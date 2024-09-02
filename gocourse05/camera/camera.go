package camera

import "time"

const (
	noLight       = iota
	infraredLight = iota
	sunLight      = iota
)

type Data struct {
	ID       int
	Animal   string
	Movement string
}

type ProcessedData struct {
	Time           time.Time
	AnimalMovement string
}

func NewProcessedData(time time.Time, animalMovement string) ProcessedData {
	return ProcessedData{
		Time:           time,
		AnimalMovement: animalMovement,
	}
}

type Processor interface {
	retrieveData() ([]Data, error)
	ProcessData() (*ProcessedData, error)
}
