package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
`)
var testInput2 = strings.TrimSpace(`
19999
19111
11191
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 40 {
		t.Errorf("expected %d but got %d", 40, p1)
	}

	if p2 != 315 {
		t.Errorf("expected %d but got %d", 315, p2)
	}
}
