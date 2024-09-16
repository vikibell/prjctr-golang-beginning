package feeder

import (
	"cmp"
	"gocourse09/zone"
	"slices"
)

type FoodBracket struct {
	Amount int
	Type   string
}

type Feeder interface {
	PourOn(ar zone.Result) []FoodBracket
}

type AutomaticFeeder struct{}

func (f AutomaticFeeder) PourOn(ar zone.Result) []FoodBracket {
	if !ar.AnimalsInZone {
		return nil
	}

	var brackets []FoodBracket

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

	slices.SortFunc(brackets, func(a, b FoodBracket) int {
		return cmp.Compare(a.Amount, b.Amount)
	})

	return brackets
}
