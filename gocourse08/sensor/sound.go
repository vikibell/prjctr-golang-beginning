package sensor

import "math/rand/v2"

type Sound struct{}

func (s *Sound) ReadData() float64 {
	return rand.Float64() * 4
}
