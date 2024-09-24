package gocourse08

import (
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/sensor"
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/transmitter"
)

func main() {
	heartRateTempSensor := &sensor.TemperatureSensor{}
	breathSensor := &sensor.BreathSensor{}
	soundSensor := &sensor.SoundSensor{}

	gps := &transmitter.GPRSTransmitter{SignalAvailable: false}

	dataChannel := make(chan transmitter.AnimalData)

	go func() {
		for data := range dataChannel {
			_ = gps.TransmitData(data)
		}
	}()

	transmitter.CollectAndTransmit("Лев", gps, heartRateTempSensor, breathSensor, soundSensor)

	// Імітація того, що з'явився сигнал GPRS і передача буферизованих даних
	gps.SignalAvailable = true
	gps.SendBufferedData()

	close(dataChannel)
}
