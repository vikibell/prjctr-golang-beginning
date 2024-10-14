package transmitter

import (
	"github.com/stretchr/testify/assert"
	"github.com/vikibell/prjctr-golang-beginning/gocourse08/sensor"
	"testing"
)

func TestCollectData(t *testing.T) {
	temperature := &sensor.Temperature{}
	breath := &sensor.Breath{}
	sound := &sensor.Sound{}
	pulse := &sensor.Pulse{}

	animalData := Collect(pulse, temperature, breath, sound)

	assert.NotEmpty(t, animalData.Temperature)
	assert.NotEmpty(t, animalData.Pulse)
	assert.NotEmpty(t, animalData.BreathRate)
	assert.NotEmpty(t, animalData.SoundLevel)
}
