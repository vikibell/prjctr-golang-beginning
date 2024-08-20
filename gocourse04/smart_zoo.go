package main

import "fmt"

type Animal struct {
	ID   string
	Name string
}

type Areas map[string]Area

type Area struct {
	ID      string
	Name    string
	Type    string            //пернаті
	Sectors map[string]Sector //загони для кожного виду тварин "папуги", "качки"
}

type Sector struct {
	ID            string
	Name          string
	Animals       map[string]Animal
	TechnicalRoom *TechnicalRoom
}

type TechnicalRoom struct {
	ID          string
	Instruments map[string]int
}

func (tr *TechnicalRoom) clean(sectorName string) {
	fmt.Printf("I'm cleaning sector: %s\n", sectorName)
}

func (tr *TechnicalRoom) feedAnimal(sector Sector, animalId string) {
	animal, exists := sector.Animals[animalId]

	if !exists {
		fmt.Printf("There is no animal %s in sector: %s\n", animalId, sector.Name)
		return
	}

	fmt.Println("I'm feeding animal:", animal.Name)
}

type Zoo struct {
	Areas Areas
}

//func (z Zoo) LookupAnimalByName(name string) Sector {
//	for _, area := range z.Areas {
//		for key, sector := range area.Sectors {
//			for _, animal := range sector.Animals {
//				if animal.Name == name {
//
//				}
//			}
//		}
//	}
//}

func moveAnimal(sectorFrom *Sector, sectorTo *Sector, animal Animal) {
	animalToMove, exists := sectorFrom.Animals[animal.ID]

	if !exists {
		fmt.Printf("There is no animal %s in sector: %s\n", animal.ID, sectorFrom.Name)
		return
	}

	delete(sectorFrom.Animals, animal.ID)

	_, exists = sectorTo.Animals[animal.ID]

	if exists {
		fmt.Printf("Animal with id %s already exists in sector: %s\n", animal.ID, sectorTo.Name)
		return
	}

	sectorTo.Animals[animal.ID] = animalToMove
}
