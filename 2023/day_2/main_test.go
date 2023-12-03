package main

import (
	"testing"
)

const p1Answer int = 8
const p2Answer int = 2286

var testInput1 = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
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
