package main

import (
	"math/rand/v2"
	"sync"
	"time"
)

func generateLightSensors(n int) []Sensor {
	var lightSensors []Sensor
	for i := 0; i < n; i++ {
		lightSensor := LightSensor{
			ID:    i,
			Light: rand.IntN(100),
		}
		lightSensors = append(lightSensors, lightSensor)
	}
	return lightSensors
}

func generateWetnessSensors(n int) []Sensor {
	var wetnessSensors []Sensor
	for i := 0; i < n; i++ {
		wetnessSensor := WetnessSensor{
			ID:      i,
			Wetness: rand.IntN(100),
		}
		wetnessSensors = append(wetnessSensors, wetnessSensor)
	}
	return wetnessSensors
}

func generateTemperatureSensors(n int) []Sensor {
	var temperatureSensors []Sensor
	for i := 0; i < n; i++ {
		temperatureSensor := TemperatureSensor{
			ID:          i,
			Temperature: rand.IntN(100),
		}
		temperatureSensors = append(temperatureSensors, temperatureSensor)
	}
	return temperatureSensors
}

func main() {
	dataChannel := make(chan Data)
	sensorStopChannel := make(chan struct{})
	centralStopChannel := make(chan struct{})

	var wg sync.WaitGroup

	lightSensors := generateLightSensors(1)
	wetnessSensors := generateWetnessSensors(2)
	temperatureSensors := generateTemperatureSensors(3)

	wg.Add(3)
	go measureSensors(lightSensors, dataChannel, sensorStopChannel, &wg)
	go measureSensors(wetnessSensors, dataChannel, sensorStopChannel, &wg)
	go measureSensors(temperatureSensors, dataChannel, sensorStopChannel, &wg)

	wg.Add(1)
	go centralSystem(dataChannel, centralStopChannel, &wg)

	time.Sleep(10 * time.Second)

	go func() {
		wg.Wait()
		close(sensorStopChannel)
		close(dataChannel)
		close(centralStopChannel)
	}()
}
