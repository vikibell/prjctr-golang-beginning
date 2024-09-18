package terrarium

import (
	"reflect"
	"testing"
)

func TestLizardAquariumBuilder(t *testing.T) {
	la := NewLizardAquariumBuilder().
		SetSize(100).
		SetAnimal("Kur").
		SetSaltLevel(10).
		SetPollutionLevel(20).
		Build()

	want := &Aquarium{
		size:           100,
		animalName:     "Lizard Kur",
		saltLevel:      10,
		pollutionLevel: 20,
	}

	if got := la; !reflect.DeepEqual(got, want) {
		t.Errorf("LizardAquariumBuilder.Build(): got = %v, want %v", got, want)
	}
}
