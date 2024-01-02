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

func Part1(field common.Graph[string]) int {
	field.Print()

	startX, startY := field.Find("S")

	bfsContainer := bfsContainer{}

	level := 64

	bfsContainer.Start(startX, startY, field, level)

	field.Print()

	ans := bfsContainer.AtLevel(level)

	common.Log.Debug().Int("Level", level).Int("ans", ans).Send()

	return ans
}

func Part2(field common.Graph[string]) int {
	field.Print()

	startX, startY := field.Find("S")

	common.Log.Debug().Ints("", []int{startX, startY}).Send()

	level := 50

	bfsContainer := bfsContainer{}
	bfsContainer.StartInfinite(startX, startY, field, level)

	ans := bfsContainer.AtLevel(level)

	common.Log.Debug().Int("Level", level).Int("ans", ans).Send()

	return ans
}

type point struct {
	x, y int
}

func (p *point) fromString(s string) {
	fmt.Sscanf(s, "%d,%d", &p.x, &p.y)
}

func (p *point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p *point) Neighbours4() []point {
	return []point{
		{p.x - 1, p.y},
		{p.x + 1, p.y},
		{p.x, p.y + 1},
		{p.x, p.y - 1},
	}
}

func (p *point) outOfBounds(maxX, maxY int) bool {
	if p.x < 0 || p.x >= maxX {
		return true
	}

	if p.y < 0 || p.y >= maxY {
		return true
	}

	return false
}

type bfsContainer struct {
	distance [][]int

	atLevel map[int]int
}

func (b *bfsContainer) Start(x, y int, field common.Graph[string], maxLevel int) {
	b.atLevel = make(map[int]int)
	b.basicBFS(point{x, y}, field, maxLevel)
}

func (b *bfsContainer) StartInfinite(x, y int, field common.Graph[string], maxLevel int) {
	b.atLevel = make(map[int]int)
	b.infinteBFS(point{x, y}, field, maxLevel)
}

func (b *bfsContainer) AtLevel(level int) int {
	return b.atLevel[level]
}

func (b *bfsContainer) basicBFS(startPoint point, field common.Graph[string], maxLevel int) {
	m, n := len(field), len(field[0])

	set := common.QueueSet[string]{}
	set.Push(startPoint.String())

	level := 0

	for level < maxLevel {
		level++

		l := set.Size()

		// field.Replace("O", "x")

		for nPopped := 0; nPopped < l; nPopped++ {
			pStr := set.Pop()

			p := point{}
			p.fromString(pStr)

			// field[p.x][p.y] = "O"

			for _, neighbour := range p.Neighbours4() {
				if neighbour.outOfBounds(m, n) {
					continue
				}
				if field[neighbour.x][neighbour.y] == "#" {
					continue
				}

				set.Push(neighbour.String())
			}
		}
		common.Log.Debug().Int("popped", l).Int("Level", level-1).Int("Next", set.Size()).Send()

		b.atLevel[level] = set.Size()
	}

}

func (b *bfsContainer) infinteBFS(startPoint point, field common.Graph[string], maxLevel int) {
	set := common.QueueSet[string]{}
	set.Push(startPoint.String())

	level := 0

	for level < maxLevel+1 {
		level++

		l := set.Size()

		for nPopped := 0; nPopped < l; nPopped++ {
			pStr := set.Pop()

			p := point{}
			p.fromString(pStr)

			for _, neighbour := range p.Neighbours4() {
				if field.AtInfinite(neighbour.x, neighbour.y) == "#" {
					continue
				}

				set.Push(neighbour.String())
			}
		}
		// d1 := set.Size() - b.atLevel[level-1]
		// d2 := b.atLevel[level-1] - b.atLevel[level-2]
		// d3 := b.atLevel[level-2] - b.atLevel[level-3]

		// d4 := d1 - d2
		// d5 := d2 - d3

		// d6 := d4 - d5

		// common.Log.Debug().Int("d1", d1).Int("d2", d2).Int("d3", d3).Int("d4", d4).Int("d5", d5).Int("d6", d6).Send()

		if (level-1)%len(field) == 0 {
			common.Log.Debug().Int(" b=>  ", l).Int(" a ", level-1).Send()

		}

		b.atLevel[level] = set.Size()
	}

}
