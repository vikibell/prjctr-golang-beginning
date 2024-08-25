package camera

import "time"

const (
	noLight       = 0
	infraredLight = 1
	sunLight      = 3
)

type Data struct {
	ID       int
	Animal   string
	Movement string
}

func NewCameraData(id int, animal, movement string) Data {
	return Data{
		ID:       id,
		Animal:   animal,
		Movement: movement,
	}
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
	ProcessData() (ProcessedData, error)
}
