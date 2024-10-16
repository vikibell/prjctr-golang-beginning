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

func IsValidRating(rating int32) bool {
	return rating >= 1 && rating <= 5
}

type ReviewHistory map[int32][]Review

func NewReviewHistory() ReviewHistory {
	return make(ReviewHistory)
}

func (rh ReviewHistory) AddReview(driverID int32, review Review) {
	rh[driverID] = append(rh[driverID], review)
}

func (rh ReviewHistory) GerReviews(driverID int32) ([]Review, bool) {
	a, exists := rh[driverID]
	return a, exists
}
