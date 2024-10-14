package sensor

import "math/rand/v2"

type Temperature struct{}

func (s *Temperature) ReadData() float64 {
	return rand.Float64()
}
