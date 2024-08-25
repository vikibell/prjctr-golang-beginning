package main

import (
	"fmt"

	"gocourse05/camera"
)

const serverUrl = "http://remote.animals.control"

type Memory []camera.ProcessedData

func saveProcessedData(p camera.Processor, m *Memory) error {
	processedData, err := p.ProcessData()
	if err != nil {
		return err
	}

	*m = append(*m, processedData)

	return nil
}

func sendProcessedData(m Memory, url string) {
	dataToSend := ""
	for _, data := range m {
		dataToSend = dataToSend + data.AnimalMovement
	}

	fmt.Printf("Sending data \"%s\" to %s", dataToSend, url)
}
