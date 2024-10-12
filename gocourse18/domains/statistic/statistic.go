package statistic

const (
	Child = iota + 1
	Teenager
	Adult
	Elderly
	OldMan
)

type Statistic struct {
	ID           uint
	City         string
	AgeRange     rune
	AverageTrips int64
}
