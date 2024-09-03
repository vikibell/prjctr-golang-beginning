package main

import (
	"fmt"
	"sync"
	"time"
)

type Sensor interface {
	measure() Data
}

type TemperatureSensor struct {
	ID          int
	Temperature int
}

func (t TemperatureSensor) measure() Data {
	return NewData(t.ID, t.Temperature)
}

type LightSensor struct {
	ID    int
	Light int
}

func (l LightSensor) measure() Data {
	return NewData(l.ID, l.Light)
}

type WetnessSensor struct {
	ID      int
	Wetness int
}

func (w WetnessSensor) measure() Data {
	return NewData(w.ID, w.Wetness)
}

type Data struct {
	SensorID    int
	Measurement int
}

func NewData(ID int, measurement int) Data {
	return Data{
		SensorID:    ID,
		Measurement: measurement,
	}
}

func measureSensors(sensors []Sensor, dataCh chan<- Data, stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	select {
	case <-stopCh:
		fmt.Println("Sensor measurement is shutting down.")
		return
	default:
		for _, sensor := range sensors {
			fmt.Printf("Sensor measure data: %+v\n", sensor.measure())
			dataCh <- sensor.measure()
		}
		time.Sleep(2 * time.Second)
	}
}

func centralSystem(dataCh <-chan Data, stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stopCh:
			fmt.Println("Central system is shutting down after processing all data.")
			return
		case data, ok := <-dataCh:
			if !ok {
				return
			}
			fmt.Printf("Central system processing data: %+v\n", data)
			time.Sleep(2 * time.Second)
		}
	}
}
