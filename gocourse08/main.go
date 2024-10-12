package main

import (
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/sensor"
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/transmitter"
)

func collectData(pulse sensor.Sensor[int], temperature sensor.Sensor[float64], breath sensor.Sensor[int], sound sensor.Sensor[float64]) transmitter.AnimalData {
	pulseRate := pulse.ReadData()
	temperatureLevel := temperature.ReadData()
	breathRate := breath.ReadData()
	soundLevel := sound.ReadData()

	return transmitter.AnimalData{
		Pulse:       pulseRate,
		Temperature: temperatureLevel,
		BreathRate:  breathRate,
		SoundLevel:  soundLevel,
	}
}

func main() {
	temperature := &sensor.TemperatureSensor{}
	breath := &sensor.BreathSensor{}
	sound := &sensor.SoundSensor{}
	pulse := &sensor.PulseSensor{}

	gps := &transmitter.GPRSTransmitter{SignalAvailable: true}
	for range 5 {
		gps.TransmitData(collectData(pulse, temperature, breath, sound))
	}

	gps.SignalAvailable = false
	for range 5 {
		gps.TransmitData(collectData(pulse, temperature, breath, sound))
	}

	gps.SendBufferedData()
}
