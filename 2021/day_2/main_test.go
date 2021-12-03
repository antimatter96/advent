package main

import (
	"testing"
)

var testInput = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2",
}

func TestRun(t *testing.T) {
	p1, p2 := Run(testInput)

	if p1 != 150 {
		t.Errorf("expected %d but got %d", 150, p1)
	}

	if p2 != 900 {
		t.Errorf("expected %d but got %d", 900, p2)
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
