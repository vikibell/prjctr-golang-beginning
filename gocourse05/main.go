package main

import (
	"fmt"

	"gocourse05/camera"
)

func main() {
	var ramMemory Memory

	nightCamera := camera.NewNightCamera(1, "Canon Z200", []camera.Data{
		camera.NewNightCameraData(1, "Тигр", "Біжить ліворуч"),
		camera.NewNightCameraData(2, "Вовк", "Сидить на місці"),
	})

	dayCamera := camera.NewDayCamera(1, "Canon XY-22", []camera.Data{
		camera.NewDayCameraData(1, "Заєць", "Біжить вперед"),
		camera.NewDayCameraData(2, "Лисиця", "Біжить вперед"),
	})

	infraredCamera := camera.NewInfraredCamera(1, "Nikon S60", []camera.Data{
		camera.NewInfraredCameraData(1, "Медвідь", "Чешить спину"),
		camera.NewInfraredCameraData(2, "Білка", "Ість горішок"),
	})

	err := saveProcessedData(nightCamera, &ramMemory)
	if err != nil {
		fmt.Printf("Saving from camera id %d data to memory failed. See details: %v\n", nightCamera.ID, err)
	}

	err = saveProcessedData(dayCamera, &ramMemory)
	if err != nil {
		fmt.Printf("Saving from camera id %d data to memory failed. See details: %v\n", dayCamera.ID, err)
	}

	err = saveProcessedData(infraredCamera, &ramMemory)
	if err != nil {
		fmt.Printf("Saving from camera id %d data to memory failed. See details: %v\n", infraredCamera.ID, err)
	}

	sendProcessedData(ramMemory, serverUrl)
}
