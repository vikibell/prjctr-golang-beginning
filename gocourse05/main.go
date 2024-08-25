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

	dayCamera := camera.NewDayCamera(2, "Canon XY-22", []camera.Data{
		camera.NewDayCameraData(1, "Заєць", "Біжить вперед"),
		camera.NewDayCameraData(2, "Лисиця", "Біжить вперед"),
	})

	infraredCamera := camera.NewInfraredCamera(1, "Nikon S60", []camera.Data{
		camera.NewInfraredCameraData(1, "Медвідь", "Чешить спину"),
		camera.NewInfraredCameraData(2, "Білка", "Ість горішок"),
	})

	err := saveProcessedData(&nightCamera, &ramMemory)
	if err != nil {
		fmt.Printf("Saving from camera %d data to memory failed. Error: %v\n", nightCamera.ID, err)
	}

	err = saveProcessedData(&dayCamera, &ramMemory)
	if err != nil {
		fmt.Printf("Saving from camera %d data to memory failed. Error: %v\n", dayCamera.ID, err)
	}

	err = saveProcessedData(&infraredCamera, &ramMemory)
	if err != nil {
		fmt.Printf("Saving from camera %d data to memory failed. Error: %v\n", infraredCamera.ID, err)
	}

	err = sendProcessedData(ramMemory, serverUrl)
	if err != nil {
		fmt.Printf("Sending data failed: %v\n", err)
	}
}
