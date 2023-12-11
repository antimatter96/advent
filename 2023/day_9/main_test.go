package main

import (
	"testing"
)

const p1Answer int = 114
const p2Answer int = 2

var testInput1 = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
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
