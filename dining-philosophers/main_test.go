package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()
		if len(orderFinished) != 5 {
			t.Errorf("incorrect length of slice; expected 5, but got %v", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTest = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", 0 * time.Second},
		{"quarter second delay", 250 * time.Millisecond},
		{"half second delay", 500 * time.Millisecond},
	}

	for _, e := range theTest {
		orderFinished = []string{}

		eatTime = e.delay
		thinkTime = e.delay
		sleepTime = e.delay

		dine()
		if len(orderFinished) != 5 {
			t.Errorf("%s: incorrect length of slice; expected 5, but got %v", e.name, len(orderFinished))
		}
	}
}
