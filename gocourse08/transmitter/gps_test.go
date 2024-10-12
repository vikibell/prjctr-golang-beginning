package transmitter

import (
	"github.com/stretchr/testify/assert"
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/sensor"
	"testing"
)

func TestGPRSTransmitter_SendBufferedData(t *testing.T) {
	temperature := &sensor.TemperatureSensor{}
	breath := &sensor.BreathSensor{}
	sound := &sensor.SoundSensor{}
	pulse := &sensor.PulseSensor{}

	gps := &GPRSTransmitter{SignalAvailable: false}
	for range 5 {
		gps.TransmitData(CollectData(pulse, temperature, breath, sound))
	}

	assert.Len(t, gps.Buffer, 5)

	gps.SendBufferedData()
	assert.Len(t, gps.Buffer, 0)
}

func TestGPRSTransmitter_TransmitData(t *testing.T) {
	temperature := &sensor.TemperatureSensor{}
	breath := &sensor.BreathSensor{}
	sound := &sensor.SoundSensor{}
	pulse := &sensor.PulseSensor{}

	gps := &GPRSTransmitter{SignalAvailable: true}
	for range 5 {
		gps.TransmitData(CollectData(pulse, temperature, breath, sound))
	}

	assert.Len(t, gps.Buffer, 0)
}
