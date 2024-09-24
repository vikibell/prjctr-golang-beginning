package sensor

import "math/rand"

type PulseSensor struct{}

func (s *PulseSensor) ReadData() (int, float64) {
	return rand.Intn(40) + 60, rand.Float64()*5 + 35
}
