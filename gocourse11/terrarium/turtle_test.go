package terrarium

import (
	"reflect"
	"testing"
)

func TestTurtleAquariumBuilder(t *testing.T) {
	la := NewTurtleAquariumBuilder().
		SetSize(100).
		SetAnimal("Kur").
		SetSaltLevel(10).
		SetPollutionLevel(20).
		Build()

	want := &Aquarium{
		size:           100,
		animalName:     "Turtle Kur",
		saltLevel:      10,
		pollutionLevel: 20,
	}

	if got := la; !reflect.DeepEqual(got, want) {
		t.Errorf("TurtleAquariumBuilder.Build(): got = %v, want %v", got, want)
	}
}
