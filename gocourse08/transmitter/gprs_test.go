package transmitter

import (
	"github.com/stretchr/testify/assert"
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/sensor"
	"testing"
)

func TestGPRSTransmitter_SendBufferedData(t *testing.T) {
	temperature := &sensor.Temperature{}
	breath := &sensor.Breath{}
	sound := &sensor.Sound{}
	pulse := &sensor.Pulse{}

	gprs := &GPRS{SignalAvailable: false}
	for range 5 {
		gprs.Transmit(Collect(pulse, temperature, breath, sound))
	}

	assert.Len(t, gprs.Buffer, 5)

	gprs.SendBufferedData()
	assert.Len(t, gprs.Buffer, 0)
}

func TestGPRSTransmitter_TransmitData(t *testing.T) {
	temperature := &sensor.Temperature{}
	breath := &sensor.Breath{}
	sound := &sensor.Sound{}
	pulse := &sensor.Pulse{}

	gprs := &GPRS{SignalAvailable: true}
	for range 5 {
		gprs.Transmit(Collect(pulse, temperature, breath, sound))
	}

	assert.Empty(t, gprs.Buffer)
}
