package main

import "fmt"

func main() {
	bigCow := Animal{
		Name:   "Milka",
		Specie: bull,
	}

	bigHorse := Animal{
		Name:   "Okl",
		Specie: horse,
	}

	smallHorse := Animal{
		Name:   "Willy",
		Specie: horse,
	}

	cat := Animal{
		Name:   "Anna",
		Specie: 100,
	}

	zone := FeedingZone{
		Animals: make([]Animal, 0),
	}
	zone.AddAnimal(bigCow)
	zone.AddAnimal(bigHorse)
	zone.AddAnimal(smallHorse)
	zone.AddAnimal(cat)

	zoneAnalyzer := ZoneAnalyzer{}
	result := zoneAnalyzer.Analyze(&zone)

	feeder := AutomaticFeeder{}
	foodBracket := feeder.pourOn(result)

	for _, food := range foodBracket {
		fmt.Printf("Poured on %s in amount of %d\n", food.Type, food.Amount)
	}
}
