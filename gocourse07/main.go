package main

import (
	"context"
	"math/rand/v2"
	"sync"
	"time"
)

func generateLightSensors(n int) []Sensor {
	lightSensors := make([]Sensor, n)
	for i := range lightSensors {
		lightSensors[i] = LightSensor{
			ID:    i,
			Light: rand.IntN(100),
		}
	}
	return lightSensors
}

func generateWetnessSensors(n int) []Sensor {
	wetnessSensors := make([]Sensor, n)
	for i := range wetnessSensors {
		wetnessSensors[i] = WetnessSensor{
			ID:      i,
			Wetness: rand.IntN(100),
		}
	}
	return wetnessSensors
}

func generateTemperatureSensors(n int) []Sensor {
	temperatureSensors := make([]Sensor, n)
	for i := range temperatureSensors {
		temperatureSensors[i] = TemperatureSensor{
			ID:          i,
			Temperature: rand.IntN(100),
		}
	}
	return temperatureSensors
}

func main() {

	deadline := time.Now().Add(100 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	dataChannel := make(chan Data)

	var wg sync.WaitGroup

	lightSensors := generateLightSensors(1)
	wetnessSensors := generateWetnessSensors(2)
	temperatureSensors := generateTemperatureSensors(3)

	wg.Add(4)
	go measureSensors(lightSensors, dataChannel, ctx, &wg)
	go measureSensors(wetnessSensors, dataChannel, ctx, &wg)
	go measureSensors(temperatureSensors, dataChannel, ctx, &wg)
	go centralSystem(dataChannel, ctx, &wg)

	time.Sleep(10 * time.Second)

	go func() {
		wg.Wait()
		close(dataChannel)
	}()
}
