package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rawInput := takeInput()

	fmt.Println(Run(rawInput))
}

func takeInput() []string {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	return inp
}

func parsePart1(inp []string) [][]int {
	grid := make([][]int, 0)

	for _, s := range inp {
		nums := strings.Split(s, "")

		var temp []int

		for _, num := range nums {
			x, _ := strconv.Atoi(num)
			temp = append(temp, x)
		}

		grid = append(grid, temp)
	}

	return grid
}

func parsePart2(inp []string) [][]int {
	grid := make([][]int, 0)

	for _, s := range inp {
		nums := strings.Split(s, "")

		var temp []int

		for _, num := range nums {
			x, _ := strconv.Atoi(num)
			temp = append(temp, x)
		}

		n := len(temp)
		for i := 1; i < 5; i++ {
			extra := make([]int, 0)

			for j := 0; j < n; j++ {
				if temp[j]+i > 9 {
					extra = append(extra, temp[j]+i-9)
				} else {
					extra = append(extra, temp[j]+i)
				}
			}
			temp = append(temp, extra...)
		}

		grid = append(grid, temp)
	}

	l := len(grid)

	original := make([][]int, 0)
	for i := 0; i < len(grid); i++ {
		tt := make([]int, 0)
		for j := 0; j < len(grid[i]); j++ {
			tt = append(tt, grid[i][j])
		}
		original = append(original, tt)
	}

	for i := 1; i < 5; i++ {

		pp := make([][]int, 0)
		for i := 0; i < len(original); i++ {
			tt := make([]int, 0)
			for j := 0; j < len(original[i]); j++ {
				tt = append(tt, original[i][j])
			}
			pp = append(pp, tt)
		}

		grid = append(grid, pp...)
	}

	for i := 1; i < 5; i++ {
		// printGrid(grid)
		for k := (i * l); k < (i+1)*l; k++ {
			for j := 0; j < len(grid[0]); j++ {
				grid[k][j] += i
				if grid[k][j] > 9 {
					grid[k][j] -= 9
				}
			}
		}

		// printGrid(grid)
	}

	return grid
}

func Run(inp []string) (int, int) {
	p := parsePart1(inp)
	//p2 := parsePart2(inp)

	return Part1(p), Part2(p)
}

func Vertex(i, j int) string {
	return fmt.Sprintf("%03d %03d", i, j)
}

func Part1(grid [][]int) int {
	start := time.Now()
	G := make(map[string]map[string]int)

	for i := range grid {
		for j := range grid {
			u := Vertex(i, j)
			G[u] = make(map[string]int)

			if i > 0 {
				G[u][Vertex(i-1, j)] = grid[i-1][j]
			}
			if i+1 < len(grid) {
				G[u][Vertex(i+1, j)] = grid[i+1][j]
			}

			if j > 0 {
				G[u][Vertex(i, j-1)] = grid[i][j-1]
			}
			if j+1 < len(grid[0]) {
				G[u][Vertex(i, j+1)] = grid[i][j+1]
			}
		}
	}

	visited := make(map[string]bool)
	distance := make(map[string]int)

	Q := Queue{arr: []string{}, dist: &distance}
	for i := range grid {
		for j := range grid {
			distance[Vertex(i, j)] = math.MaxInt
			Q.Add(Vertex(i, j))
		}
	}
	distance[Vertex(0, 0)] = 0
	visited[Vertex(0, 0)] = true

	fmt.Println(Q.Len())
	for !Q.Empty() {
		v := Q.Least()
		visited[v] = true

		for u := range G[v] {
			if !visited[u] {
				alt := distance[v] + G[v][u]

				if alt < distance[u] {
					distance[u] = alt
				}
			}
		}

	}

	t := time.Now()
	fmt.Println(t.Sub(start), Q.Len())

	// fmt.Println(distance)
	return distance[Vertex(len(grid)-1, len(grid[0])-1)]
}

func Part2(grid [][]int) int {
	start := time.Now()
	G := make(map[string]map[string]int)

	for i := range grid {
		for j := range grid {
			u := Vertex(i, j)
			G[u] = make(map[string]int)

			if i > 0 {
				G[u][Vertex(i-1, j)] = grid[i-1][j]
			}
			if i+1 < len(grid) {
				G[u][Vertex(i+1, j)] = grid[i+1][j]
			}

			if j > 0 {
				G[u][Vertex(i, j-1)] = grid[i][j-1]
			}
			if j+1 < len(grid[0]) {
				G[u][Vertex(i, j+1)] = grid[i][j+1]
			}
		}
	}

	visited := make(map[string]bool)
	distance := make(map[string]*Item)

	pq := make(PriorityQueue, 0)
	for i := range grid {
		for j := range grid {
			distance[Vertex(i, j)] = &Item{value: Vertex(i, j), index: (len(grid) * i) + j, priority: math.MaxInt}
			heap.Push(&pq, distance[Vertex(i, j)])
		}
	}
	distance[Vertex(0, 0)].priority = 0
	pq.update(distance[Vertex(0, 0)], Vertex(0, 0), 0)
	visited[Vertex(0, 0)] = true

	fmt.Println(pq.Len())

	for pq.Len() > 0 {
		vItem := heap.Pop(&pq).(*Item)
		visited[vItem.value] = true

		for u := range G[vItem.value] {
			if !visited[u] {
				alt := distance[vItem.value].priority + G[vItem.value][u]

				if alt < distance[u].priority {
					distance[u].priority = alt

					pq.update(distance[u], u, alt)
				}
			}
		}

	}

	t := time.Now()
	fmt.Println(t.Sub(start), pq.Len())
	return distance[Vertex(len(grid)-1, len(grid[0])-1)].priority
}

func printGrid(grid [][]int) {
	s := strings.Builder{}
	for _, row := range grid {
		for _, num := range row {
			s.WriteString(fmt.Sprintf("% 2d", num))
		}
		s.WriteString("\n")
	}
	fmt.Println(s.String())
}
