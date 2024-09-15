package feeder

import (
	"gocourse09/analyzer"
	"gocourse09/zone"
	"sort"
)

type FoodBracket struct {
	Amount int
	Type   string
}

type Feeder interface {
	PourOn(ar analyzer.Result) SortedFoodBrackets
}

type AutomaticFeeder struct{}

func (f AutomaticFeeder) PourOn(ar analyzer.Result) SortedFoodBrackets {
	if !ar.AnimalsInZone {
		return nil
	}

	var brackets SortedFoodBrackets

	for specie, count := range ar.Species {
		fb := FoodBracket{
			Amount: count,
		}
		switch specie {
		case zone.Bull, zone.Buffalo:
			fb.Type = "gross"
		case zone.Horse:
			fb.Type = "apples"
		case zone.Elk:
			fb.Type = "eggs"
		default:
			fb.Type = "dry food"
		}
		brackets = append(brackets, fb)
	}
	sort.Sort(brackets)

	return brackets
}

type SortedFoodBrackets []FoodBracket

func (fb SortedFoodBrackets) Len() int      { return len(fb) }
func (fb SortedFoodBrackets) Swap(i, j int) { fb[i], fb[j] = fb[j], fb[i] }
func (fb SortedFoodBrackets) Less(i, j int) bool {
	return fb[i].Amount > fb[j].Amount
}
