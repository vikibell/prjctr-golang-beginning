package main

import "fmt"

type Specie string

type Animal struct {
	ID     int
	Name   string
	Specie Specie
}

type FoodBracket struct {
	ID     int
	Amount int
	Type   string
}

type Zone interface {
	getAnimals() []Animal
	addAnimal(a Animal)
}

type FeedingZone struct {
	ID      int
	Animals []Animal
}

func (fz FeedingZone) getAnimals() []Animal {
	return fz.getAnimals()
}

func (fz FeedingZone) addAnimal(a Animal) {
	fz.Animals = append(fz.Animals, a)
}

type AnalyzerResult struct {
	AnimalsInZone bool
	Species       []Specie
	Count         int
}

type ZoneAnalyzer struct{}

func (za ZoneAnalyzer) analyze(zone Zone) AnalyzerResult {
	animals := zone.getAnimals()
	result := AnalyzerResult{
		AnimalsInZone: false,
		Species:       make([]Specie, 0),
		Count:         0,
	}

	if len(animals) == 0 {
		return result
	}

	result.AnimalsInZone = true
	for i, animal := range animals {
		result.Count += 1
		result.Species[i] = animal.Specie
	}

	return result
}

type Feeder interface {
	feed(a Animal, fb FoodBracket)
}

type AutomaticFeeder struct{}

func (f AutomaticFeeder) feed(a Animal, fb FoodBracket) {
	fmt.Printf("Animal %s is eating %s.", a.Name, fb.Type)
}

func NewFoodBracket(ar AnalyzerResult) FoodBracket {
	return FoodBracket{
		ID:     0,
		Amount: 0,
		Type:   "",
	}
}
