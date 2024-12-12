package common

func FloodFill[T comparable](graph Graph[T], goDiagnol bool, replaceWith T) map[int][]*Point {
	visited := make(map[string]struct{})
	fills := make(map[int][]*Point)

	N := -1

	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[0]); j++ {
			points := floodFillRec(graph, i, j, goDiagnol, replaceWith, visited)
			if len(points) > 0 {
				fills[N] = points
				N++
			}
		}
	}

	return fills
}

func floodFillRec[T comparable](graph Graph[T], i, j int, goDiagnol bool, replaceWith T, visited map[string]struct{}) []*Point {
	start := &Point{i, j}

	if _, ok := visited[start.String()]; ok {
		return nil
	}

	visited[start.String()] = struct{}{}

	points := []*Point{{X: i, Y: j}}

	for incX := -1; incX < 2; incX++ {
		x := i + incX
		for incY := -1; incY < 2; incY++ {
			y := j + incY
			if x < 0 || x >= len(graph) || y < 0 || y >= len(graph[0]) {
				continue
			}
			if !goDiagnol && incX*incY != 0 {
				continue
			}

			if x == i && y == j {
				continue
			}

			if graph[x][y] == graph[i][j] {
				points = append(points, floodFillRec(graph, x, y, goDiagnol, replaceWith, visited)...)
			}
		}
	}

	if ADVENT_DEBUG {
		graph[i][j] = replaceWith
	}
	return points
}
