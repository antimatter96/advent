package main

import (
	"testing"
)

const p1Answer int = 3
const p2Answer int = 6

var testInput1 = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

var testInput2 = []string{
	"R1000",
}

func TestRun(t *testing.T) {
	p1, _ := Run(testInput1)

	if p1 != p1Answer {
		t.Errorf("expected %d but got %d", p1Answer, p1)
	}

	_, p2 := Run(testInput1)
	if p2 != p2Answer {
		t.Errorf("expected %d but got %d", p2Answer, p2)
	}
}
