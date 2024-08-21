package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestAddSector(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", NewTechnicalRoom("tr_1"))
	area.AddSector(sector)

	s, exists := area.GetSector("sector_1")
	assert.True(t, exists, "Sector should exist")
	assert.Equal(t, area, s, "Sectors should be equal")
}

func TestDeleteSector(t *testing.T) {

}

func TestAddAnimal(t *testing.T) {
	sector := NewSector("sector_1", "Parrots", NewTechnicalRoom("tr_1"))
	animal := NewAnimal("animal_1", "Frodo")
	sector.AddAnimal(animal)

	s, exists := sector.GetAnimal("animal_1")
	assert.True(t, exists, "Animal should exist")
	assert.Equal(t, animal, s, "Animals should be equal")
}

func TestDeleteAnimal(t *testing.T) {

}

func TestAddArea(t *testing.T) {
	zoo := NewZoo("Happy zoo")
	area := NewArea("area_1", "Happy birds", "Feathered")
	zoo.AddArea(area)

	z, exists := zoo.GetArea("area_1")
	assert.True(t, exists, "Area should exist")
	assert.Equal(t, area, z, "Areas should be equal")
}

func TestDeleteArea(t *testing.T) {

}

func TestCleanSector(t *testing.T) {

}

func TestFeedAnimal(t *testing.T) {

}

func TestLookupAnimalByID(t *testing.T) {

}

func TestFailLookupAnimalByID(t *testing.T) {

}

func TestLookupAnimalByName(t *testing.T) {

}

func TestFailLookupAnimalByName(t *testing.T) {

}

func TestMoveAnimal(t *testing.T) {

}

func TestFailMoveAnimal(t *testing.T) {

}

func BenchmarkLookupAnimalByID(b *testing.B) {
}
