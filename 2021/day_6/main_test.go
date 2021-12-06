package main

import (
	"fmt"
	"strings"
	"testing"
)

var testInput = strings.TrimSpace(`
3,4,3,1,2
`)

func TestRun(t *testing.T) {
	inp := strings.Split(testInput, "\n")
	p1, p2 := Run(inp)

	if p1 != 5934 {
		t.Errorf("expected %d but got %d", 5934, p1)
	}

	if p2 != 26984457539 {
		t.Errorf("expected %d but got %d", 26984457539, p2)
	}
}

func TestRun2(t *testing.T) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, PreCompute2(i, 256))
	}
}
