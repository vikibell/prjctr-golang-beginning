package main

import (
	"math/rand/v2"
	"sync"
	"testing"
	"time"
)

func TestRaceConditions(t *testing.T) {
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

	processLogs(logs)
}

func TestDeadlocks(t *testing.T) {
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

	done := make(chan bool)
	go func() {
		wg.Wait()
		close(monitorSystem)
		close(feedersCh)
		close(logs)
		done <- true
	}()

	processLogs(logs)

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Test timed out, possible deadlock detected")
	}
}
