package main

import "time"

const RodentRat rodentType = "Rat"

type (
	Sector     string
	rodentType string
	FromTo     [2]Sector

	//SectorHistory

	DailyMovements struct {
		RodentMovement []Movement
	}

	Movement struct {
		Time time.Time
		FromTo
	}

	Rodent struct {
		ID        int
		Type      rodentType
		History   FromTo //find first and last sector in daily movements
		Movements *DailyMovements
	}
)

func NewMovement(from, to Sector) Movement {
	return Movement{
		time.Now(),
		FromTo{
			from,
			to,
		},
	}
}

func NewDailyMovements() *DailyMovements {
	return &DailyMovements{
		RodentMovement: make([]Movement, 0),
	}
}

func (dm *DailyMovements) addMovement(movement Movement) {
	dm.RodentMovement = append(dm.RodentMovement, movement)
}
