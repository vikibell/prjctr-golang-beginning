package camera

import (
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
	// TODO error
	return nc.Data, nil
}

func (nc NightCamera) ProcessData() (ProcessedData, error) {
	// TODO error
	processedData := NewProcessedData(time.Now(), "")
	cameraData, err := nc.retrieveData()
	if err != nil {
		return ProcessedData{}, err
	}

	for _, data := range cameraData {
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		animalMovement = strings.ReplaceAll(animalMovement, nightCameraArtifact, "")
		processedData.AnimalMovement = processedData.AnimalMovement + animalMovement
	}

	return processedData, nil
}
