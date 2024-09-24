package sensor

type SoundSensor struct{}

func (s *SoundSensor) ReadData() float64 {
	// Генерація випадкових значень для прикладу
	return rand.Float64() * 100
}
