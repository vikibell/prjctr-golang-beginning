package sensor

import "math/rand"

type TemperatureSensor struct{}

func (s *TemperatureSensor) ReadData() (int, float64) {
	// Генерація випадкових значень для прикладу
	return rand.Intn(40) + 60, rand.Float64()*5 + 35
}
