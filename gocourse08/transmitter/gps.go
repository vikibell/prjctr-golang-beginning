package transmitter

import "fmt"

type GPRSTransmitter struct {
	SignalAvailable bool
	Buffer          []AnimalData
}

func (g *GPRSTransmitter) TransmitData(data AnimalData) {
	if g.SignalAvailable {
		fmt.Println("Передача даних на сервер:", data)
	} else {
		g.Buffer = append(g.Buffer, data)
	}
}

func (g *GPRSTransmitter) SendBufferedData() {
	if len(g.Buffer) > 0 {
		for _, data := range g.Buffer {
			fmt.Println("Передача буферизованих даних на сервер:", data)
		}
		g.Buffer = nil
	}
}
