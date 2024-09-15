package main

import (
	"fmt"

	"gocourse09/analyzer"
	"gocourse09/feeder"
	"gocourse09/zone"
)

func main() {
	bigCow := zone.Animal{
		Name:   "Milka",
		Specie: zone.Bull,
	}

	bigHorse := zone.Animal{
		Name:   "Okl",
		Specie: zone.Horse,
	}

	smallHorse := zone.Animal{
		Name:   "Willy",
		Specie: zone.Horse,
	}

	cat := zone.Animal{
		Name:   "Anna",
		Specie: 100,
	}

	z := zone.FeedingZone{}
	z.AddAnimal(bigCow)
	z.AddAnimal(bigHorse)
	z.AddAnimal(smallHorse)
	z.AddAnimal(cat)

	zoneAnalyzer := analyzer.ZoneAnalyzer{}
	result := zoneAnalyzer.Analyze(&z)

	f := feeder.AutomaticFeeder{}
	foodBracket := f.PourOn(result)

	for _, food := range foodBracket {
		fmt.Printf("Poured on %s in amount of %d\n", food.Type, food.Amount)
	}
}
