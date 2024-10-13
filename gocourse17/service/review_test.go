package service

import (
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewReview(t *testing.T) {
	cs := int32(2)
	sq := int32(3)
	fs := int32(4)

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

func TestIsValidRating(t *testing.T) {
	cs := int32(2)
	sq := int32(6)

	t.Run("Valid", func(t *testing.T) {
		if got := IsValidRating(cs); !cmp.Equal(got, true) {
			t.Errorf("NewReview(): got = %v, want %v", got, true)
		}
	})
	t.Run("Not valid", func(t *testing.T) {
		if got := IsValidRating(sq); !cmp.Equal(got, false) {
			t.Errorf("NewReview(): got = %v, want %v", got, false)
		}
	})
}

func TestReviewHistory_AddGet(t *testing.T) {
	cs := int32(2)
	sq := int32(3)
	fs := int32(4)
	driverId := int32(1)

	review := NewReview(cs, sq, fs)
	history := NewReviewHistory()
	history.AddReview(driverId, review)

	got, _ := history.GerReviews(driverId)

	want := []Review{
		{
			CargoState:       cs,
			ServiceQuality:   sq,
			FulfillmentSpeed: fs,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReviewHistory.Add(): got = %v, want %v", got, want)
	}
}
