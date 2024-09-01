package main

import (
	"errors"
	"fmt"

	"gocourse05/camera"
)

const serverURL = "http://remote.animals.control"

type Memory []camera.ProcessedData

func saveProcessedData(p camera.Processor, m *Memory) error {
	processedData, err := p.ProcessData()
	if err != nil {
		return err
	}

	*m = append(*m, *processedData)

	return nil
}

func sendProcessedData(m Memory, url string) error {
	dataToSend := ""
	for _, data := range m {
		dataToSend += data.AnimalMovement
	}

	if len(dataToSend) > 0 {
		fmt.Printf("Sending data %q to %s", dataToSend, url)
		return nil
	}

	return errors.New("no data for send")
}
