package main

import (
	"testing"
)

var testInput = []string{
	"199",
	"200",
	"208",
	"210",
	"200",
	"207",
	"240",
	"269",
	"260",
	"263",
}

func TestRun(t *testing.T) {
	p1, p2 := Run(testInput)

	if p1 != 7 {
		t.Errorf("expected %q but got %q", 7, p1)
	}

	if p2 != 5 {
		t.Errorf("expected %q but got %q", 5, p2)
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

var parsedInputForPart2 = parsePart2(testInput)

func BenchmarkPart2(b *testing.B) {

	implementations := []struct {
		name string
		fn   func([]int) int
	}{
		{"Original", Part2},
		{"Smart", Part2_1},
		{"Dumb", Part2_2},
	}

	var x int
	for _, implementation := range implementations {
		b.Run(implementation.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x = implementation.fn(parsedInputForPart2)

			}
		})
	}

	temp1 = x
}
