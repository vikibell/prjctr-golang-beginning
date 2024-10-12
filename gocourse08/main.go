package main

import (
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/sensor"
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/transmitter"
)

func main() {
	temperature := &sensor.TemperatureSensor{}
	breath := &sensor.BreathSensor{}
	sound := &sensor.SoundSensor{}
	pulse := &sensor.PulseSensor{}

	gps := &transmitter.GPRSTransmitter{SignalAvailable: true}
	for range 5 {
		gps.TransmitData(transmitter.CollectData(pulse, temperature, breath, sound))
	}

	gps.SignalAvailable = false
	for range 5 {
		gps.TransmitData(transmitter.CollectData(pulse, temperature, breath, sound))
	}

	gps.SendBufferedData()
}
