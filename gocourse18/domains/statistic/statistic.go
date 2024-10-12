package statistic

const (
	Child    = iota + 1 // 1-8 years
	Teenager            // 9-21 years
	Adult               // 22-40 years
	Elderly             // 41-60 years
	OldMan              // 61-100 years
)

type Statistic struct {
	ID           uint
	City         string
	AverageTrips int64 `db:"average_trips"`
	AgeRange     rune  `db:"age_range"`
}
