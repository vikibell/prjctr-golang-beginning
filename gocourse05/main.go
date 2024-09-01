package main

import (
	"fmt"

	"gocourse05/camera"
)

func main() {
	server := Server{Memory: make([]camera.ProcessedData, 0)}

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

	err := server.saveProcessedData(&nightCamera)
	if err != nil {
		fmt.Printf("Saving from camera %d data to memory failed. Error: %v\n", nightCamera.ID, err)
	}

	err = server.saveProcessedData(&dayCamera)
	if err != nil {
		fmt.Printf("Saving from camera %d data to memory failed. Error: %v\n", dayCamera.ID, err)
	}

	err = server.saveProcessedData(&infraredCamera)
	if err != nil {
		fmt.Printf("Saving from camera %d data to memory failed. Error: %v\n", infraredCamera.ID, err)
	}

	err = server.sendProcessedData(serverURL)
	if err != nil {
		fmt.Printf("Sending data failed: %v\n", err)
	}
}
