package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`)

var testInput2 = strings.TrimSpace(`
1,1 -> 3,3
9,7 -> 7,9
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 5 {
		t.Errorf("expected %d but got %d", 5, p1)
	}

	if p2 != 12 {
		t.Errorf("expected %d but got %d", 12, p2)
	}
}
