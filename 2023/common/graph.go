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
			strB.WriteString(fmt.Sprintf("%c", graph[i][j]))
		}
		Log.Debug().Str("", strB.String()).Send()
	}

	Log.Debug().Str("", "===========").Send()
}
