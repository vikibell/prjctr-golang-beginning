package main

import (
	"fmt"
	"time"
)

const openRequest = "open"
const closeRequest = "close"

type Logs map[int64]string // Index is timestamp unix nanoseconds

func (logs Logs) printLogs() {
	for t, l := range logs {
		fmt.Printf("%d %s\n", t, l)
	}
}

type Animal struct {
	ID     int
	Health int
	Hunger int
	Mood   int
}

type Enclosure struct {
	ID     int
	IsOpen bool
}

type Request struct {
	EnclosureID int
	Request     string
}

type Feeder struct {
	ID      int
	IsEmpty bool
}

func manageEnclosures(enclosures map[int]Enclosure, requests <-chan Request, logs chan<- string) {
	for request := range requests {
		e, exists := enclosures[request.EnclosureID]
		if exists {
			time.Sleep(1 * time.Second)
			switch request.Request {
			case openRequest:
				if e.IsOpen {
					logs <- fmt.Sprintf("[Warning] Enclosure %d is already opened", e.ID)
				} else {
					e.IsOpen = true
					logs <- fmt.Sprintf("[Info] Enclosure %d was opened", e.ID)
				}
			case closeRequest:
				if !e.IsOpen {
					logs <- fmt.Sprintf("[Warning] Enclosure %d is already closed", e.ID)
				} else {
					e.IsOpen = false
					logs <- fmt.Sprintf("[Info] Enclosure %d was closed", e.ID)
				}
			}
		}
	}
}

func collectState(a Animal, monitorSystem chan<- Animal, logs chan<- string) {
	time.Sleep(5 * time.Second)

	if a.Mood < 50 {
		logs <- fmt.Sprintf("[Info] Animal %d is not happy", a.ID)
	}

	if a.Hunger < 30 {
		logs <- fmt.Sprintf("[Warning] Animal %d is hungry", a.ID)
	}

	if a.Health < 70 {
		logs <- fmt.Sprintf("[Critical] Animal %d health need to be checked", a.ID)
	}

	monitorSystem <- a
}

func checkFeeder(f Feeder, statuses chan<- Feeder, logs chan<- string) {
	time.Sleep(2 * time.Second)

	if f.IsEmpty {
		logs <- fmt.Sprintf("[Info] Feeder %d is empty", f.ID)
	}

	statuses <- f
}

func processLogs(logsCh <-chan string) Logs {
	logs := make(Logs)
	for l := range logsCh {
		logs[time.Now().UnixNano()] = l
	}

	return logs
}
