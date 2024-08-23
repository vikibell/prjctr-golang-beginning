package main

import (
	"fmt"
)

//TODO defer/tests/main/functions

const noLight = 0
const infraredLight = 1
const sunLight = 3

type Picture string

type CameraData struct {
	ID         int
	Picture    Picture
	AnimalData AnimalData
}

type AnimalData struct {
	ID       int
	Animal   string
	Movement string
}

type Memory []AnimalData

type processor interface {
	processData()
	retrieveData() ([]CameraData, error)
}

type sender interface {
	sendData(m Memory, url string) error
}

type NightCamera struct {
	ID        int
	Name      string
	Data      []CameraData
	lightType int
}

func NeeNightCamera(id int, name string, data []CameraData) InfraredCamera {
	return InfraredCamera{
		ID:        id,
		Name:      name,
		Data:      data,
		lightType: noLight,
	}
}

func (nc NightCamera) processData() {
	//TODO process CameraData and add AnimalData, add some errors?
}

func (nc NightCamera) retrieveData() ([]CameraData, error) {
	//TODO
	return nc.Data, nil
}

func (nc NightCamera) sendData(m Memory, url string) error {
	//TODO
	return nil
}

type InfraredCamera struct {
	ID        int
	Name      string
	Data      []CameraData
	lightType int
}

func NewInfraredCamera(id int, name string, data []CameraData) InfraredCamera {
	return InfraredCamera{
		ID:        id,
		Name:      name,
		Data:      data,
		lightType: infraredLight,
	}
}

func (ic InfraredCamera) processData() {
	//TODO
}

func (ic InfraredCamera) retrieveData() ([]CameraData, error) {
	//TODO
	return ic.Data, nil
}

func (ic InfraredCamera) sendData(m Memory, url string) error {
	//TODO
	return nil
}

type DayCamera struct {
	ID        int
	Name      string
	Data      []CameraData
	lightType int
}

func NewDayCamera(id int, name string, data []CameraData) InfraredCamera {
	return InfraredCamera{
		ID:        id,
		Name:      name,
		Data:      data,
		lightType: sunLight,
	}
}

func (dc DayCamera) processData() {
	//TODO
}

func (dc DayCamera) retrieveData() ([]CameraData, error) {
	//TODO
	return dc.Data, nil
}

func (dc DayCamera) sendData(m Memory, url string) error {
	//TODO
	return nil
}

func processCameraData(p processor) {
	p.processData()
}

func saveCameraData(p processor, m Memory) (bool, error) {
	cameraData, err := p.retrieveData()

	if err != nil {
		fmt.Println("Something went wrong....")
		return false, err
	}

	for _, data := range cameraData {
		m = append(m, data.AnimalData)
	}

	fmt.Printf("Succesfully saved %d pictures to memory.\n", len(m))
	return true, nil
}

func sendCameraData(s sender, m Memory, url string) error {
	err := s.sendData(m, url)

	if err != nil {
		return err
	}

	return nil
}
