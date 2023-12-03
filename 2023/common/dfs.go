package common

import "fmt"

func FloodFill[K comparable](graph [][]K, x, y int, whatToReplaceWith K, whenToProcees map[K]struct{}, whenToStop map[K]struct{}) {

	Log.Debug().Ints("flood filling", []int{x, y}).Send()

	if x > len(graph)-1 || x < 0 {
		return
	}
	if y > len(graph[0])-1 || y < 0 {
		return
	}

	Log.Debug().Str("graph[x,y]", fmt.Sprintf("%c", graph[x][y])).Send()

	if _, ok := whenToStop[graph[x][y]]; ok {
		return
	}

	if _, ok := whenToProcees[graph[x][y]]; !ok {
		return
	}

	graph[x][y] = whatToReplaceWith

	FloodFill[K](graph, x, y+1, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x, y-1, whatToReplaceWith, whenToProcees, whenToStop)

	FloodFill[K](graph, x+1, y+1, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x+1, y, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x+1, y-1, whatToReplaceWith, whenToProcees, whenToStop)

	FloodFill[K](graph, x-1, y+1, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x-1, y, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x-1, y-1, whatToReplaceWith, whenToProcees, whenToStop)
}

func StartFloodFill[K comparable](graph [][]K, x, y int, whatToReplaceWith K, whenToProcees map[K]struct{}, whenToStop map[K]struct{}) {
	FloodFill[K](graph, x, y+1, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x, y-1, whatToReplaceWith, whenToProcees, whenToStop)

	FloodFill[K](graph, x+1, y+1, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x+1, y, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x+1, y-1, whatToReplaceWith, whenToProcees, whenToStop)

	FloodFill[K](graph, x-1, y+1, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x-1, y, whatToReplaceWith, whenToProcees, whenToStop)
	FloodFill[K](graph, x-1, y-1, whatToReplaceWith, whenToProcees, whenToStop)
}
