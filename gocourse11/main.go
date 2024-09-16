package main

import (
	"fmt"
	"gocourse11/filtering"
	"gocourse11/terrarium"
	"math/rand/v2"
	"slices"
)

func generateAquariums(n int, builder terrarium.Builder) []terrarium.Aquarium {
	aquariums := make([]terrarium.Aquarium, n)
	director := terrarium.NewDirector(builder)

	for i := range aquariums {
		size := rand.IntN(100)
		animalName := "name" + string(rune(rand.IntN(1000)))
		level := rand.IntN(100)
		pLevel := rand.IntN(200)
		aquariums[i] = *director.Construct(size, animalName, level, pLevel)
	}

	return aquariums
}

func main() {
	snakeBuilder := terrarium.NewSnakeAquariumBuilder()
	turtleBuilder := terrarium.NewTurtleAquariumBuilder()
	lizardBuilder := terrarium.NewLizardAquariumBuilder()

	snakeA := generateAquariums(2, snakeBuilder)
	turtleA := generateAquariums(3, turtleBuilder)
	lizardA := generateAquariums(5, lizardBuilder)

	aquariums := slices.Concat(snakeA, turtleA, lizardA)
	run(aquariums)
}

func run(aquariums []terrarium.Aquarium) {
	t := terrarium.NewTerrarium()
	t.SetAquariums(aquariums)

	pl := t.CalculatePollutionLevel()
	t.SetFilter(filtering.SelectFilter(pl))

	fmt.Printf("Based on pollution level %d filter with such characteristics was selected:\ncleaner level: %d, absorber type: %s, water improver: %s\n",
		pl, t.Filter().CleanLevel(), t.Filter().Absorber(), t.Filter().WaterImprover())
}
