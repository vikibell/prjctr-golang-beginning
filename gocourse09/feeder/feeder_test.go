package feeder

import (
	a "gocourse09/analyzer"
	z "gocourse09/zone"
	"reflect"
	"sort"
	"testing"
)

func TestAutomaticFeeder_PourOn(t *testing.T) {
	feedingZone := z.FeedingZone{}
	zoneAnalyzer := a.ZoneAnalyzer{}
	feeder := AutomaticFeeder{}

	t.Run("No animals in zone", func(t *testing.T) {
		result := zoneAnalyzer.Analyze(&feedingZone)

		got := feeder.PourOn(result)
		if got != nil {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, nil)
		}
	})

	t.Run("Animals in zone", func(t *testing.T) {
		bigCow := z.Animal{
			Name:   "Milka",
			Specie: z.Bull,
		}
		smallCow := z.Animal{
			Name:   "Milka",
			Specie: z.Bull,
		}
		smallHorse := z.Animal{
			Name:   "Willy",
			Specie: z.Horse,
		}
		feedingZone.AddAnimal(bigCow)
		feedingZone.AddAnimal(smallCow)
		feedingZone.AddAnimal(smallHorse)

		result := zoneAnalyzer.Analyze(&feedingZone)
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
		sort.Sort(got)

		if len(got) != len(want) {
			t.Fatalf("Analyze() should return slice with 2 items.")
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, want)
		}
	})
}
