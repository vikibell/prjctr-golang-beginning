package main

import "fmt"

func main() {
	zoo := NewZoo("zoo_1", "Zoo 12")

	areaFeathered := NewArea("area_1", "Fly fly", "feathered")
	areaUngulates := NewArea("area_2", "Jump jump", "ungulated")
	zoo.AddArea(areaFeathered)
	zoo.AddArea(areaUngulates)

	technicalRoomA := NewTechnicalRoom("room_1")
	sectorA := NewSector("sector_1", "Sector A", &technicalRoomA)
	sectorB := NewSector("sector_2", "Sector B", nil)
	technicalRoomC := NewTechnicalRoom("room_3")
	sectorC := NewSector("sector_3", "Sector C", &technicalRoomC)
	areaFeathered.AddSector(sectorA)
	areaFeathered.AddSector(sectorB)
	areaFeathered.AddSector(sectorC)

	technicalRoomD := NewTechnicalRoom("room_4")
	sectorD := NewSector("sector_4", "Sector D", &technicalRoomD)
	sectorE := NewSector("sector_5", "Sector E", nil)
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

	sector := zoo.LookupAnimalByID("animal_4")

	if sector != nil {
		fmt.Println(sector.Name)
	} else {
		fmt.Println("Animal was not found.")
	}

	sector = zoo.LookupAnimalByID("animal_88")

	if sector != nil {
		fmt.Println(sector.Name)
	} else {
		fmt.Println("Animal was not found.")
	}

	sectors := zoo.LookupAnimalByName("A")
	if len(sectors) > 0 {
		for _, sector := range sectors {
			fmt.Println(sector.Name)
		}
	} else {
		fmt.Println("Animal was not found.")
	}

	sectors = zoo.LookupAnimalByName("Y")
	if len(sectors) > 0 {
		for _, sector := range sectors {
			fmt.Println(sector.Name)
		}
	} else {
		fmt.Println("Animal was not found.")
	}
}
