package main

import (
	"testing"
)

const p1Answer int = 2
const p2Answer int = 6

var testInput1 = []string{
	"RL",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

var testInput2 = []string{
	"LR",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func TestRun(t *testing.T) {
	p1, _ := Run(testInput1, 1)

	if p1 != p1Answer {
		t.Errorf("expected %d but got %d", p1Answer, p1)
	}

	_, p2 := Run(testInput2, 2)
	if p2 != p2Answer {
		t.Errorf("expected %d but got %d", p2Answer, p2)
	}
}
