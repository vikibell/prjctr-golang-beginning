package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSector(t *testing.T) {
	tr := NewTechnicalRoom("tr_1")
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", &tr)
	area.AddSector(sector)

	s, exists := area.GetSector("sector_1")
	assert.True(t, exists, "Sector should exist")
	assert.Equal(t, sector, s, "Sectors should be equal")
}

func TestDeleteSector(t *testing.T) {
	tr := NewTechnicalRoom("tr_1")
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", &tr)
	area.AddSector(sector)

	area.DeleteSector("sector_1")
	_, exists := area.GetSector("sector_1")
	assert.False(t, exists, "Sector shouldn't exist")
}

func TestAddAnimal(t *testing.T) {
	tr := NewTechnicalRoom("tr_1")
	sector := NewSector("sector_1", "Parrots", &tr)
	animal := NewAnimal("animal_1", "Frodo")
	sector.AddAnimal(animal)

	a, exists := sector.GetAnimal("animal_1")
	assert.True(t, exists, "Animal should exist")
	assert.Equal(t, animal, a, "Animals should be equal")
}

func TestDeleteAnimal(t *testing.T) {
	tr := NewTechnicalRoom("tr_1")
	sector := NewSector("sector_1", "Parrots", &tr)
	animal := NewAnimal("animal_1", "Frodo")
	sector.AddAnimal(animal)

	sector.DeleteAnimal("animal_1")
	_, exists := sector.GetAnimal("animal_1")
	assert.False(t, exists, "Animal shouldn't exist")
}

func TestAddArea(t *testing.T) {
	zoo := NewZoo("Happy zoo")
	area := NewArea("area_1", "Happy birds", "Feathered")
	zoo.AddArea(area)

	a, exists := zoo.GetArea("area_1")
	assert.True(t, exists, "Area should exist")
	assert.Equal(t, area, a, "Areas should be equal")
}

func TestDeleteArea(t *testing.T) {
	zoo := NewZoo("Happy zoo")
	area := NewArea("area_1", "Happy birds", "Feathered")
	zoo.AddArea(area)

	zoo.DeleteArea("area_1")
	_, exists := zoo.GetArea("area_1")
	assert.False(t, exists, "Area shouldn't exist")
}

func TestLookupAnimalByID(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	parrotsSector := NewSector("sector_1", "Parrots", nil)
	ducksSector := NewSector("sector_2", "Ducks", nil)
	frodo := NewAnimal("animal_1", "Frodo")
	bilbo := NewAnimal("animal_2", "Bilbo")
	parrotsSector.AddAnimal(frodo)
	parrotsSector.AddAnimal(bilbo)
	area.AddSector(parrotsSector)
	area.AddSector(ducksSector)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	sectorFound := zoo.LookupAnimalByID("animal_2")

	assert.Equal(t, parrotsSector.ID, sectorFound.ID, "Sectors should be equal")
}

func TestFailLookupAnimalByID(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	parrotsSector := NewSector("sector_1", "Parrots", nil)
	ducksSector := NewSector("sector_2", "Ducks", nil)
	frodo := NewAnimal("animal_1", "Frodo")
	bilbo := NewAnimal("animal_2", "Bilbo")
	parrotsSector.AddAnimal(frodo)
	parrotsSector.AddAnimal(bilbo)
	area.AddSector(parrotsSector)
	area.AddSector(ducksSector)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	sectorFound := zoo.LookupAnimalByID("animal_3")

	assert.Nil(t, sectorFound, "Sector should be nil")
}

func TestLookupAnimalByName(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	parrotsSector := NewSector("sector_1", "Parrots", nil)
	ducksSector := NewSector("sector_2", "Ducks", nil)
	animal := NewAnimal("animal_1", "Frodo")
	bilbo := NewAnimal("animal_2", "Bilbo")
	frodo := NewAnimal("animal_3", "Frodo")
	parrotsSector.AddAnimal(animal)
	parrotsSector.AddAnimal(bilbo)
	ducksSector.AddAnimal(frodo)
	area.AddSector(parrotsSector)
	area.AddSector(ducksSector)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	sectorsFound := zoo.LookupAnimalByName("Frodo")

	assert.Len(t, sectorsFound, 2, "Should be 2 sectors")
}

func TestFailLookupAnimalByName(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	parrotsSector := NewSector("sector_1", "Parrots", nil)
	ducksSector := NewSector("sector_2", "Ducks", nil)
	animal := NewAnimal("animal_1", "Frodo")
	bilbo := NewAnimal("animal_2", "Bilbo")
	frodo := NewAnimal("animal_3", "Frodo")
	parrotsSector.AddAnimal(animal)
	parrotsSector.AddAnimal(bilbo)
	ducksSector.AddAnimal(frodo)
	area.AddSector(parrotsSector)
	area.AddSector(ducksSector)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	sectorsFound := zoo.LookupAnimalByName("Gaga")

	assert.Empty(t, sectorsFound, "Sectors should be empty")
}

func TestMoveAnimal(t *testing.T) {
	tr := NewTechnicalRoom("tr_1")
	from := NewSector("sector_1", "Parrots", &tr)
	animal := NewAnimal("animal_1", "Frodo")
	from.AddAnimal(animal)

	to := NewSector("sector_2", "Ducks", nil)

	MoveAnimal(&from, &to, animal)

	_, exists := from.GetAnimal("animal_1")
	assert.False(t, exists, "Animal shouldn't exist")

	a, ok := to.GetAnimal("animal_1")
	assert.True(t, ok, "Animal should exist")
	assert.Equal(t, animal, a, "Animals should be equal")
}

func TestFailMoveAnimal(t *testing.T) {
	tr := NewTechnicalRoom("tr_1")
	from := NewSector("sector_1", "Parrots", &tr)
	animal := NewAnimal("animal_1", "Frodo")

	to := NewSector("sector_2", "Ducks", nil)

	MoveAnimal(&from, &to, animal)

	_, exists := from.GetAnimal("animal_1")
	assert.False(t, exists, "Animal shouldn't exist")

	to.AddAnimal(animal)
	from.AddAnimal(animal)

	_, ok := to.GetAnimal("animal_1")
	assert.True(t, ok, "Animal should exist") // Movement was not made
}

func BenchmarkLookupAnimalByID(b *testing.B) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	parrotsSector := NewSector("sector_1", "Parrots", nil)
	animal := NewAnimal("animal_1", "Frodo")
	parrotsSector.AddAnimal(animal)
	ducksSector := NewSector("sector_2", "Ducks", nil)
	area.AddSector(parrotsSector)
	area.AddSector(ducksSector)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	animals := make(map[string]Animal)
	for i := 10; i < 100; i++ {
		id := fmt.Sprintf("animal_%d", i)
		animal := NewAnimal(id, fmt.Sprintf("Frodo %d", i))
		animals[id] = animal
		parrotsSector.AddAnimal(animal)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zoo.LookupAnimalByID("animal_1")
	}
}
