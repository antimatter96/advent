package main

import (
	"fmt"
	"strings"
)

type File struct {
	name string
	size int
}

type Folder struct {
	name    string
	folders []*Folder
	files   []*File
}

func (folder *Folder) Size() int {
	sum := 0
	for _, file := range folder.files {
		sum += file.size
	}

	for _, folder := range folder.folders {
		sum += folder.Size()
	}
	return sum
}

func (folder *Folder) Print(level int) {
	ls := strings.Repeat("  ", level-1)
	fmt.Println(ls, "[d]", folder.name)
	ls = strings.Repeat("  ", level)
	for _, f := range folder.files {
		fmt.Println(ls, "[f]", f.name)
	}
	for _, f := range folder.folders {
		f.Print(level + 1)
	}
}
