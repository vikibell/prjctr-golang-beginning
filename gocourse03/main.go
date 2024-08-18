package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rodents := []Rodent{
		{ID: 11, Type: RodentRat, History: FromTo{}, Movements: NewDailyMovements()},
		{ID: 12, Type: RodentRat, History: FromTo{}, Movements: NewDailyMovements()},
		{ID: 13, Type: RodentRat, History: FromTo{}, Movements: NewDailyMovements()},
	}

	for _, rodent := range rodents {
		sector := getRandomSector()
		rodent.Movements.AddMovement(NewMovement(getRandomSector(), SectorCenter))
		rodent.Movements.AddMovement(NewMovement(SectorCenter, sector))
		rodent.Movements.AddMovement(NewMovement(sector, SectorCenter))
		rodent.Movements.AddMovement(NewMovement(SectorCenter, getRandomSector()))
	}

	for _, rodent := range rodents {
		firstMovement := rodent.Movements.FindMovementByKey(0)
		lastMovement := rodent.Movements.FindMovementByKey(len(rodent.Movements.RodentMovement) - 1)
		rodent.History = FromTo{
			firstMovement.FromTo[0],
			lastMovement.FromTo[1],
		}
		fmt.Printf(
			"Rodent number: %d. Start of the day: %s. End of the day: %s\n",
			rodent.ID, rodent.History[0], rodent.History[1],
		)
		rodent.Movements.PrintRodentMovements()
	}
}

func getRandomSector() Sector {
	sectors := [5]Sector{
		"Sector A", "Sector B", "Sector C", "Sector D", "Sector E",
	}

	return sectors[rand.Intn(len(sectors))]
}
