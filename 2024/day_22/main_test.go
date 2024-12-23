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

func TestMix(t *testing.T) {
	out := mix(42, 15)
	if out != 37 {
		t.Errorf("expected %d but got %d", 37, out)
	}
}

func TestPrune(t *testing.T) {
	out := prune(100000000)
	if out != 16113920 {
		t.Errorf("expected %d but got %d", 16113920, out)
	}
}

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
