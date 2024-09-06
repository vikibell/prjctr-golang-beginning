package camera

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewInfraredCameraData(id int, animal, movement string) Data {
	return Data{
		ID:       id,
		Animal:   animal + "_" + uuid.NewString(),
		Movement: movement,
	}
}

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

func (ic InfraredCamera) RetrieveData() ([]Data, error) {
	if len(ic.Data) == 0 {
		return nil, errors.New("no camera data found")
	}

	return ic.Data, nil
}

func (ic InfraredCamera) ProcessData() (*ProcessedData, error) {
	cameraData, err := ic.RetrieveData()
	if err != nil {
		return nil, err
	}

	processedData := NewProcessedData(time.Now(), "")
	re := regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`)
	for _, data := range cameraData {
		data.Animal = strings.ReplaceAll(re.ReplaceAllString(data.Animal, ""), "_", "")
		animalMovement := fmt.Sprintf("%s, %s; ", data.Animal, data.Movement)
		processedData.AnimalMovement += animalMovement
	}

	return &processedData, nil
}
