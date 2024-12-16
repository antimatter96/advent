package main

import (
	"testing"
)

const p1Answer int = 12
const p2Answer int = 0

var testInput1 = []string{
	"11 7",
	"p=0,4 v=3,-3",
	"p=6,3 v=-1,-3",
	"p=10,3 v=-1,2",
	"p=2,0 v=2,-1",
	"p=0,0 v=1,3",
	"p=3,0 v=-2,-2",
	"p=7,6 v=-1,-3",
	"p=3,0 v=-1,-2",
	"p=9,3 v=2,3",
	"p=7,3 v=-1,2",
	"p=2,4 v=2,-3",
	"p=9,5 v=-3,-3",
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
