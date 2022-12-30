package main

import (
	"testing"
)

var testInput = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

const p1Answer int = 21
const p2Answer int = 8

func TestRun(t *testing.T) {
	p1, p2 := Run(testInput)

	if p1 != p1Answer {
		t.Errorf("expected %d but got %d", p1Answer, p1)
	}

	if p2 != p2Answer {
		t.Errorf("expected %d but got %d", p2Answer, p2)
	}
}
