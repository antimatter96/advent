package main

import (
	"fmt"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

type robot struct {
	x, y   int
	vX, vY int
}

var M, N int

func parsePart1(inp []string) []*robot {
	robots := make([]*robot, 0, len(inp)-1)
	for i, inp_s := range inp {
		if i == 0 {
			fmt.Sscanf(inp_s, "%d %d", &M, &N)
			continue
		}
		robot := &robot{}
		fmt.Sscanf(inp_s, "p=%d,%d v=%d,%d", &robot.x, &robot.y, &robot.vX, &robot.vY)
		robots = append(robots, robot)
	}

	return robots
}

func parsePart2(inp []string) []*robot {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(robots []*robot) int {

	const moves int = 100

	xHalf := (M - 1) / 2
	yHalf := (N - 1) / 2

	n1, n2, n3, n4 := 0, 0, 0, 0

	for _, robot := range robots {

		robot.x += (100 * robot.vX) % M
		robot.y += (100 * robot.vY) % N

		if robot.x >= M {
			robot.x = robot.x % M
		} else if robot.x < 0 {
			robot.x += M
		}
		if robot.y >= N {
			robot.y = robot.y % N
		} else if robot.y < 0 {
			robot.y += N
		}

		common.Log.Debug().Int("x", robot.x).Int("y", robot.y).Send()

		if robot.x == xHalf || robot.y == yHalf {
			continue
		}

		if robot.x < xHalf {
			if robot.y < yHalf {
				n1++
			} else {
				n2++
			}

		} else {
			if robot.y < yHalf {
				n3++
			} else {
				n4++
			}
		}

	}

	common.Log.Debug().Int("n1", n1).Int("n2", n2).Int("n3", n3).Int("n4", n4).Send()
	return n1 * n2 * n3 * n4
}

func Part2(robots []*robot) int {
	var goodPoint int

	graph := make(common.Graph[string], M)
	for i := 0; i < M; i++ {
		graph[i] = make([]string, N)
		for j := 0; j < N; j++ {
			graph[i][j] = "."
		}
	}

timeLoop:
	for i := 0; i < 10000; i++ {
		for _, robot := range robots {
			graph[robot.x][robot.y] = "."
		}

		for _, robot := range robots {
			robot.x += robot.vX
			robot.y += robot.vY

			if robot.x >= M {
				robot.x = robot.x % M
			} else if robot.x < 0 {
				robot.x += M
			}
			if robot.y >= N {
				robot.y = robot.y % N
			} else if robot.y < 0 {
				robot.y += N
			}
		}

		for _, robot := range robots {
			graph[robot.x][robot.y] = "*"
		}

		for _, row := range graph {
			if strings.Contains(strings.Join(row, ""), "***********************") {
				goodPoint = i + 1
				break timeLoop
			}
		}
		graph.Print()
	}

	return goodPoint
}
