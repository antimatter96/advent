package main

import (
	"testing"
)

var testInput = []string{
	"2-4,6-8",
	"2-3,4-5",
	"5-7,7-9",
	"2-8,3-7",
	"6-6,4-6",
	"2-6,4-8",
}

const p1Answer int = 2
const p2Answer int = 4

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
