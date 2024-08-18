package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	startDayTime := time.Now()

	fmt.Println(startDayTime)
	rodents := []Rodent{
		{ID: 11, Type: RodentRat, Movements: NewDailyMovements()},
		{ID: 12, Type: RodentRat, Movements: NewDailyMovements()},
		{ID: 13, Type: RodentRat, Movements: NewDailyMovements()},
	}

	for _, rodent := range rodents {
		rodentMovements := rodent.Movements
		rodentMovements.addMovement(NewMovement(getRandomSector(), getRandomSector()))
	}
}

func getRandomSector() Sector {
	sectors := [5]Sector{
		"Sector A", "Sector B", "Sector C", "Sector D", "Sector E",
	}

	return sectors[rand.Intn(len(sectors))]
}
