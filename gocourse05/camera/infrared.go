package camera

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const infraredCameraArtifact = " infrared"

func NewInfraredCameraData(id int, animal, movement string) Data {
	return Data{
		ID:       id,
		Animal:   animal + infraredCameraArtifact,
		Movement: movement + infraredCameraArtifact,
	}
}

type InfraredCamera struct {
	id        int
	name      string
	data      []Data
	lightType lightType
}

func NewInfraredCamera(id int, name string, data []Data) InfraredCamera {
	return InfraredCamera{
		id:        id,
		name:      name,
		data:      data,
		lightType: infraredLight,
	}
}

func (ic InfraredCamera) ID() int {
	return ic.id
}

func (ic InfraredCamera) RetrieveData() []Data {
	return ic.data
}

func (ic InfraredCamera) ProcessData() (ProcessedData, error) {
	cameraData := ic.RetrieveData()
	if len(cameraData) == 0 {
		return ProcessedData{}, errors.New("not enough data for processing")
	}

	processedData := NewProcessedData(time.Now(), "")
	for _, data := range cameraData {
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		animalMovement = strings.ReplaceAll(animalMovement, infraredCameraArtifact, "")
		processedData.AnimalMovement += animalMovement
	}

	return processedData, nil
}
