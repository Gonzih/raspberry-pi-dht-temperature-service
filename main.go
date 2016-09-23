package main

import (
	"sync"

	"github.com/d2r2/go-dht"
	"github.com/kataras/iris"
)

var gpioMutex sync.Mutex

func temperatureHandler(ctx *iris.Context) {
	gpioMutex.Lock()
	defer gpioMutex.Unlock()

	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT11, 4, true, 10)

	if err != nil {
		ctx.Write("Error occured during temp readout: %s", err)
	} else {
		ctx.Write("Temperature = %v*C, Humidity = %v%% (retried %d times)", temperature, humidity, retried)
	}
}

func main() {
	iris.Get("/", temperatureHandler)

	iris.Listen(":8080")
}
