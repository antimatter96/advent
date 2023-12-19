package main

import (
	"testing"
)

const p1Answer int = 62
const p2Answer int = 952408144115

var testInput1 = []string{
	"R 6 (#70c710)",
	"D 5 (#0dc571)",
	"L 2 (#5713f0)",
	"D 2 (#d2c081)",
	"R 2 (#59c680)",
	"D 2 (#411b91)",
	"L 5 (#8ceee2)",
	"U 2 (#caa173)",
	"L 1 (#1b58a2)",
	"U 2 (#caa171)",
	"R 2 (#7807d2)",
	"U 3 (#a77fa3)",
	"L 2 (#015232)",
	"U 2 (#7a21e3)",
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
