package main

import (
	"fmt"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

func main() {
	rawInput := common.TakeInputAsString()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp string) common.Graph[string] {
	inp = strings.TrimSpace(inp)

	inpLines := strings.Split(inp, "\n")

	field := make(common.Graph[string], len(inpLines))

	for i := 0; i < len(inpLines); i++ {
		field[i] = strings.Split(strings.TrimSpace(inpLines[i]), "")
	}

	return field
}

func parsePart2(inp string) common.Graph[string] {
	return parsePart1(inp)
}

func Run(inp string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

var options = []string{"|", "-", "J", "7", "F", "L"}

func Part1(field common.Graph[string]) int {
	max := 0

	startX, startY := field.Find("S")
	common.Log.Debug().Int("x", startX).Int("y", startY).Send()

	for _, option := range options {
		bfsContainer := bfsContainer{}

		if canFit(startX, startY, option, field) {
			common.Log.Debug().Str("option", option).Send()

			field[startX][startY] = option

			bfsContainer.Start(startX, startY, &field, same)

			max = bfsContainer.MaxDistance()

			common.Log.Debug().Int("max", max).Str("for option", option).Send()
			break
		}

	}

	return max
}

func Part2(field common.Graph[string]) int {
	startX, startY := field.Find("S")

	field.Print()

	bfsContainer := bfsContainer{}

	finalOption := ""
	for _, possibleOption := range options {
		if canFit(startX, startY, possibleOption, field) {
			finalOption = possibleOption
			break
		}
	}
	common.Log.Debug().Str("finalOption", finalOption).Send()

	field[startX][startY] = finalOption
	bfsContainer.Start(startX, startY, &field, replacer)

	field.Print()

	bfsContainer.cleanUp(&field)
	field.Print()
	ans := bfsContainer.countInside(&field)
	field.Print()

	return ans

}

func outOfBounds(x, y, maxX, maxY int) bool {
	if x < 0 || x >= maxX {
		return true
	}

	if y < 0 || y >= maxY {
		return true
	}

	return false
}

type fitOption struct {
	newX1, newY1  func(int) int
	notAllowedIn1 string

	newX2, newY2 func(int) int

	notAllowedIn2 string
}

func nop(i int) int { return i }
func inc(i int) int { return i + 1 }
func dec(i int) int { return i - 1 }

var fitOptions = map[string]*fitOption{
	"|": {
		// .1.
		// .|.
		// .2.

		newX1: dec, newY1: nop,
		notAllowedIn1: "-LJ", // |7F

		newX2: inc, newY2: nop,
		notAllowedIn2: "-7F", // |LJ
	},
	"-": {
		// ...
		// 1-2
		// ...

		newX1: nop, newY1: dec,
		notAllowedIn1: "|J7", // -LF

		newX2: nop, newY2: inc,
		notAllowedIn2: "|LF", // -J7
	},
	"L": {
		// .1.
		// .L2
		// ...

		newX1: dec, newY1: nop,
		notAllowedIn1: "-LJ", // |7F

		newX2: nop, newY2: inc,
		notAllowedIn2: "|LF", // -J7
	},
	"J": {
		// .1.
		// 2J.
		// ...

		newX1: dec, newY1: nop,
		notAllowedIn1: "-LJ", // |7F

		newX2: nop, newY2: dec,
		notAllowedIn2: "|J7", // -LF
	},
	"7": {
		// ...
		// 17.
		// .2.

		newX1: nop, newY1: dec,
		notAllowedIn1: "|7J", // -LF

		newX2: inc, newY2: nop,
		notAllowedIn2: "-7F", // |LJ
	},
	"F": {
		// ...
		// .F1
		// .2.

		newX1: nop, newY1: inc,
		notAllowedIn1: "|LF", // -J7

		newX2: inc, newY2: nop,
		notAllowedIn2: "-7F", // |LJ
	},
}

func canFit(x, y int, option string, field common.Graph[string]) bool {
	maxX := len(field)
	maxY := len(field[0])

	fitOption := fitOptions[option]

	x1, y1 := fitOption.newX1(x), fitOption.newY1(y)
	x2, y2 := fitOption.newX2(x), fitOption.newY2(y)

	if outOfBounds(x1, y1, maxX, maxY) || outOfBounds(x2, y2, maxX, maxY) {
		return false
	}

	if field[x1][y1] == "." || field[x2][y2] == "." {
		return false
	}

	if strings.ContainsAny(field[x1][y1], fitOption.notAllowedIn1) {
		return false
	}

	if strings.ContainsAny(field[x2][y2], fitOption.notAllowedIn2) {
		return false
	}

	return true
}

type point struct {
	x, y     int
	distance int
}

func (p *point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func pointToString(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

type bfsContainer struct {
	queue    common.Queue[point]
	done     map[string]bool
	distance [][]int
}

func (b *bfsContainer) Start(x, y int, field *common.Graph[string], replacerFunc func(string) string) {
	b.queue = common.Queue[point]{}
	b.queue.Push(point{x, y, 0})

	b.distance = make([][]int, len(*field))
	for i := 0; i < len(b.distance); i++ {
		b.distance[i] = make([]int, len((*field)[i]))
	}

	b.done = make(map[string]bool)

	b.basicBFS(field, replacerFunc)

	common.Log.Debug().Int("don", len(b.done)).Send()
}

func (b *bfsContainer) MaxDistance() int {
	return common.MaxIntInts(b.distance)
}

func (b *bfsContainer) basicBFS(field *common.Graph[string], replacerFunc func(string) string) {
	for !b.queue.Empty() {

		p := b.queue.Pop()

		if b.done[p.String()] {
			continue
		}
		b.done[p.String()] = true

		b.distance[p.x][p.y] = p.distance

		tile := (*field)[p.x][p.y]

		option := fitOptions[tile]

		x1, y1 := option.newX1(p.x), option.newY1(p.y)
		x2, y2 := option.newX2(p.x), option.newY2(p.y)

		b.queue.Push(point{x1, y1, p.distance + 1})
		b.queue.Push(point{x2, y2, p.distance + 1})

		(*field)[p.x][p.y] = replacerFunc(tile)
	}

}

func (b *bfsContainer) cleanUp(field *common.Graph[string]) {

	for x := 0; x < len(*field); x++ {

		for y := 0; y < len((*field)[0]); y++ {
			if b.done[pointToString(x, y)] {
				break
			}
			(*field)[x][y] = " "
		}

		for y := len((*field)[0]) - 1; y > -1; y-- {
			if b.done[pointToString(x, y)] {
				break
			}
			(*field)[x][y] = " "
		}
	}

}

func (b *bfsContainer) countInside(field *common.Graph[string]) int {
	count := 0

	for x := 0; x < len(*field); x++ {
		common.Log.Debug().Int("line", x).Send()

		lastWas := "x"
		inside := false

		for y := 0; y < len((*field)[0]); y++ {

			tile := (*field)[x][y]

			if !b.done[pointToString(x, y)] {
				(*field)[x][y] = "O"
				if inside {
					count++
					(*field)[x][y] = "I"
				}
				continue
			}

			if tile == "-" {
				continue
			}

			if tile == "|" {
				inside = !inside
				continue
			}

			verticalTile := (tile == "J" || tile == "L" || tile == "7" || tile == "F")

			if !verticalTile {
				continue
			}

			if lastWas == "x" {
				lastWas = tile
				continue
			}

			if lastWas == "L" && tile == "7" {
				inside = !inside
			}

			if lastWas == "F" && tile == "J" {
				inside = !inside
			}
			lastWas = "x"

		}
	}
	common.Log.Debug().Int("Final", count).Send()

	return count
}

func same(in string) string {
	return in
}

// var replaceMap = map[string]string{
// 	"J": "┘",
// 	"L": "└",
// 	"7": "┐",
// 	"F": "┌",
// 	"|": "│",
// 	"-": "─",
// }

func replacer(in string) string {
	return in
}
