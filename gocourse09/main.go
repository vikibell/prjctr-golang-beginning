package main

import (
	"fmt"

	a "gocourse09/analyzer"
	f "gocourse09/feeder"
	z "gocourse09/zone"
)

func main() {
	bigCow := z.Animal{
		Name:   "Milka",
		Specie: z.Bull,
	}

	bigHorse := z.Animal{
		Name:   "Okl",
		Specie: z.Horse,
	}

	smallHorse := z.Animal{
		Name:   "Willy",
		Specie: z.Horse,
	}

	cat := z.Animal{
		Name:   "Anna",
		Specie: 100,
	}

	zone := z.FeedingZone{
		Animals: make([]z.Animal, 0),
	}
	zone.AddAnimal(bigCow)
	zone.AddAnimal(bigHorse)
	zone.AddAnimal(smallHorse)
	zone.AddAnimal(cat)

	zoneAnalyzer := a.ZoneAnalyzer{}
	result := zoneAnalyzer.Analyze(&zone)

	feeder := f.AutomaticFeeder{}
	foodBracket := feeder.PourOn(result)

	for _, food := range foodBracket {
		fmt.Printf("Poured on %s in amount of %d\n", food.Type, food.Amount)
	}
}
