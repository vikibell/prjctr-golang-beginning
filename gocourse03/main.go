package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"time"
)

func main() {
	rodents := []Rodent{
		NewRodent(11, RodentRat, FromTo{}, NewDailyMovements()),
		NewRodent(12, RodentRat, FromTo{}, NewDailyMovements()),
		NewRodent(13, RodentRat, FromTo{}, NewDailyMovements()),
		NewRodent(66, "To delete type", FromTo{}, NewDailyMovements()),
	}

	rodents = slices.Delete(rodents, 3, 4)

	startMovement := time.Now()
	for key, _ := range rodents {
		sector := chooseRandomSector()
		rodents[key].Movements = append(
			rodents[key].Movements,
			NewMovement(chooseRandomSector(), SectorCenter, startMovement),
			NewMovement(SectorCenter, sector, startMovement.Local().Add(time.Minute*time.Duration(2+rand.IntN(10)))),
			NewMovement(sector, SectorCenter, startMovement.Local().Add(time.Minute*time.Duration(6+rand.IntN(10)))),
			NewMovement(SectorCenter, chooseRandomSector(), startMovement.Local().Add(time.Minute*time.Duration(11+rand.IntN(10)))),
		)
	}

	for key, rodent := range rodents {
		firstMovement := rodent.Movements[0]
		lastMovement := rodent.Movements[len(rodent.Movements)-1]
		rodents[key].History = FromTo{
			firstMovement.FromTo[0],
			lastMovement.FromTo[1],
		}
	}

	for _, rodent := range rodents {
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
