package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"

	"gocourse09/analyzer"
	"gocourse09/feeder"
	"gocourse09/zone"
)

func generateAnimals(n int) []zone.Animal {
	animals := make([]zone.Animal, n)
	for i := range animals {
		animals[i] = zone.Animal{
			Name:   "Name" + strconv.Itoa(i),
			Specie: zone.Specie(rand.IntN(10)),
		}
	}

	return animals
}

func main() {
	animals := generateAnimals(100)
	z := zone.FeedingZone{}
	for _, animal := range animals {
		z.AddAnimal(animal)
	}

	zoneAnalyzer := analyzer.ZoneAnalyzer{}
	result := zoneAnalyzer.Analyze(&z)

	f := feeder.AutomaticFeeder{}
	foodBracket := f.PourOn(result)

	for _, food := range foodBracket {
		fmt.Printf("Poured on %s in amount of %d\n", food.Type, food.Amount)
	}
}
