package main

import (
	"fmt"
	"time"
)

const RodentRat RodentType = "Rat"

const SectorCenter Sector = "Sector Center"

type (
	Sector     string
	RodentType string
	FromTo     [2]Sector

	DailyMovements struct {
		RodentMovement []Movement
	}

	Movement struct {
		Time time.Time
		FromTo
	}

	Rodent struct {
		ID        int
		Type      RodentType
		History   FromTo
		Movements *DailyMovements
	}
)

func NewRodent(id int, Type RodentType, history FromTo, movements *DailyMovements) Rodent {
	return Rodent{
		ID:        id,
		Type:      Type,
		History:   history,
		Movements: movements,
	}
}

func NewMovement(from, to Sector, time time.Time) Movement {
	return Movement{
		time,
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

func (dm *DailyMovements) AddMovement(movement Movement) {
	dm.RodentMovement = append(dm.RodentMovement, movement)
}

func (dm *DailyMovements) RemoveMovement(index int) {
	if index < 0 || index >= len(dm.RodentMovement) {
		return
	}
	dm.RodentMovement = append(dm.RodentMovement[:index], dm.RodentMovement[index+1:]...)
}

func (dm *DailyMovements) FindMovementByKey(index int) Movement {
	return dm.RodentMovement[index]
}

func (dm *DailyMovements) PrintRodentMovements() {
	fmt.Println("Rodent recorded movements:")
	for _, movement := range dm.RodentMovement {
		fmt.Printf("%s: %s - %s\n", movement.Time.Format("15:04:01"), movement.FromTo[0], movement.FromTo[1])
	}
}
