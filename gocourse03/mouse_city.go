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

	DailyMovements []Movement

	Movement struct {
		Time time.Time
		FromTo
	}

	Rodent struct {
		ID        int
		Type      RodentType
		History   FromTo
		Movements DailyMovements
	}
)

func NewRodent(id int, rodentType RodentType, history FromTo, movements DailyMovements) Rodent {
	return Rodent{
		ID:        id,
		Type:      rodentType,
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

func (dm DailyMovements) PrintRodentMovements() {
	fmt.Println("Rodent recorded movements:")
	for _, movement := range dm {
		fmt.Printf("%s: %s - %s\n", movement.Time.Format("15:04:01"), movement.FromTo[0], movement.FromTo[1])
	}
}
