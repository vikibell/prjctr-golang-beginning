package camera

import "time"

type lightType int

const (
	noLight lightType = iota
	infraredLight
	sunLight
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
