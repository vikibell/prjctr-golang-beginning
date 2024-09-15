package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
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

}

func Test_centralSystem(t *testing.T) {

}

//TODO race conditions
//TODO dead locks
