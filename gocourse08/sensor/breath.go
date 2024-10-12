package sensor

import "math/rand/v2"

type BreathSensor struct{}

func (s *BreathSensor) ReadData() int {
	return rand.IntN(200)
}
