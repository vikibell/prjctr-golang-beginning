package filter

import (
	"reflect"
	"testing"
)

func TestNewFilter(t *testing.T) {
	filter := NewFilter(
		WithCleanLevel(Middle),
		WithAbsorber("coal"),
		WithWaterImprover("cn2"),
	)

	want := Filter{
		cleanLevel:    Middle,
		absorber:      "coal",
		waterImprover: "cn2",
	}

	if got := filter; !reflect.DeepEqual(got, want) {
		t.Errorf("NewFilter(): got = %v, want %v", got, want)
	}
}
