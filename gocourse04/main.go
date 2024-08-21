package main

import "fmt"

func main() {
	zoo := NewZoo("Zoo 12")

	areaFeathered := NewArea("area_1", "Fly fly", "feathered")
	areaUngulates := NewArea("area_2", "Jump jump", "ungulated")
	zoo.AddArea(areaFeathered)
	zoo.AddArea(areaUngulates)

	sectorA := NewSector("sector_1", "Sector A", NewTechnicalRoom("room_1"))
	sectorB := NewSector("sector_2", "Sector B", NewTechnicalRoom("room_2"))
	sectorC := NewSector("sector_3", "Sector C", NewTechnicalRoom("room_3"))
	areaFeathered.AddSector(sectorA)
	areaFeathered.AddSector(sectorB)
	areaFeathered.AddSector(sectorC)

	sectorD := NewSector("sector_4", "Sector D", NewTechnicalRoom("room_4"))
	sectorE := NewSector("sector_5", "Sector E", NewTechnicalRoom("room_5"))
	areaUngulates.AddSector(sectorD)
	areaFeathered.AddSector(sectorE)

	animalA := NewAnimal("animal_1", "A")
	animalB := NewAnimal("animal_2", "A")
	sectorA.AddAnimal(animalA)
	sectorA.AddAnimal(animalB)
	animalC := NewAnimal("animal_3", "C")
	animalD := NewAnimal("animal_4", "D")
	animalE := NewAnimal("animal_5", "E")
	sectorD.AddAnimal(animalC)
	sectorD.AddAnimal(animalD)
	sectorE.AddAnimal(animalE)

	sectorD.TechnicalRoom.clean(sectorD)

	sectorA.TechnicalRoom.feedAnimal(sectorA, animalA.ID)
	sectorC.TechnicalRoom.feedAnimal(sectorC, animalA.ID)

	MoveAnimal(&sectorA, &sectorB, animalA)
	MoveAnimal(&sectorA, &sectorB, animalA)

	sector, found := zoo.LookupAnimalByID("animal_4")

	if found {
		fmt.Println(sector.Name)
	}

	sector, found = zoo.LookupAnimalByID("animal_88")

	if found {
		fmt.Println(sector.Name)
	}

	sectors, found := zoo.LookupAnimalByName("A")

	if found {
		for _, sector := range sectors {
			fmt.Println(sector.Name)
		}
	}

	sectors, found = zoo.LookupAnimalByName("Y")

	if found {
		for _, sector := range sectors {
			fmt.Println(sector.Name)
		}
	}
}
