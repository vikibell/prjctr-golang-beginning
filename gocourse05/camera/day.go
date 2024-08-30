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

func (dc DayCamera) retrieveData() (*[]Data, error) {
	if len(dc.Data) == 0 {
		return nil, fmt.Errorf("no camera data found")
	}
	return &dc.Data, nil
}

func (dc DayCamera) ProcessData() (*ProcessedData, error) {
	cameraData, err := dc.retrieveData()
	if err != nil {
		return nil, err
	}

	processedData := NewProcessedData(time.Now(), "")
	for _, data := range *cameraData {
		if len(data.Animal) == 0 || len(data.Movement) == 0 {
			return nil, fmt.Errorf("not enought data for processing")
		}
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		processedData.AnimalMovement = processedData.AnimalMovement + animalMovement
	}

	return &processedData, nil
}
