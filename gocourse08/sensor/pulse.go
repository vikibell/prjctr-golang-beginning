package sensor

import "math/rand/v2"

type PulseSensor struct{}

func (s *PulseSensor) ReadData() int {
	return rand.IntN(100)
}
