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

func ExtractNumbersFrom[K comparable](graph [][]K, x, y int, direction int, whenToProcees map[K]struct{}, whenToStop map[K]struct{}) []K {
	Log.Debug().Ints("extracting numbers", []int{x, y}).Send()

	numbers := make([]K, 0)

	if x > len(graph)-1 || x < 0 {
		return numbers
	}
	if y > len(graph[0])-1 || y < 0 {
		return numbers
	}

	if _, ok := whenToStop[graph[x][y]]; ok {
		return numbers
	}

	if _, ok := whenToProcees[graph[x][y]]; !ok {
		return numbers
	}

	Log.Debug().Str("graph[x,y]", fmt.Sprintf("%c", graph[x][y])).Send()

	numbers = append(numbers, graph[x][y])

	if direction == 1 || direction == 0 {
		for j := y + 1; j < len(graph[0]); j++ {

			if _, ok := whenToStop[graph[x][j]]; ok {
				break
			}
			if _, ok := whenToProcees[graph[x][j]]; !ok {
				break
			}

			numbers = append(numbers, graph[x][j])
			Log.Debug().Str("graph[x,y]", fmt.Sprintf("%c", graph[x][j])).Send()
		}
	}

	if direction == -1 || direction == 0 {
		for j := y - 1; j > -1; j-- {

			if _, ok := whenToStop[graph[x][j]]; ok {
				break
			}
			if _, ok := whenToProcees[graph[x][j]]; !ok {
				break
			}

			numbers = append([]K{graph[x][j]}, numbers...)
			Log.Debug().Str("graph[x,y]", fmt.Sprintf("%c", graph[x][j])).Send()
		}
	}
	return numbers
}

func StartExtractNumbersFrom[K comparable](graph [][]K, x, y int, whenToProcees map[K]struct{}, whenToStop map[K]struct{}) [][]K {

	Log.Debug().Ints("StartExtractNumbersFrom", []int{x, y}).Send()

	r := ExtractNumbersFrom[K](graph, x, y+1, 1, whenToProcees, whenToStop)
	l := ExtractNumbersFrom[K](graph, x, y-1, -1, whenToProcees, whenToStop)

	var u [][]K
	if x-1 > -1 {
		if _, ok := whenToProcees[graph[x-1][y]]; ok {
			u = append(u, ExtractNumbersFrom[K](graph, x-1, y, 0, whenToProcees, whenToStop))
		} else {
			ur := ExtractNumbersFrom[K](graph, x-1, y+1, 1, whenToProcees, whenToStop)
			ul := ExtractNumbersFrom[K](graph, x-1, y-1, -1, whenToProcees, whenToStop)

			u = append(u, ul)
			u = append(u, ur)
		}

	}

	var d [][]K
	if x+1 < len(graph) {
		if _, ok := whenToProcees[graph[x+1][y]]; ok {
			d = append(d, ExtractNumbersFrom[K](graph, x+1, y, 0, whenToProcees, whenToStop))
		} else {
			dr := ExtractNumbersFrom[K](graph, x+1, y+1, 1, whenToProcees, whenToStop)
			dl := ExtractNumbersFrom[K](graph, x+1, y-1, -1, whenToProcees, whenToStop)

			d = append(d, dl)
			d = append(d, dr)
		}

	}

	var final [][]K

	Log.Debug().Str("StartExtractNumbersFrom left", toString[K](l)).Send()
	Log.Debug().Str("StartExtractNumbersFrom right", toString[K](r)).Send()

	for _, uu := range u {
		Log.Debug().Str("StartExtractNumbersFrom up", toString[K](uu)).Send()
	}
	for _, dd := range d {
		Log.Debug().Str("StartExtractNumbersFrom down", toString[K](dd)).Send()
	}

	final = append(final, l)
	final = append(final, r)

	final = append(final, u...)
	final = append(final, d...)

	var nonEmptyfinal [][]K

	for _, num := range final {
		if len(num) > 0 {
			nonEmptyfinal = append(nonEmptyfinal, num)
		}
	}

	return nonEmptyfinal
}

func toString[K comparable](arr []K) string {
	var s string
	for _, e := range arr {
		s += fmt.Sprintf("%c", e)
	}

	return s
}
