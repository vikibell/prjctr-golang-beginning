package sensor

import "math/rand/v2"

type Breath struct{}

func (s *Breath) ReadData() int {
	return rand.IntN(200)
}
