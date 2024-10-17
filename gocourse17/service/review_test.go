package service

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewReview(t *testing.T) {
	cs := Rating(2)
	sq := Rating(3)
	fs := Rating(4)

	review := NewReview(cs, sq, fs)
	want := Review{
		CargoState:       cs,
		ServiceQuality:   sq,
		FulfillmentSpeed: fs,
	}

	if got := review; !reflect.DeepEqual(got, want) {
		t.Errorf("NewReview(): got = %v, want %v", got, want)
	}
}

func TestIsValid(t *testing.T) {
	cs := Rating(2)
	sq := Rating(6)

	t.Run("Valid", func(t *testing.T) {
		if got := IsValid(cs); !cmp.Equal(got, true) {
			t.Errorf("NewReview(): got = %v, want %v", got, true)
		}
	})
	t.Run("Not valid", func(t *testing.T) {
		if got := IsValid(sq); !cmp.Equal(got, false) {
			t.Errorf("NewReview(): got = %v, want %v", got, false)
		}
	})
}

func TestReviewHistory_GetReviews(t *testing.T) {
	cs := Rating(2)
	sq := Rating(3)
	fs := Rating(4)
	driverID := 1

	review := NewReview(cs, sq, fs)
	history := NewReviewHistory()
	history.AddReview(driverID, review)
	history.AddReview(driverID, review)

	got := history.GetReviews(driverID)

	want := []Review{
		{
			CargoState:       cs,
			ServiceQuality:   sq,
			FulfillmentSpeed: fs,
		},
		{
			CargoState:       cs,
			ServiceQuality:   sq,
			FulfillmentSpeed: fs,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReviewHistory.GetReviews(): got = %v, want %v", got, want)
	}
}

func TestReviewHistory_AddReview(t *testing.T) {
	cs := Rating(2)
	sq := Rating(3)
	fs := Rating(4)
	driverID := 1

	review := NewReview(cs, sq, fs)
	history := NewReviewHistory()
	history.AddReview(driverID, review)

	got := history.GetReviews(driverID)[0]
	want := Review{
		CargoState:       cs,
		ServiceQuality:   sq,
		FulfillmentSpeed: fs,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReviewHistory.Add(): got = %v, want %v", got, want)
	}
}
