package main

import (
	"fmt"
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`)

var testInput2 = strings.TrimSpace(`
11111
19991
19191
19991
11111
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 26397 {
		t.Errorf("expected %d but got %d", 26397, p1)
	}

	if p2 != 288957 {
		t.Errorf("expected %d but got %d", 288957, p2)
	}
}
func TestRun2(t *testing.T) {
	inp := strings.Split(testInput2, "\n")
	p1, p2 := Run(inp)

	fmt.Println(p1, p2)
}
