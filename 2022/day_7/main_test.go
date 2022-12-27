package main

import (
	"testing"
)

var testInput = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

const p1Answer int = 95437
const p2Answer int = 24933642

func TestRun(t *testing.T) {
	p1, p2 := Run(testInput)

	if p1 != p1Answer {
		t.Errorf("expected %d but got %d", p1Answer, p1)
	}

	if p2 != p2Answer {
		t.Errorf("expected %d but got %d", p2Answer, p2)
	}
}
