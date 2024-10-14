package main

import (
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/sensor"
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/transmitter"
)

func main() {
	temperature := &sensor.Temperature{}
	breath := &sensor.Breath{}
	sound := &sensor.Sound{}
	pulse := &sensor.Pulse{}

	gprs := &transmitter.GPRS{SignalAvailable: true}
	for range 5 {
		gprs.Transmit(transmitter.Collect(pulse, temperature, breath, sound))
	}

	gprs.SignalAvailable = false
	for range 5 {
		gprs.Transmit(transmitter.Collect(pulse, temperature, breath, sound))
	}

	gprs.SendBufferedData()
}
