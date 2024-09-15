package zone

import "testing"

func TestFeedingZone_GetAnimals(t *testing.T) {
	bigCow := Animal{
		Name:   "Milka",
		Specie: Bull,
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

func TestFeedingZone_AddAnimal(t *testing.T) {
	bigCow := Animal{
		Name:   "Milka",
		Specie: Bull,
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
