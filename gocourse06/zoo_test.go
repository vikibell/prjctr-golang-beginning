package main

import (
	"testing"
	"time"
)

func TestRaceConditions_run(t *testing.T) {
	run()
}

func TestDeadlocks_run(t *testing.T) {
	done := make(chan bool)
	go func() {
		run()
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("Test timed out, possible deadlock detected")
	}
}
