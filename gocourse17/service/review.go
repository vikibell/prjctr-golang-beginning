package service

type Rating int

const (
	RatingPoor Rating = iota + 1
	RatingFair
	RatingGood
	RatingGreat
	RatingExcellent
)

type Review struct {
	CargoState       Rating
	ServiceQuality   Rating
	FulfillmentSpeed Rating
}

func NewReview(cs, sq, fs Rating) Review {
	return Review{
		CargoState:       cs,
		ServiceQuality:   sq,
		FulfillmentSpeed: fs,
	}
}

func IsValid(rating Rating) bool {
	return rating >= RatingPoor && rating <= RatingExcellent
}

type ReviewHistory map[int][]Review

func NewReviewHistory() ReviewHistory {
	return make(ReviewHistory)
}

func (rh ReviewHistory) AddReview(driverID int, review Review) {
	rh[driverID] = append(rh[driverID], review)
}

func (rh ReviewHistory) GetReviews(driverID int) []Review {
	return rh[driverID]
}
