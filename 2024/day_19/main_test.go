package main

import (
	"testing"
)

const p1Answer int = 6
const p2Answer int = 16

var testInput1 = []string{
	"r, wr, b, g, bwu, rb, gb, br",
	"",
	"bggr",
	"brgr",
	"gbbr",
	"ubwu",
	"brwrr",
	"bbrgwb",
	"bwurrg",
	"rrbgbr",
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
