package camera

import (
	"fmt"
	"time"
)

func NewDayCameraData(id int, animal, movement string) Data {
	return Data{
		ID:       id,
		Animal:   animal,
		Movement: movement,
	}
}

type DayCamera struct {
	ID        int
	Name      string
	Data      []Data
	lightType int
}

func NewDayCamera(id int, name string, data []Data) DayCamera {
	return DayCamera{
		ID:        id,
		Name:      name,
		Data:      data,
		lightType: sunLight,
	}
}

func (dc DayCamera) retrieveData() ([]Data, error) {
	// TODO error
	return dc.Data, nil
}

func (dc DayCamera) ProcessData() (ProcessedData, error) {
	// TODO error, change way
	processedData := NewProcessedData(time.Now(), "")
	cameraData, err := dc.retrieveData()
	if err != nil {
		return ProcessedData{}, err
	}

	for _, data := range cameraData {
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		processedData.AnimalMovement = processedData.AnimalMovement + animalMovement
	}

	return processedData, nil
}
