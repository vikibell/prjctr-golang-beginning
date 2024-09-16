package zone

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestAnimalsAnalyzer_Analyze(t *testing.T) {
	feedingZone := FeedingZone{}
	animalsAnalyzer := AnimalsAnalyzer{}

	t.Run("No animals in zone", func(t *testing.T) {
		got := animalsAnalyzer.Analyze(&feedingZone)

		want := Result{
			AnimalsInZone: false,
			Species:       nil,
		}

		if !cmp.Equal(got, want) {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, want)
		}
	})

	t.Run("Animals in zone", func(t *testing.T) {
		bigCow := Animal{
			Name:   "Milka",
			Specie: Bull,
		}
		smallCow := Animal{
			Name:   "Milka",
			Specie: Bull,
		}
		feedingZone.AddAnimal(bigCow)
		feedingZone.AddAnimal(smallCow)

		got := animalsAnalyzer.Analyze(&feedingZone)
		want := Result{
			AnimalsInZone: true,
			Species: map[Specie]int{
				Bull: 2,
			},
		}

		if !cmp.Equal(got, want) {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, want)
		}
	})
}
