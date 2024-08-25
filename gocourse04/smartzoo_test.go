package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddSector(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", NewTechnicalRoom("tr_1"))
	area.AddSector(sector)

	s, exists := area.GetSector("sector_1")
	assert.True(t, exists, "Sector should exist")
	assert.Equal(t, sector, s, "Sectors should be equal")
}

func TestDeleteSector(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", NewTechnicalRoom("tr_1"))
	area.AddSector(sector)

	_, exists := area.GetSector("sector_1")
	assert.True(t, exists, "Sector shouldn't exist")

	area.DeleteSector("sector_1")
	_, exists = area.GetSector("sector_1")
	assert.False(t, exists, "Sector shouldn't exist")
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
	sector := NewSector("sector_1", "Parrots", NewTechnicalRoom("tr_1"))
	animal := NewAnimal("animal_1", "Frodo")
	sector.AddAnimal(animal)

	_, exists := sector.GetAnimal("animal_1")
	assert.True(t, exists, "Animal should exist")

	sector.DeleteAnimal("animal_1")
	_, exists = sector.GetAnimal("animal_1")
	assert.False(t, exists, "Animal shouldn't exist")
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
	zoo := NewZoo("Happy zoo")
	area := NewArea("area_1", "Happy birds", "Feathered")
	zoo.AddArea(area)

	_, exists := zoo.GetArea("area_1")
	assert.True(t, exists, "Area should exist")

	zoo.DeleteArea("area_1")
	_, exists = zoo.GetArea("area_1")
	assert.False(t, exists, "Area shouldn't exist")
}

func TestLookupAnimalByID(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", nil)
	sector2 := NewSector("sector_2", "Ducks", nil)
	animal := NewAnimal("animal_1", "Frodo")
	animal2 := NewAnimal("animal_2", "Bilbo")
	sector.AddAnimal(animal)
	sector.AddAnimal(animal2)
	area.AddSector(sector)
	area.AddSector(sector2)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	sectorFound, ok := zoo.LookupAnimalByID("animal_2")

	assert.True(t, ok, "Animal should be found")
	assert.Equal(t, sectorFound, sector, "Sectors should be equal")
}

func TestFailLookupAnimalByID(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", nil)
	sector2 := NewSector("sector_2", "Ducks", nil)
	animal := NewAnimal("animal_1", "Frodo")
	animal2 := NewAnimal("animal_2", "Bilbo")
	sector.AddAnimal(animal)
	sector.AddAnimal(animal2)
	area.AddSector(sector)
	area.AddSector(sector2)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	sectorFound, ok := zoo.LookupAnimalByID("animal_3")

	assert.False(t, ok, "Animal should be found")
	assert.Equal(t, sectorFound, Sector{}, "Sectors should be equal")
}

func TestLookupAnimalByName(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", nil)
	sector2 := NewSector("sector_2", "Ducks", nil)
	animal := NewAnimal("animal_1", "Frodo")
	animal2 := NewAnimal("animal_2", "Bilbo")
	animal3 := NewAnimal("animal_3", "Frodo")
	sector.AddAnimal(animal)
	sector.AddAnimal(animal2)
	sector2.AddAnimal(animal3)
	area.AddSector(sector)
	area.AddSector(sector2)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	sectorsFound, ok := zoo.LookupAnimalByName("Frodo")

	assert.True(t, ok, "Animals should be found")
	assert.Len(t, sectorsFound, 2, "Should be 2 sectors")
}

func TestFailLookupAnimalByName(t *testing.T) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", nil)
	sector2 := NewSector("sector_2", "Ducks", nil)
	animal := NewAnimal("animal_1", "Frodo")
	animal2 := NewAnimal("animal_2", "Bilbo")
	animal3 := NewAnimal("animal_3", "Frodo")
	sector.AddAnimal(animal)
	sector.AddAnimal(animal2)
	sector2.AddAnimal(animal3)
	area.AddSector(sector)
	area.AddSector(sector2)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	sectorsFound, ok := zoo.LookupAnimalByName("Gaga")

	assert.False(t, ok, "Animals shouldn't be found")
	assert.Len(t, sectorsFound, 0, "Should be 0 sectors")
}

func TestMoveAnimal(t *testing.T) {
	sectorFrom := NewSector("sector_1", "Parrots", NewTechnicalRoom("tr_1"))
	animal := NewAnimal("animal_1", "Frodo")
	sectorFrom.AddAnimal(animal)

	sectorTo := NewSector("sector_2", "Ducks", nil)

	MoveAnimal(&sectorFrom, &sectorTo, animal)

	_, exists := sectorFrom.GetAnimal("animal_1")
	assert.False(t, exists, "Animal shouldn't exist")

	a, ok := sectorTo.GetAnimal("animal_1")
	assert.True(t, ok, "Animal should exist")
	assert.Equal(t, animal, a, "Animals should be equal")
}

func TestFailMoveAnimal(t *testing.T) {
	sectorFrom := NewSector("sector_1", "Parrots", NewTechnicalRoom("tr_1"))
	animal := NewAnimal("animal_1", "Frodo")

	sectorTo := NewSector("sector_2", "Ducks", nil)

	MoveAnimal(&sectorFrom, &sectorTo, animal)

	_, exists := sectorFrom.GetAnimal("animal_1")
	assert.False(t, exists, "Animal shouldn't exist")

	sectorTo.AddAnimal(animal)
	sectorFrom.AddAnimal(animal)

	_, ok := sectorTo.GetAnimal("animal_1")
	assert.True(t, ok, "Animal should exist") // Movement was not made
}

func BenchmarkLookupAnimalByID(b *testing.B) {
	area := NewArea("area_1", "Happy birds", "Feathered")
	sector := NewSector("sector_1", "Parrots", nil)
	animal := NewAnimal("animal_1", "Frodo")
	sector.AddAnimal(animal)
	sector2 := NewSector("sector_2", "Ducks", nil)
	area.AddSector(sector)
	area.AddSector(sector2)
	zoo := NewZoo("Happy zoo")
	zoo.AddArea(area)

	animals := make(map[string]Animal)
	for i := 10; i < 100; i++ {
		id := fmt.Sprintf("animal_%d", i)
		animals[id] = NewAnimal(id, fmt.Sprintf("Frodo %d", i))
	}

	for _, a := range animals {
		sector.AddAnimal(a)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		zoo.LookupAnimalByID("animal_1")
	}
}
