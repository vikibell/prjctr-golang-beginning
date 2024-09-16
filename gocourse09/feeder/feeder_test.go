package feeder

import (
	"github.com/google/go-cmp/cmp"
	"gocourse09/zone"
	"sort"
	"testing"
)

func TestAutomaticFeeder_PourOn(t *testing.T) {
	feedingZone := zone.FeedingZone{}
	animalsAnalyzer := zone.AnimalsAnalyzer{}
	feeder := AutomaticFeeder{}

	t.Run("No animals in zone", func(t *testing.T) {
		result := animalsAnalyzer.Analyze(&feedingZone)

		got := feeder.PourOn(result)
		if got != nil {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, nil)
		}
	})

	t.Run("Animals in zone", func(t *testing.T) {
		bigCow := zone.Animal{
			Name:   "Milka",
			Specie: zone.Bull,
		}
		smallCow := zone.Animal{
			Name:   "Milka",
			Specie: zone.Bull,
		}
		smallHorse := zone.Animal{
			Name:   "Willy",
			Specie: zone.Horse,
		}
		feedingZone.AddAnimal(bigCow)
		feedingZone.AddAnimal(smallCow)
		feedingZone.AddAnimal(smallHorse)

		result := animalsAnalyzer.Analyze(&feedingZone)
		got := feeder.PourOn(result)

		want := SortedFoodBrackets{
			{
				Amount: 2,
				Type:   "gross",
			},
			{
				Amount: 1,
				Type:   "apples",
			},
		}
		sort.Sort(want)

		if !cmp.Equal(got, want) {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, want)
		}
	})
}
