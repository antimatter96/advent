package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 17 {
		t.Errorf("expected %d but got %d", 17, p1)
	}

	if p2 != 16 {
		t.Errorf("expected %d but got %d", 16, p2)
	}
}
