package terrarium

import (
	"reflect"
	"testing"
)

func TestTerrarium_CalculatePollutionLevel(t *testing.T) {
	tr := NewTerrarium()
	tr.AddAquarium(*NewLizardAquariumBuilder().
		SetSize(100).
		SetAnimal("Kur").
		SetSaltLevel(10).
		SetPollutionLevel(20).
		Build(),
	)
	tr.AddAquarium(*NewSnakeAquariumBuilder().
		SetSize(100).
		SetAnimal("Kur").
		SetSaltLevel(10).
		SetPollutionLevel(60).
		Build(),
	)
	tr.AddAquarium(*NewTurtleAquariumBuilder().
		SetSize(100).
		SetAnimal("Kur").
		SetSaltLevel(10).
		SetPollutionLevel(20).
		Build(),
	)

	pl := tr.CalculatePollutionLevel()
	want := 100

	if got := pl; !reflect.DeepEqual(got, want) {
		t.Errorf("Terrarium.CalculatePollutionLevel(): got = %v, want %v", got, want)
	}
}
