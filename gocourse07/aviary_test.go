package main

import (
	"sync"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestTemperatureSensor_measure(t *testing.T) {
	ts := TemperatureSensor{
		ID:          1,
		Temperature: 100,
	}

	data := ts.measure()

	want := Data{SensorID: 1, Measurement: 100}
	if data != want {
		t.Errorf("TemperatureSensor.measure(): got=%+v, want=%+v", data, want)
	}
}

func TestLightSensor_measure(t *testing.T) {
	ts := LightSensor{
		ID:    1,
		Light: 100,
	}

	data := ts.measure()

	want := Data{SensorID: 1, Measurement: 100}
	if !cmp.Equal(data, want) {
		t.Errorf("LightSensor.measure(): got=%+v, want=%+v", data, want)
	}
}

func TestWetnessSensor_measure(t *testing.T) {
	ts := WetnessSensor{
		ID:      1,
		Wetness: 100,
	}

	data := ts.measure()

	want := Data{SensorID: 1, Measurement: 100}
	if !cmp.Equal(data, want) {
		t.Errorf("WetnessSensor.measure(): got=%+v, want=%+v", data, want)
	}

}

func Test_measureSensors(t *testing.T) {
	dataChannel := make(chan Data, 3)

	var wg sync.WaitGroup

	lightSensors := LightSensor{
		ID:    1,
		Light: 20,
	}
	wetnessSensors := WetnessSensor{
		ID:      2,
		Wetness: 40,
	}
	temperatureSensors := TemperatureSensor{
		ID:          3,
		Temperature: 10,
	}
	sensors := []Sensor{lightSensors, wetnessSensors, temperatureSensors}

	wg.Add(1)
	go measureSensors(sensors, dataChannel, &wg)
	wg.Wait()
	close(dataChannel)

	received := make([]int, 4)

	for data := range dataChannel {
		received[data.SensorID] = data.Measurement
	}

	want := make([]int, 4)
	want[1] = 20
	want[2] = 40
	want[3] = 10

	for i, got := range received {
		if got != want[i] {
			t.Errorf("Want %d at position %d, but got %d", want[i], i, got)
		}
	}
}

func TestRaceConditions_run(t *testing.T) {
	run()
}

func TestDeadLock_run(t *testing.T) {
	done := make(chan bool)
	go func() {
		run()
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Test timed out, possible deadlock detected")
	}
}
