package common

import (
	"fmt"
	"strings"
)

type Graph[T comparable] [][]T

func (graph *Graph[T]) InBounds(x, y int) bool {
	return x > -1 && y > -1 && x < len(*graph) && y < len((*graph)[0])
}

func (graph Graph[T]) Print() {
	Log.Debug().Str("", "===========").Send()

	_, isRune := any(graph[0][0]).(rune)

	for i := 0; i < len(graph); i++ {
		strB := strings.Builder{}
		for j := 0; j < len(graph[i]); j++ {
			if isRune {
				strB.WriteString(fmt.Sprintf("%c", graph[i][j]))
				continue
			}
			strB.WriteString(fmt.Sprintf("%s", graph[i][j]))
		}
		Log.Debug().Str("", strB.String()).Send()
	}

	Log.Debug().Str("", "===========").Send()
}

func CopyPointerGraph[T comparable](graph Graph[*T]) Graph[*T] {
	mp := make([][]*T, len(graph))
	for i := 0; i < len(mp); i++ {
		mp[i] = make([]*T, len(graph[0]))
		copy(mp[i], graph[i])
	}
	return mp

}

func CopyGraph[T comparable](graph Graph[T]) Graph[T] {
	mp := make([][]T, len(graph))
	for i := 0; i < len(mp); i++ {
		mp[i] = make([]T, len(graph[0]))
		copy(mp[i], graph[i])
	}
	return mp

}
