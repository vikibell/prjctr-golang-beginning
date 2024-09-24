package transmitter

import "fmt"

type GPRSTransmitter struct {
	SignalAvailable bool
	Buffer          []AnimalData
}

// TODO rework
func (g *GPRSTransmitter) TransmitData(data AnimalData) error {
	if g.SignalAvailable {
		fmt.Println("Передача даних на сервер:", data)
	} else {
		g.Buffer = append(g.Buffer, data)
	}
	return nil
}

// TODO rework
func (g *GPRSTransmitter) SendBufferedData() {
	if g.SignalAvailable && len(g.Buffer) > 0 {
		fmt.Println("Передача накопичених даних на сервер:")
		for _, data := range g.Buffer {
			fmt.Println(data)
		}
		g.Buffer = nil
	}
}
