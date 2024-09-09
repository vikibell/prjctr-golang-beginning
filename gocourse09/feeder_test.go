package main

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"sort"
	"testing"
)

func TestFeedingZone_GetAnimals(t *testing.T) {
	bigCow := Animal{
		Name:   "Milka",
		Specie: bull,
	}
	feedingZone := FeedingZone{Animals: []Animal{bigCow}}

	animals := feedingZone.GetAnimals()
	if len(animals) != 1 {
		t.Fatalf("GetAnimals() should return 1 item.")
	}

	got := animals[0]
	if got != bigCow {
		t.Errorf("Unexpected data: got=%+v, want=%+v", got, bigCow)
	}
}

func TestFeedingZone_AddAnimal(t *testing.T) {
	bigCow := Animal{
		Name:   "Milka",
		Specie: bull,
	}
	feedingZone := FeedingZone{}
	feedingZone.AddAnimal(bigCow)

	animals := feedingZone.GetAnimals()
	if len(animals) != 1 {
		t.Fatalf("GetAnimals() should return 1 item.")
	}

	got := animals[0]
	if got != bigCow {
		t.Errorf("Unexpected data: got=%+v, want=%+v", got, bigCow)
	}
}

func TestZoneAnalyzer_Analyze(t *testing.T) {
	feedingZone := FeedingZone{}
	zoneAnalyzer := ZoneAnalyzer{}

	t.Run("No animals in zone", func(t *testing.T) {
		got := zoneAnalyzer.Analyze(&feedingZone)

		want := AnalyzerResult{
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
			Specie: bull,
		}
		smallCow := Animal{
			Name:   "Milka",
			Specie: bull,
		}
		feedingZone.AddAnimal(bigCow)
		feedingZone.AddAnimal(smallCow)

		got := zoneAnalyzer.Analyze(&feedingZone)
		want := AnalyzerResult{
			AnimalsInZone: true,
			Species: map[Specie]int{
				bull: 2,
			},
		}

		if !cmp.Equal(got, want) {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, want)
		}
	})
}

func TestAutomaticFeeder_PourOn(t *testing.T) {
	feedingZone := FeedingZone{}
	zoneAnalyzer := ZoneAnalyzer{}
	feeder := AutomaticFeeder{}

	t.Run("No animals in zone", func(t *testing.T) {
		result := zoneAnalyzer.Analyze(&feedingZone)

		got := feeder.pourOn(result)
		if got != nil {
			t.Errorf("Unexpected data: got=%+v, want=%+v", got, nil)
		}
	})

	t.Run("Animals in zone", func(t *testing.T) {
		bigCow := Animal{
			Name:   "Milka",
			Specie: bull,
		}
		smallCow := Animal{
			Name:   "Milka",
			Specie: bull,
		}
		smallHorse := Animal{
			Name:   "Willy",
			Specie: horse,
		}
		feedingZone.AddAnimal(bigCow)
		feedingZone.AddAnimal(smallCow)
		feedingZone.AddAnimal(smallHorse)

		result := zoneAnalyzer.Analyze(&feedingZone)
		got := feeder.pourOn(result)

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
