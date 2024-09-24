package sensor

type BreathSensor struct{}

func (s *BreathSensor) ReadData() int {
	return rand.Float64() * 10
}
