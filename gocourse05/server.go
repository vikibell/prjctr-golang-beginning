package main

import (
	"errors"
	"fmt"

	"gocourse05/camera"
)

const serverURL = "http://remote.animals.control"

type Server struct {
	Memory []camera.ProcessedData
}

func (s *Server) saveProcessedData(p camera.Processor) error {
	processedData, err := p.ProcessData()
	if err != nil {
		return err
	}

	s.Memory = append(s.Memory, *processedData)

	return nil
}

func (s *Server) sendProcessedData(url string) error {
	dataToSend := ""
	for _, data := range s.Memory {
		dataToSend += data.AnimalMovement
	}

	if len(dataToSend) > 0 {
		fmt.Printf("Sending data %q to %s", dataToSend, url)
		return nil
	}

	return errors.New("no data for send")
}
