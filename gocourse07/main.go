package main

import (
	"context"
	"math/rand/v2"
	"sync"
	"time"
)

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

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	dataChannel := make(chan Data, 4)

	var wg sync.WaitGroup

	lightSensors := generateLightSensors(1)
	wetnessSensors := generateWetnessSensors(2)
	temperatureSensors := generateTemperatureSensors(3)

	wg.Add(4)
	go measureSensors(lightSensors, dataChannel, &wg)
	go measureSensors(wetnessSensors, dataChannel, &wg)
	go measureSensors(temperatureSensors, dataChannel, &wg)
	go centralSystem(ctx, dataChannel, &wg)

	wg.Wait()
}
