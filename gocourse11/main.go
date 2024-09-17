package main

import (
	"fmt"
	"github.com/vikibell/prjctr-golang-beginning/gocourse11/filter"
	"github.com/vikibell/prjctr-golang-beginning/gocourse11/terrarium"
	"math/rand/v2"
	"slices"
)

func generateAquariums(n int, builder terrarium.Builder) []terrarium.Aquarium {
	aquariums := make([]terrarium.Aquarium, n)
	director := terrarium.NewDirector(builder)

	for i := range aquariums {
		size := rand.IntN(100)
		animalName := "name" + string(rune(rand.IntN(1000)))
		saltLevel := rand.IntN(100)
		pollutionLevel := rand.IntN(200)
		aquariums[i] = *director.Construct(size, animalName, saltLevel, pollutionLevel)
	}

	return aquariums
}

func main() {
	snakeBuilder := terrarium.NewSnakeAquariumBuilder()
	turtleBuilder := terrarium.NewTurtleAquariumBuilder()
	lizardBuilder := terrarium.NewLizardAquariumBuilder()

	snakeAquariums := generateAquariums(2, snakeBuilder)
	turtleAquariums := generateAquariums(3, turtleBuilder)
	lizardAquariums := generateAquariums(5, lizardBuilder)

	aquariums := slices.Concat(snakeAquariums, turtleAquariums, lizardAquariums)
	t := terrarium.NewTerrarium()
	t.SetAquariums(aquariums)

	pl := t.CalculatePollutionLevel()
	t.SetFilter(filter.Select(pl))

	fmt.Printf("Based on pollution level %d filter with such characteristics was selected:\ncleaner level: %d, absorber type: %s, water improver: %s\n",
		pl, t.Filter().CleanLevel(), t.Filter().Absorber(), t.Filter().WaterImprover())
}
