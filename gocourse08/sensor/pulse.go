package sensor

import "math/rand/v2"

type Pulse struct{}

func (s *Pulse) ReadData() int {
	return rand.IntN(100)
}
