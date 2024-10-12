package transmitter

type AnimalData struct {
	Pulse       int
	Temperature float64
	BreathRate  int
	SoundLevel  float64
}

type DataTransmitter interface {
	TransmitData(data AnimalData)
	SendBufferedData()
}
