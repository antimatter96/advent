package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2022/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	fmt.Println(Run(rawInput))
}

func parsePart1(inp []string) []string {
	return inp
}

func parsePart2(inp []string) []string {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

const rootDirectory = "/"

func Part1(logLines []string) int {
	lookup := parseLogLines(logLines)
	sum := 0

	for _, ptr := range lookup {
		sz := ptr.Size()
		if sz <= 100000 {
			sum += sz
		}
	}

	return sum
}

func Part2(logLines []string) int {
	lookup := parseLogLines(logLines)

	currentSize := lookup[rootDirectory].Size()
	freeSpace := 70000000 - currentSize
	requiredToDelete := 30000000 - freeSpace
	minToDel := 70000000

	for _, ptr := range lookup {
		sz := ptr.Size()

		if sz >= requiredToDelete {
			if sz < minToDel {
				minToDel = sz
			}
		}
	}

	return minToDel
}

func parseLogLines(logLines []string) map[string]*Folder {
	currentFolderString := rootDirectory
	stk := common.Stack[string]{}

	currentFolder := &Folder{name: currentFolderString, folders: []*Folder{}, files: []*File{}}

	lookup := make(map[string]*Folder)
	lookup[currentFolderString] = currentFolder

	for i := 0; i < len(logLines); i++ {
		logLine := logLines[i]

		split := strings.Split(logLine, " ")

		if split[0] != "$" {
			fmt.Println(logLines[i])
			panic("v any")
		}

		command := split[1]

		if command == "cd" {
			targetDirectory := split[2]

			switch targetDirectory {
			case rootDirectory:
				{
					currentFolderString = rootDirectory
					stk = common.Stack[string]{}
					stk.Push(currentFolderString)
				}
			case "..":
				{
					for !stk.Empty() && stk.Top() == currentFolderString {
						_ = stk.Pop()
					}
					if stk.Empty() {
						currentFolderString = rootDirectory
					} else {
						currentFolderString = stk.Pop()
					}
				}
			default:
				{
					stk.Push(currentFolderString)

					if currentFolderString != rootDirectory {
						currentFolderString = currentFolderString + "/" + targetDirectory
					} else {
						currentFolderString = rootDirectory + split[2]
					}

					if _, ok := lookup[currentFolderString]; !ok {
						lookup[currentFolderString] = &Folder{name: currentFolderString, folders: []*Folder{}, files: []*File{}}
					}
				}
			}

			currentFolder = lookup[currentFolderString]

		}

		if command == "ls" {
			j := i + 1

			for ; j < len(logLines); j++ {
				lsLogLine := logLines[j]

				split = strings.Split(lsLogLine, " ")

				if split[0] == "$" {
					j--
					break
				}

				name := split[1]

				if split[0] == "dir" {
					newFolderString := ""

					if currentFolderString != rootDirectory {
						newFolderString = currentFolderString + "/" + name
					} else {
						newFolderString = rootDirectory + name
					}

					if _, ok := lookup[newFolderString]; !ok {
						lookup[newFolderString] = &Folder{name: newFolderString, folders: []*Folder{}, files: []*File{}}
					}
					currentFolder.folders = append(currentFolder.folders, lookup[newFolderString])

				} else {

					size, _ := strconv.Atoi(split[0])
					currentFolder.files = append(currentFolder.files, &File{name: name, size: size})
				}

			}

			i = j
		}

	}
	return lookup
}
