package transmitter

import "github.com/vikibell/prjctr-golang-beginning/gocourse08/sensor"

type AnimalData struct {
	Pulse       int
	Temperature float64
	BreathRate  int
	SoundLevel  float64
}

type DataTransmitter interface {
	Transmit(data AnimalData)
	SendBufferedData()
}

func Collect(pulse sensor.Sensor[int], temperature sensor.Sensor[float64], breath sensor.Sensor[int], sound sensor.Sensor[float64]) AnimalData {
	pulseRate := pulse.ReadData()
	temperatureLevel := temperature.ReadData()
	breathRate := breath.ReadData()
	soundLevel := sound.ReadData()

	return AnimalData{
		Pulse:       pulseRate,
		Temperature: temperatureLevel,
		BreathRate:  breathRate,
		SoundLevel:  soundLevel,
	}
}
