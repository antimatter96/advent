package main

import (
	"testing"
)

const p1Answer int = 480
const p2Answer int = 875318608908

var testInput1 = []string{
	"Button A: X+94, Y+34",
	"Button B: X+22, Y+67",
	"Prize: X=8400, Y=5400",
	"",
	"Button A: X+26, Y+66",
	"Button B: X+67, Y+21",
	"Prize: X=12748, Y=12176",
	"",
	"Button A: X+17, Y+86",
	"Button B: X+84, Y+37",
	"Prize: X=7870, Y=6450",
	"",
	"Button A: X+69, Y+23",
	"Button B: X+27, Y+71",
	"Prize: X=18641, Y=10279",
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
