package camera

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const nightCameraArtifact = " night"

func NewNightCameraData(id int, animal, movement string) Data {
	return Data{
		ID:       id,
		Animal:   animal + nightCameraArtifact,
		Movement: movement + nightCameraArtifact,
	}
}

type NightCamera struct {
	id        int
	name      string
	data      []Data
	lightType lightType
}

func NewNightCamera(id int, name string, data []Data) NightCamera {
	return NightCamera{
		id:        id,
		name:      name,
		data:      data,
		lightType: noLight,
	}
}

func (nc NightCamera) ID() int {
	return nc.id
}

func (nc NightCamera) RetrieveData() []Data {
	return nc.data
}

func (nc NightCamera) ProcessData() (ProcessedData, error) {
	cameraData := nc.RetrieveData()
	if len(cameraData) == 0 {
		return ProcessedData{}, errors.New("not enough data for processing")
	}

	processedData := NewProcessedData(time.Now(), "")
	for _, data := range cameraData {
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		animalMovement = strings.ReplaceAll(animalMovement, nightCameraArtifact, "")
		processedData.AnimalMovement += animalMovement
	}

	return processedData, nil
}
