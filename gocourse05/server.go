package main

import (
	"errors"
	"fmt"

	"gocourse05/camera"
)

const serverURL = "http://remote.animals.control"

type Processor interface {
	RetrieveData() []camera.Data
	ProcessData() (camera.ProcessedData, error)
}

type Server struct {
	Memory    []camera.ProcessedData
	processor Processor
}

func NewServer(p Processor) *Server {
	return &Server{
		Memory:    make([]camera.ProcessedData, 0),
		processor: p,
	}
}

func (s *Server) setProcessor(p Processor) {
	s.processor = p
}

func (s *Server) saveProcessedData() error {
	processedData, err := s.processor.ProcessData()
	if err != nil {
		return err
	}

	s.Memory = append(s.Memory, processedData)

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
