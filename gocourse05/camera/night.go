package camera

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const nightCameraArtifact = "night"

func NewNightCameraData(id int, animal, movement string) Data {
	return Data{
		ID:       id,
		Animal:   animal + nightCameraArtifact,
		Movement: movement + nightCameraArtifact,
	}
}

type NightCamera struct {
	ID        int
	Name      string
	Data      []Data
	lightType int
}

func NewNightCamera(id int, name string, data []Data) NightCamera {
	return NightCamera{
		ID:        id,
		Name:      name,
		Data:      data,
		lightType: noLight,
	}
}

func (nc NightCamera) retrieveData() ([]Data, error) {
	if len(nc.Data) == 0 {
		return nil, errors.New("no camera data found")
	}

	return nc.Data, nil
}

func (nc NightCamera) ProcessData() (*ProcessedData, error) {
	cameraData, err := nc.retrieveData()
	if err != nil {
		return nil, err
	}

	processedData := NewProcessedData(time.Now(), "")
	for _, data := range cameraData {
		if len(data.Animal) <= len(nightCameraArtifact) || len(data.Movement) <= len(nightCameraArtifact) {
			return nil, errors.New("not enought data for processing")
		}
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		animalMovement = strings.ReplaceAll(animalMovement, nightCameraArtifact, "")
		processedData.AnimalMovement = processedData.AnimalMovement + animalMovement
	}

	return &processedData, nil
}
