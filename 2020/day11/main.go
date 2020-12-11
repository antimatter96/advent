package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Errorf(scanner.Err().Error())
	}

	day2(inp)
}

func day1(arr []string) {
	byteArr := make([][]byte, len(arr))
	for i := 0; i < len(arr); i++ {
		byteArr[i] = []byte(arr[i])
	}
	// fmt.Println(arr)

	// doChanges(byteArr)

	//j := 0
	for j := 0; j < 100000; j++ {
		xx := doChanges(byteArr, 4, getOccupiedAdjacent)
		//fmt.Println("===")
		//fmt.Println("changed", xx)
		// for i := 0; i < len(arr); i++ {
		// 	fmt.Println(string(byteArr[i]))
		// }
		//fmt.Println(count(byteArr))
		if !xx {
			fmt.Println(count(byteArr))
			break
		}
	}

	// fmt.Println("===")
}

func count(arr [][]byte) int {
	total := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == '#' {
				total++
			}
		}
	}
	return total
}

func doChanges(arr [][]byte, limit int, f func(int, int, [][]byte) int) bool {
	// fmt.Println("doChanges")
	fill := make(map[string]bool)
	empty := make(map[string]bool)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] != '.' {
				occupied := f(i, j, arr)

				// fmt.Println(i, j, occupied)

				if occupied == 0 && arr[i][j] == 'L' {
					fill[stringify(i, j)] = true
				} else if occupied >= limit && arr[i][j] == '#' {
					empty[stringify(i, j)] = true

					// fmt.Println("empty", i, j, occupied)
				}
			}
		}
	}

	// fmt.Println(fill, empty)

	var x, y int
	for s := range fill {
		//mt.Println("fill", s)
		fmt.Sscanf(s, "%d %d", &x, &y)
		//fmt.Println("fill", x, y)
		arr[x][y] = '#'
		// for i := 0; i < len(arr); i++ {
		// 	fmt.Println(string(arr[i]))
		// }
	}
	for s := range empty {
		//fmt.Println("empty", s)
		fmt.Sscanf(s, "%d %d", &x, &y)
		arr[x][y] = 'L'
	}

	return len(fill)+len(empty) > 0
}

func stringify(x, y int) string {
	return fmt.Sprintf("%d %d", x, y)
}

func getOccupiedAdjacent(i, j int, arr [][]byte) int {
	maxX := len(arr)
	maxY := len(arr[0])
	occupied := 0

	occupied += findStuff(i, j, arr, dec, dec, maxX, maxY)
	occupied += findStuff(i, j, arr, dec, inc, maxX, maxY)
	occupied += findStuff(i, j, arr, inc, inc, maxX, maxY)
	occupied += findStuff(i, j, arr, inc, dec, maxX, maxY)

	occupied += findStuff(i, j, arr, same, dec, maxX, maxY)
	occupied += findStuff(i, j, arr, same, inc, maxX, maxY)
	occupied += findStuff(i, j, arr, dec, same, maxX, maxY)
	occupied += findStuff(i, j, arr, inc, same, maxX, maxY)

	return occupied
}

func day2(arr []string) {
	byteArr := make([][]byte, len(arr))
	for i := 0; i < len(arr); i++ {
		byteArr[i] = []byte(arr[i])
	}

	//j := 0
	for j := 0; j < 100000; j++ {
		xx := doChanges(byteArr, 5, getOccupiedAdjacent2)
		//fmt.Println("===")
		//fmt.Println("changed", xx)
		// for i := 0; i < len(arr); i++ {
		// 	fmt.Println(string(byteArr[i]))
		// }
		//fmt.Println(count(byteArr))
		if !xx {
			fmt.Println(count(byteArr))
			break
		}
	}

	// fmt.Println("===")
}

func getOccupiedAdjacent2(i, j int, arr [][]byte) int {
	maxX := len(arr)
	maxY := len(arr[0])
	occupied := 0

	occupied += findStuffLong(i, j, arr, dec, dec, maxX, maxY)
	occupied += findStuffLong(i, j, arr, dec, inc, maxX, maxY)
	occupied += findStuffLong(i, j, arr, inc, inc, maxX, maxY)
	occupied += findStuffLong(i, j, arr, inc, dec, maxX, maxY)

	occupied += findStuffLong(i, j, arr, same, dec, maxX, maxY)
	occupied += findStuffLong(i, j, arr, same, inc, maxX, maxY)
	occupied += findStuffLong(i, j, arr, dec, same, maxX, maxY)
	occupied += findStuffLong(i, j, arr, inc, same, maxX, maxY)

	return occupied
}

func inc(i int) int {
	return i + 1
}
func dec(i int) int {
	return i - 1
}
func same(i int) int {
	return i
}

func findStuffLong(i, j int, arr [][]byte, updateI, updateJ func(int) int, limitX, limitY int) int {

	for true {
		i = updateI(i)
		j = updateJ(j)

		if i < 0 || i >= limitX || j < 0 || j >= limitY {
			return 0
		}

		if arr[i][j] == 'L' {
			return 0
		}

		if arr[i][j] == '#' {
			return 1
		}

	}

	return 0
}

func findStuff(i, j int, arr [][]byte, updateI, updateJ func(int) int, limitX, limitY int) int {
	i = updateI(i)
	j = updateJ(j)

	if i < 0 || i >= limitX || j < 0 || j >= limitY {
		return 0
	}

	if arr[i][j] == '#' {
		return 1
	}

	return 0
}
