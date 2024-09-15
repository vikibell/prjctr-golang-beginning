package main

import (
	"context"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {
	run()
}

func generateLightSensors(n int) []Sensor {
	sensors := make([]Sensor, n)
	for i := range sensors {
		sensors[i] = LightSensor{
			ID:    i,
			Light: rand.IntN(100),
		}
	}
	return sensors
}

func generateWetnessSensors(n int) []Sensor {
	sensors := make([]Sensor, n)
	for i := range sensors {
		sensors[i] = WetnessSensor{
			ID:      i,
			Wetness: rand.IntN(100),
		}
	}
	return sensors
}

func generateTemperatureSensors(n int) []Sensor {
	sensors := make([]Sensor, n)
	for i := range sensors {
		sensors[i] = TemperatureSensor{
			ID:          i,
			Temperature: rand.IntN(100),
		}
	}
	return sensors
}

func run() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	dataChannel := make(chan Data, 36)

	var wg sync.WaitGroup

	lightSensors := generateLightSensors(10)
	wetnessSensors := generateWetnessSensors(15)
	temperatureSensors := generateTemperatureSensors(11)

	wg.Add(4)
	go measureSensors(lightSensors, dataChannel, &wg)
	go measureSensors(wetnessSensors, dataChannel, &wg)
	go measureSensors(temperatureSensors, dataChannel, &wg)
	go centralSystem(ctx, dataChannel, &wg)
	wg.Wait()
}
