package terrarium

import (
	"reflect"
	"testing"
)

func TestSnakeAquariumBuilder(t *testing.T) {
	la := NewSnakeAquariumBuilder().
		SetSize(100).
		SetAnimal("Kur").
		SetSaltLevel(10).
		SetPollutionLevel(20).
		Build()

	want := &Aquarium{
		size:           100,
		animalName:     "Snake Kur",
		saltLevel:      10,
		pollutionLevel: 20,
	}

	if got := la; !reflect.DeepEqual(got, want) {
		t.Errorf("SnakeAquariumBuilder.Build(): got = %v, want %v", got, want)
	}
}
