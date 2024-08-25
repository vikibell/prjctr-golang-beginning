package camera

import (
	"fmt"
	"time"
)

type InfraredCamera struct {
	ID        int
	Name      string
	Data      []Data
	lightType int
}

func NewInfraredCamera(id int, name string, data []Data) InfraredCamera {
	return InfraredCamera{
		ID:        id,
		Name:      name,
		Data:      data,
		lightType: infraredLight,
	}
}

func (ic InfraredCamera) retrieveData() ([]Data, error) {
	// TODO error
	return ic.Data, nil
}

func (ic InfraredCamera) ProcessData() (ProcessedData, error) {
	// TODO error, change way
	processedData := NewProcessedData(time.Now(), "")
	cameraData, err := ic.retrieveData()
	if err != nil {
		return ProcessedData{}, err
	}

	for _, data := range cameraData {
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		processedData.AnimalMovement = processedData.AnimalMovement + animalMovement
	}

	return processedData, nil
}
