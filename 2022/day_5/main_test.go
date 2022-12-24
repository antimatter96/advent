package main

import (
	"testing"
)

const p1Answer string = "CMZ"
const p2Answer string = "MCD"

var testInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestRun(t *testing.T) {
	p1, p2 := Run(testInput)

	if p1 != p1Answer {
		t.Errorf("expected %s but got %s", p1Answer, p1)
	}

	if p2 != p2Answer {
		t.Errorf("expected %s but got %s", p2Answer, p2)
	}
}

var temp1, temp2 string

func BenchmarkRun(b *testing.B) {
	var x, y string
	for i := 0; i < b.N; i++ {
		x, y = Run(testInput)
	}
	temp1 = x
	temp2 = y
}
