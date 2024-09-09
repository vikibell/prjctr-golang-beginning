package main

type Specie int

const (
	bull Specie = iota
	horse
	buffalo
	elk
)

type Animal struct {
	Name   string
	Specie Specie
}

type Zone interface {
	GetAnimals() []Animal
	AddAnimal(a Animal)
}

type FeedingZone struct {
	Animals []Animal
}

func (fz *FeedingZone) GetAnimals() []Animal {
	return fz.Animals
}

func (fz *FeedingZone) AddAnimal(a Animal) {
	fz.Animals = append(fz.Animals, a)
}

type AnalyzerResult struct {
	AnimalsInZone bool
	Species       map[Specie]int // Map with specie type as index and count as key
}

type ZoneAnalyzer struct{}

func (za ZoneAnalyzer) Analyze(zone Zone) AnalyzerResult {
	animals := zone.GetAnimals()
	result := AnalyzerResult{
		AnimalsInZone: false,
		Species:       make(map[Specie]int),
	}

	if len(animals) == 0 {
		return result
	}

	result.AnimalsInZone = true
	for _, animal := range animals {
		result.Species[animal.Specie] += 1
	}

	return result
}

type FoodBracket struct {
	Amount int
	Type   string
}

type Feeder interface {
	pourOn(ar AnalyzerResult) []FoodBracket
}

type AutomaticFeeder struct{}

func (f AutomaticFeeder) pourOn(ar AnalyzerResult) []FoodBracket {
	if !ar.AnimalsInZone {
		return nil
	}

	var brackets []FoodBracket

	for specie, count := range ar.Species {
		fb := FoodBracket{
			Amount: count,
		}
		switch specie {
		case bull, buffalo:
			fb.Type = "gross"
		case horse:
			fb.Type = "apples"
		case elk:
			fb.Type = "eggs"
		default:
			fb.Type = "dry food"
		}
		brackets = append(brackets, fb)
	}

	return brackets
}
