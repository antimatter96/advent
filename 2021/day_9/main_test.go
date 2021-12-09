package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
2199943210
3987894921
9856789892
8767896789
9899965678
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 15 {
		t.Errorf("expected %d but got %d", 15, p1)
	}

	if p2 != 1134 {
		t.Errorf("expected %d but got %d", 1134, p2)
	}
}
