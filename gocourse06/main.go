package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

// Генерує тестові дані для тварин
func generateAnimals(n int) []Animal {
	var animals []Animal
	for i := 0; i < n; i++ {
		animal := Animal{
			ID:     i,
			Health: rand.IntN(100),
			Hunger: rand.IntN(100),
			Mood:   rand.IntN(100),
		}
		animals = append(animals, animal)
	}
	return animals
}

// Генерує тестові дані для вольєрів
func generateEnclosures(n int) []Enclosure {
	var enclosures []Enclosure
	for i := 0; i < n; i++ {
		enclosure := Enclosure{
			ID:     i,
			IsOpen: rand.IntN(2) == 1,
		}
		enclosures = append(enclosures, enclosure)
	}
	return enclosures
}

// Генерує тестові дані для кормушок
func generateFeeders(n int) []Feeder {
	var feeders []Feeder
	for i := 0; i < n; i++ {
		feeder := Feeder{
			ID:      i,
			IsEmpty: rand.IntN(2) == 1,
		}
		feeders = append(feeders, feeder)
	}
	return feeders
}

func main() {
	animals := generateAnimals(8)
	feeders := generateFeeders(5)
	enclosures := generateEnclosures(5)

	var wg sync.WaitGroup

	logs := make(chan string)

	monitorSystem := make(chan Animal, len(animals))
	for _, animal := range animals {
		wg.Add(1)
		go collectState(&wg, animal, monitorSystem, logs)
	}

	feedersCh := make(chan Feeder, len(feeders))
	for _, feeder := range feeders {
		wg.Add(1)
		go checkFeeder(&wg, feeder, feedersCh, logs)
	}

	requests := [2]string{openRequest, closeRequest}
	requestsCh := make(chan Request, len(enclosures))
	for _, enclosure := range enclosures {
		requestsCh <- Request{EnclosureID: enclosure.ID, Request: requests[rand.IntN(len(requests))]}
	}
	close(requestsCh)

	for _, enclosure := range enclosures {
		wg.Add(1)
		go manageEnclosures(&wg, &enclosure, requestsCh, logs)
	}

	go func() {
		wg.Wait()
		close(monitorSystem)
		close(feedersCh)
		close(logs)
	}()

	processedLogs := processLogs(logs)

	for animal := range monitorSystem {
		fmt.Printf("Animal statistics %d: \n Health:%d;\n Mood: %d\n Hunger:%d\n", animal.ID, animal.Health, animal.Mood, animal.Hunger)
	}

	for feeder := range feedersCh {
		if feeder.IsEmpty {
			fmt.Printf("Feeder %d is empty.\n", feeder.ID)
		} else {
			fmt.Printf("Feeder %d is full.\n", feeder.ID)
		}
	}

	processedLogs.printLogs()
}
