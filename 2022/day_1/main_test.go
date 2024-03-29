package main

import (
	"testing"
)

const p1Answer int = 24000
const p2Answer int = 45000

var testInput = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestRun(t *testing.T) {
	p1, p2 := Run(testInput)

	if p1 != p1Answer {
		t.Errorf("expected %d but got %d", p1Answer, p1)
	}

	if p2 != p2Answer {
		t.Errorf("expected %q but got %q", p2Answer, p2)
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

	// implementations := []struct {
	// 	name string
	// 	fn   func([]int) int
	// }{
	// 	{"Original", Part2},
	// 	{"Smart", Part2_1},
	// 	{"Dumb", Part2_2},
	// }

	// var x int
	// for _, implementation := range implementations {
	// 	b.Run(implementation.name, func(b *testing.B) {
	// 		for i := 0; i < b.N; i++ {
	// 			x = implementation.fn(parsedInputForPart2)

	// 		}
	// 	})
	// }

	// temp1 = x
}
