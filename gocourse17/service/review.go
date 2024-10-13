package service

type Review struct {
	CargoState       int32
	ServiceQuality   int32
	FulfillmentSpeed int32
}

func NewReview(cs, sq, fs int32) Review {
	return Review{
		CargoState:       cs,
		ServiceQuality:   sq,
		FulfillmentSpeed: fs,
	}
}

type ReviewHistory map[int32][]Review

func NewReviewHistory() ReviewHistory {
	return make(ReviewHistory)
}

func (rh ReviewHistory) AddReview(driverId int32, review Review) {
	rh[driverId] = append(rh[driverId], review)
}

func (rh ReviewHistory) GerReviews(driverId int32) ([]Review, bool) {
	a, exists := rh[driverId]
	return a, exists
}
