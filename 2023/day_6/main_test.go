package main

import (
	"testing"
)

const p1Answer int64 = 288
const p2Answer int64 = 71503

var testInput1 = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

var testInput2 = testInput1

func TestRun(t *testing.T) {
	p1, _ := Run(testInput1)

	if p1 != p1Answer {
		t.Errorf("expected %d but got %d", p1Answer, p1)
	}

	_, p2 := Run(testInput2)
	if p2 != p2Answer {
		t.Errorf("expected %d but got %d", p2Answer, p2)
	}
}
