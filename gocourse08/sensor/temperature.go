package sensor

import "math/rand/v2"

type TemperatureSensor struct{}

func (s *TemperatureSensor) ReadData() float64 {
	return rand.Float64()
}
