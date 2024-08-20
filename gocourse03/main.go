package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	rodents := []Rodent{
		NewRodent(11, RodentRat, FromTo{}, NewDailyMovements()),
		NewRodent(12, RodentRat, FromTo{}, NewDailyMovements()),
		NewRodent(13, RodentRat, FromTo{}, NewDailyMovements()),
	}

	startMovement := time.Now()
	for _, rodent := range rodents {
		sector := chooseRandomSector()
		rodent.Movements.AddMovement(NewMovement(chooseRandomSector(), SectorCenter, startMovement))
		rodent.Movements.AddMovement(NewMovement(SectorCenter, sector, startMovement.Local().Add(time.Minute*time.Duration(2))))
		rodent.Movements.AddMovement(NewMovement(sector, SectorCenter, startMovement.Local().Add(time.Minute*time.Duration(6))))
		rodent.Movements.AddMovement(NewMovement(SectorCenter, chooseRandomSector(), startMovement.Local().Add(time.Minute*time.Duration(11))))
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

func chooseRandomSector() Sector {
	sectors := [5]Sector{
		"Sector A", "Sector B", "Sector C", "Sector D", "Sector E",
	}

	return sectors[rand.IntN(len(sectors))]
}
