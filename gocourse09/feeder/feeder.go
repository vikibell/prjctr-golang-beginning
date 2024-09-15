package feeder

import (
	a "gocourse09/analyzer"
	z "gocourse09/zone"
)

type FoodBracket struct {
	Amount int
	Type   string
}

type Feeder interface {
	PourOn(ar a.Result) SortedFoodBrackets
}

type AutomaticFeeder struct{}

func (f AutomaticFeeder) PourOn(ar a.Result) SortedFoodBrackets {
	if !ar.AnimalsInZone {
		return nil
	}

	var brackets []FoodBracket

	for specie, count := range ar.Species {
		fb := FoodBracket{
			Amount: count,
		}
		switch specie {
		case z.Bull, z.Buffalo:
			fb.Type = "gross"
		case z.Horse:
			fb.Type = "apples"
		case z.Elk:
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
