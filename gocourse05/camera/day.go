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
	id        int
	name      string
	data      []Data
	lightType lightType
}

func NewDayCamera(id int, name string, data []Data) DayCamera {
	return DayCamera{
		id:        id,
		name:      name,
		data:      data,
		lightType: sunLight,
	}
}

func (dc DayCamera) ID() int {
	return dc.id
}

func (dc DayCamera) RetrieveData() []Data {
	return dc.data
}

func (dc DayCamera) ProcessData() (ProcessedData, error) {
	cameraData := dc.RetrieveData()
	if len(cameraData) == 0 {
		return ProcessedData{}, errors.New("not enough data for processing")
	}

	processedData := NewProcessedData(time.Now(), "")
	for _, data := range cameraData {
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		processedData.AnimalMovement += animalMovement
	}

	return processedData, nil
}
