package sensor

import "math/rand/v2"

type SoundSensor struct{}

func (s *SoundSensor) ReadData() float64 {
	return rand.Float64() * 4
}
