package main

import (
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
16,1,2,0,4,2,7,1,2,14
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 37 {
		t.Errorf("expected %d but got %d", 37, p1)
	}

	if p2 != 168 {
		t.Errorf("expected %d but got %d", 168, p2)
	}
}
