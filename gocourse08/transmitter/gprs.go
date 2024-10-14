package transmitter

import "fmt"

type GPRS struct {
	SignalAvailable bool
	Buffer          []AnimalData
}

func (g *GPRS) Transmit(data AnimalData) {
	if g.SignalAvailable {
		fmt.Println("Передача даних на сервер:", data)
	} else {
		g.Buffer = append(g.Buffer, data)
	}
}

func (g *GPRS) SendBufferedData() {
	if len(g.Buffer) == 0 {
		return
	}
	for _, data := range g.Buffer {
		fmt.Println("Передача буферизованих даних на сервер:", data)
	}
	g.Buffer = nil
}
