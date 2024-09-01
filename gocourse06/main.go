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
func generateEnclosures(n int) map[int]Enclosure {
	enclosures := make(map[int]Enclosure)
	for i := 0; i < n; i++ {
		enclosures[i] = Enclosure{
			ID:     i,
			IsOpen: rand.IntN(2) == 1,
		}
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
		go func() {
			collectState(animal, monitorSystem, logs)
			wg.Done()
		}()
	}

	feedersCh := make(chan Feeder, len(feeders))
	for _, feeder := range feeders {
		wg.Add(1)
		go func() {
			checkFeeder(feeder, feedersCh, logs)
			wg.Done()
		}()
	}

	requests := [2]string{openRequest, closeRequest}
	requestsCh := make(chan Request, len(enclosures))
	for _, enclosure := range enclosures {
		requestsCh <- Request{EnclosureID: enclosure.ID, Request: requests[rand.IntN(len(requests))]}
	}
	close(requestsCh)

	wg.Add(1)
	go func() {
		manageEnclosures(enclosures, requestsCh, logs)
		wg.Done()
	}()

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
