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

type ZoneAnalyzer struct{}

type AnalyzerResult struct {
	AnimalsInZone bool
	Species       map[Specie]int // Map with specie type as index and count as key
}

func (za ZoneAnalyzer) Analyze(zone Zone) AnalyzerResult {
	animals := zone.GetAnimals()
	result := AnalyzerResult{
		AnimalsInZone: false,
		Species:       nil,
	}

	if len(animals) == 0 {
		return result
	}

	result.AnimalsInZone = true
	species := make(map[Specie]int)
	for _, animal := range animals {
		species[animal.Specie] += 1
	}
	result.Species = species

	return result
}

type FoodBracket struct {
	Amount int
	Type   string
}

type Feeder interface {
	pourOn(ar AnalyzerResult) SortedFoodBrackets
}

type AutomaticFeeder struct{}

func (f AutomaticFeeder) pourOn(ar AnalyzerResult) SortedFoodBrackets {
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

type SortedFoodBrackets []FoodBracket

func (fb SortedFoodBrackets) Len() int      { return len(fb) }
func (fb SortedFoodBrackets) Swap(i, j int) { fb[i], fb[j] = fb[j], fb[i] }
func (fb SortedFoodBrackets) Less(i, j int) bool {
	return fb[i].Amount > fb[j].Amount
}
