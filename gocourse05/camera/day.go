package camera

import (
	"errors"
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
	lightType lightType
}

func NewDayCamera(id int, name string, data []Data) DayCamera {
	return DayCamera{
		ID:        id,
		Name:      name,
		Data:      data,
		lightType: sunLight,
	}
}

func (dc DayCamera) RetrieveData() ([]Data, error) {
	if len(dc.Data) == 0 {
		return nil, errors.New("no camera data found")
	}
	return dc.Data, nil
}

func (dc DayCamera) ProcessData() (*ProcessedData, error) {
	cameraData, err := dc.RetrieveData()
	if err != nil {
		return nil, err
	}

	processedData := NewProcessedData(time.Now(), "")
	for _, data := range cameraData {
		if data.Animal == "" || data.Movement == "" {
			return nil, errors.New("not enough data for processing")
		}
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		processedData.AnimalMovement += animalMovement
	}

	return &processedData, nil
}
