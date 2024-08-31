package main

import (
	"fmt"
	"sync"
	"time"
)

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

type Feeder struct {
	ID      int
	IsEmpty bool
}

//func manageEnclosures(requests chan<-) {
//
//}

func collectState(wg *sync.WaitGroup, a Animal, monitorSystem chan<- Animal, logs chan<- string) {
	defer wg.Done()
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

func checkFeeder(wg *sync.WaitGroup, f Feeder, statuses chan<- Feeder, logs chan<- string) {
	defer wg.Done()
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
