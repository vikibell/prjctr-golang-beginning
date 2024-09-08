package main

func main() {
	cow := Animal{
		ID:     0,
		Name:   "",
		Specie: "",
	}

	horse := Animal{
		ID:     0,
		Name:   "",
		Specie: "",
	}

	zone := FeedingZone{
		ID:      0,
		Animals: nil,
	}
	zone.addAnimal(cow)
	zone.addAnimal(horse)

	zoneAnalyzer := ZoneAnalyzer{}
	result := zoneAnalyzer.analyze(zone)

	food := NewFoodBracket(result)

	feeder := AutomaticFeeder{}
	feeder.feed(cow, food)
	feeder.feed(horse, food)
}
