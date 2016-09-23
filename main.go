package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/d2r2/go-dht"
	"github.com/julienschmidt/httprouter"
)

var gpioMutex sync.Mutex

func temperatureHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	gpioMutex.Lock()
	defer gpioMutex.Unlock()

	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, 4, true, 10)

	if err != nil {
		fmt.Fprintf(w, "Error occured during temp readout: %s", err)
	} else {
		fmt.Fprintf(w, "Temperature = %v*C, Humidity = %v%% (retried %d times)", temperature, humidity, retried)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", temperatureHandler)

	log.Println("Ready to serve!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
