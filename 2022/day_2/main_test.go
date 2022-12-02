package main

import (
	"testing"
)

const p1Answer int = 15
const p2Answer int = 12

var testInput = []string{
	"A Y",
	"B X",
	"C Z",
}

func TestRun(t *testing.T) {
	p1, p2 := Run(testInput)

	if p1 != p1Answer {
		t.Errorf("expected %d but got %d", p1Answer, p1)
	}

	if p2 != p2Answer {
		t.Errorf("expected %d but got %d", p2Answer, p2)
	}
}

var temp1, temp2 int

func BenchmarkRun(b *testing.B) {
	var x, y int
	for i := 0; i < b.N; i++ {
		x, y = Run(testInput)
	}
	temp1 = x
	temp2 = y
}
