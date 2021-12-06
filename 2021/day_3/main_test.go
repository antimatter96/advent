package main

import (
	"testing"
)

var testInput = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestRun(t *testing.T) {
	p1, p2 := Run(testInput)

	if p1 != 198 {
		t.Errorf("expected %d but got %d", 198, p1)
	}

	if p2 != 230 {
		t.Errorf("expected %d but got %d", 230, p2)
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
