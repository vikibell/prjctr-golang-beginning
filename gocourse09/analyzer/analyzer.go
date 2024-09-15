package analyzer

import "gocourse09/zone"

type Analyzer interface {
	Analyze(zone zone.Zone) Result
}

type Result struct {
	AnimalsInZone bool
	Species       map[zone.Specie]int // Map with specie type as index and count as key
}

type ZoneAnalyzer struct{}

func (za ZoneAnalyzer) Analyze(z zone.Zone) Result {
	animals := z.GetAnimals()

	if len(animals) == 0 {
		return Result{}
	}

	result := Result{
		AnimalsInZone: true,
	}

	species := make(map[zone.Specie]int)
	for _, animal := range animals {
		species[animal.Specie]++
	}
	result.Species = species

	return result
}
