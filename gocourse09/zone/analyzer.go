package zone

type Analyzer interface {
	Analyze(zone Zone) Result
}

type Result struct {
	AnimalsInZone bool
	Species       map[Specie]int // Map with specie type as index and count as key
}

type AnimalsAnalyzer struct{}

func (za AnimalsAnalyzer) Analyze(z Zone) Result {
	animals := z.GetAnimals()

	if len(animals) == 0 {
		return Result{}
	}

	result := Result{
		AnimalsInZone: true,
	}

	species := make(map[Specie]int)
	for _, animal := range animals {
		species[animal.Specie]++
	}
	result.Species = species

	return result
}
