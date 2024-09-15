package analyzer

import z "gocourse09/zone"

type Analyzer interface {
	Analyze(zone z.Zone) Result
}

type Result struct {
	AnimalsInZone bool
	Species       map[z.Specie]int // Map with specie type as index and count as key
}

type ZoneAnalyzer struct{}

func (za ZoneAnalyzer) Analyze(zone z.Zone) Result {
	animals := zone.GetAnimals()
	result := Result{
		AnimalsInZone: false,
		Species:       nil,
	}

	if len(animals) == 0 {
		return result
	}

	result.AnimalsInZone = true
	species := make(map[z.Specie]int)
	for _, animal := range animals {
		species[animal.Specie] += 1
	}
	result.Species = species

	return result
}
