package service

const (
	poor int = iota + 1
	fair
	good
	great
	excellent
)

type Review struct {
	CargoState       int
	ServiceQuality   int
	FulfillmentSpeed int
}

func NewReview(cs, sq, fs int) Review {
	return Review{
		CargoState:       cs,
		ServiceQuality:   sq,
		FulfillmentSpeed: fs,
	}
}

func IsValidRating(rating int) bool {
	return rating >= poor && rating <= excellent
}

type ReviewHistory map[int][]Review

func NewReviewHistory() ReviewHistory {
	return make(ReviewHistory)
}

func (rh ReviewHistory) AddReview(driverID int, review Review) {
	rh[driverID] = append(rh[driverID], review)
}

func (rh ReviewHistory) GerReviews(driverID int) ([]Review, bool) {
	a, exists := rh[driverID]
	return a, exists
}
