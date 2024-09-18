package terrarium

import (
	"reflect"
	"testing"
)

func TestDirector_Construct(t *testing.T) {
	builder := NewSnakeAquariumBuilder()
	director := NewDirector(builder)

	aquarium := director.Construct(10, "Test", 12, 50)

	want := &Aquarium{
		size:           10,
		animalName:     "Snake Test",
		saltLevel:      12,
		pollutionLevel: 50,
	}

	if got := aquarium; !reflect.DeepEqual(got, want) {
		t.Errorf("Director.Construct(): got = %v, want %v", got, want)
	}
}
