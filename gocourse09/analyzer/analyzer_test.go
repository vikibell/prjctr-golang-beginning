package analyzer

import (
	"github.com/google/go-cmp/cmp"
	"gocourse09/zone"
	"testing"
)

func TestZoneAnalyzer_Analyze(t *testing.T) {
	feedingZone := zone.FeedingZone{}
	zoneAnalyzer := ZoneAnalyzer{}

	t.Run("No animals in zone", func(t *testing.T) {
		got := zoneAnalyzer.Analyze(&feedingZone)

		want := Result{
			AnimalsInZone: false,
			Species:       nil,
		}

		if !cmp.Equal(got, want) {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, want)
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
		feedingZone.AddAnimal(bigCow)
		feedingZone.AddAnimal(smallCow)

		got := zoneAnalyzer.Analyze(&feedingZone)
		want := Result{
			AnimalsInZone: true,
			Species: map[zone.Specie]int{
				zone.Bull: 2,
			},
		}

		if !cmp.Equal(got, want) {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, want)
		}
	})
}
