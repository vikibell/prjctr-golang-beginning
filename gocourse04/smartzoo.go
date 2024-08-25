package main

import "fmt"

type Animal struct {
	ID   string
	Name string
}

func NewAnimal(id, name string) Animal {
	return Animal{
		ID:   id,
		Name: name,
	}
}

type Area struct {
	ID      string
	Name    string
	Type    string
	Sectors map[string]Sector // Index is Sector.ID
}

func NewArea(id, name, areaType string) Area {
	return Area{
		ID:      id,
		Name:    name,
		Type:    areaType,
		Sectors: make(map[string]Sector),
	}
}

func (a *Area) AddSector(s Sector) {
	a.Sectors[s.ID] = s
}

func (a *Area) GetSector(id string) (Sector, bool) {
	s, exists := a.Sectors[id]
	return s, exists
}

func (a *Area) DeleteSector(id string) {
	delete(a.Sectors, id)
}

type Sector struct {
	ID            string
	Name          string
	Animals       map[string]Animal
	TechnicalRoom *TechnicalRoom
}

func NewSector(id string, name string, tr *TechnicalRoom) Sector {
	return Sector{
		ID:            id,
		Name:          name,
		Animals:       make(map[string]Animal),
		TechnicalRoom: tr,
	}
}

func (s *Sector) AddAnimal(a Animal) {
	s.Animals[a.ID] = a
}

func (s *Sector) GetAnimal(id string) (Animal, bool) {
	a, exists := s.Animals[id]
	return a, exists
}

func (s *Sector) DeleteAnimal(id string) {
	delete(s.Animals, id)
}

type TechnicalRoom struct {
	ID          string
	Instruments map[string]int
}

func NewTechnicalRoom(id string) *TechnicalRoom {
	return &TechnicalRoom{
		ID: id,
		Instruments: map[string]int{
			"shovel":  3,
			"basin":   4,
			"nippers": 2,
		},
	}
}

func (tr *TechnicalRoom) clean(sector Sector) {
	fmt.Printf("I'm cleaning sector: %s\n", sector.Name)
}

func (tr *TechnicalRoom) feedAnimal(sector Sector, animalID string) {
	animal, exists := sector.GetAnimal(animalID)

	if !exists {
		fmt.Printf("There is no animal %s in sector: %s\n", animalID, sector.Name)
		return
	}

	fmt.Println("I'm feeding animal:", animal.Name)
}

type Zoo struct {
	Name  string
	Areas map[string]Area // Index is Area.ID
}

func NewZoo(name string) Zoo {
	return Zoo{
		Name:  name,
		Areas: make(map[string]Area),
	}
}

func (z *Zoo) AddArea(a Area) {
	z.Areas[a.ID] = a
}

func (z *Zoo) GetArea(id string) (Area, bool) {
	a, exists := z.Areas[id]
	return a, exists
}

func (z *Zoo) DeleteArea(id string) {
	delete(z.Areas, id)
}

func (z *Zoo) LookupAnimalByName(name string) ([]Sector, bool) {
	foundSectors := make([]Sector, 0)

	for _, area := range z.Areas {
		for key, sector := range area.Sectors {
			for _, animal := range sector.Animals {
				if animal.Name == name {
					foundSectors = append(foundSectors, area.Sectors[key])
				}
			}
		}
	}

	if len(foundSectors) > 0 {
		return foundSectors, true
	}

	fmt.Printf("Animal with name %s was not found.\n", name)

	return foundSectors, false
}

func (z *Zoo) LookupAnimalByID(id string) (Sector, bool) {
	foundSector := Sector{}
	found := false

loopExit:
	for _, area := range z.Areas {
		for key, sector := range area.Sectors {
			for _, animal := range sector.Animals {
				if animal.ID == id {
					foundSector = area.Sectors[key]
					found = true
					break loopExit
				}
			}
		}
	}

	if !found {
		fmt.Printf("Animal with ID %s was not found.\n", id)
	}

	return foundSector, found
}

func MoveAnimal(from, to *Sector, animal Animal) {
	animalToMove, exists := from.GetAnimal(animal.ID)

	if !exists {
		fmt.Printf("There is no animal %s in sector: %s\n", animal.ID, from.Name)
		return
	}

	from.DeleteAnimal(animal.ID)

	_, exists = to.GetAnimal(animal.ID)

	if exists {
		fmt.Printf("Animal with id %s already exists in sector: %s.\n", animal.ID, to.Name)
		return
	}

	fmt.Printf("Moving %s from sector %s to sector %s.\n", animal.Name, from.Name, to.Name)
	to.AddAnimal(animalToMove)
}
