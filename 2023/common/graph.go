package common

import (
	"fmt"
	"strings"
)

type Graph[T comparable] [][]T

func (graph Graph[T]) Print() {
	Log.Debug().Str("", "===========").Send()

	for i := 0; i < len(graph); i++ {
		strB := strings.Builder{}
		for j := 0; j < len(graph[i]); j++ {
			strB.WriteString(fmt.Sprintf("%v", graph[i][j]))
		}
		Log.Debug().Str("", strB.String()).Send()
	}

	Log.Debug().Str("", "===========").Send()
}

func (graph Graph[T]) At(x, y int) T {
	return graph[x][y]
}

func (graph Graph[T]) AtInfinite(x, y int) T {
	// Log.Debug().Int("orignalX", x).Int("orignalY", y).Send()

	for x < 0 {
		x += len(graph)
	}
	for y < 0 {
		y += len(graph[0])
	}

	x = x % len(graph)
	y = y % len(graph[0])

	// Log.Debug().Int("newX", x).Int("newT", y).Send()
	return graph[x][y]
}

func PrintGraphRune(graph Graph[rune]) {
	Log.Debug().Str("", "===========").Send()

	for i := 0; i < len(graph); i++ {
		strB := strings.Builder{}
		for j := 0; j < len(graph[i]); j++ {
			strB.WriteRune(graph[i][j])
		}
		Log.Debug().Str("", strB.String()).Send()
	}

	Log.Debug().Str("", "===========").Send()
}

func (graph Graph[T]) Find(ele T) (int, int) {

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == ele {
				return i, j
			}
		}
	}

	return -1, -1
}

func (graph Graph[T]) Count(ele T) int {
	count := 0

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == ele {
				count++
			}
		}
	}

	return count
}

func (graph Graph[T]) Replace(what T, with T) {

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == what {
				graph[i][j] = with
			}
		}
	}

}
