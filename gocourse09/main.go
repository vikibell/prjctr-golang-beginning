package main

func main() {
	cow := Animal{
		ID:     1,
		Name:   "Milka",
		Specie: "bull",
	}

	horse := Animal{
		ID:     1,
		Name:   "Okl",
		Specie: "horse",
	}

	zone := FeedingZone{
		ID:      1,
		Animals: make([]Animal, 0),
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
