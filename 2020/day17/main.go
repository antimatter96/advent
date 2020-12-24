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
		panic(scanner.Err().Error())
	}

	day22(inp)
}

func count(mp map[string]bool) int {
	total := 0
	for _, v := range mp {
		if v {
			total++
		}
	}
	return total
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

func intify3D(s string) (x, y, z int) {
	fmt.Sscanf(s, "%d,%d,%d", &x, &y, &z)
	return
}
func stringify3D(x, y, z int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, z)
}

func getOccupied(i, j, k int, mp *map[string]bool) int {
	occupied := 0

	for x, fI := range funcs {
		for y, fJ := range funcs {
			for z, fK := range funcs {
				if x+y+z != 0 { // same,same,same,same
					occupied += findStuff(i, j, k, mp, fI, fJ, fK)
				}
			}
		}
	}

	return occupied
}

func findStuff(i, j, k int, mp *map[string]bool, updateI, updateJ, updateK func(int) int) int {
	i = updateI(i)
	j = updateJ(j)
	k = updateK(k)

	v, ok := (*mp)[stringify3D(i, j, k)]
	if ok && v {
		return 1
	} else if !ok {
		(*mp)[stringify3D(i, j, k)] = false
	}

	return 0
}

func doChanges3D(mp *map[string]bool) bool {
	fill := make(map[string]bool)
	empty := make(map[string]bool)

	copy := make(map[string]bool)
	for key, value := range *mp {
		copy[key] = value
	}

	for key, v := range *mp {
		x, y, z := intify3D(key)

		occupied := getOccupied(x, y, z, &copy)

		if v {
			if occupied == 2 || occupied == 3 {
				fill[key] = true
			} else {
				empty[key] = true
			}
		} else {
			if occupied == 3 {
				fill[key] = false
			}
		}
	}

	for s := range fill {
		copy[s] = true
	}
	for s := range empty {
		copy[s] = false
	}

	for key, value := range copy {
		(*mp)[key] = value
	}

	return len(fill)+len(empty) > 0
}

func day12(inp []string) {

	mp := make(map[string]bool)

	n := len(inp)

	for i := -1; i < n+1; i++ {
		for j := -1; j < n+1; j++ {
			for k := -1; k < n+1; k++ {
				mp[stringify3D(i, j, k)] = false
			}
		}
	}

	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp); j++ {
			if inp[i][j] == '#' {
				mp[stringify3D(i, j, 0)] = true
			} else {
				mp[stringify3D(i, j, 0)] = false
			}
		}
	}

	_print3D(mp)

	for j := 0; j < 6; j++ {
		doChanges3D(&mp)
	}

	fmt.Println(count(mp))
}

func _stringify2D(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
func _intify2D(s string) (x, y int) {
	fmt.Sscanf(s, "%d,%d", &x, &y)
	return
}

func _print3D(mp map[string]bool) {
	mp2 := make(map[int]map[string]bool)

	minX, minY := 1<<62, 1<<62
	maxX, maxY := -minX-1, -minY-1

	for key, v := range mp {
		x, y, z := intify3D(key)

		_, ok := mp2[z]
		if !ok {
			mp2[z] = make(map[string]bool)
		}
		mp2[z][_stringify2D(x, y)] = v
	}

	for _, v := range mp2 {
		for xy := range v {
			x, y := _intify2D(xy)

			if x < minX {
				minX = x
			}
			if y < minY {
				minY = y
			}

			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	// fmt.Println(minX, minY)
	// fmt.Println(maxX, maxY)

	for key, v := range mp2 {
		fmt.Println("z = ", key)

		plain := make([][]byte, maxX-minX+1)
		for i := 0; i < maxX-minX+1; i++ {
			plain[i] = make([]byte, maxY-minY+1)
		}

		for xy, vv := range v {
			x, y := _intify2D(xy)

			if vv {
				plain[x-minX][y-minY] = '#'
			} else {
				plain[x-minX][y-minY] = '.'
			}
		}

		for _, byteArr := range plain {
			fmt.Println(string(byteArr))
		}
	}
}

// =====================================

func intify4D(s string) (x, y, z, w int) {
	fmt.Sscanf(s, "%d,%d,%d,%d", &x, &y, &z, &w)
	return
}

func stringify4D(x, y, z, w int) string {
	return fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)
}

func day22(inp []string) {

	mp := make(map[string]bool)

	n := len(inp)

	for i := -1; i < n+1; i++ {
		for j := -1; j < n+1; j++ {
			for k := -1; k < n+1; k++ {
				for l := -1; l < n+1; l++ {
					mp[stringify4D(i, j, k, l)] = false
				}
			}
		}
	}

	for i := 0; i < len(inp); i++ {
		for j := 0; j < len(inp); j++ {
			if inp[i][j] == '#' {
				mp[stringify4D(i, j, 0, 0)] = true
			} else {
				mp[stringify4D(i, j, 0, 0)] = false
			}
		}
	}

	// _print3D(mp)

	for j := 0; j < 6; j++ {
		doChanges4D(&mp)
	}

	fmt.Println(count(mp))
}

var funcs = [](func(int) int){same, dec, inc}

func getOccupied4D(i, j, k, l int, mp *map[string]bool) int {
	occupied := 0

	for x, fI := range funcs {
		for y, fJ := range funcs {
			for z, fK := range funcs {
				for w, fL := range funcs {
					if x+y+z+w != 0 { // same,same,same,same,same
						occupied += findStuff4D(i, j, k, l, mp, fI, fJ, fK, fL)
					}
				}
			}
		}
	}

	// fmt.Println(occupied)

	return occupied
}

func findStuff4D(i, j, k, l int, mp *map[string]bool, updateI, updateJ, updateK, updateL func(int) int) int {
	i = updateI(i)
	j = updateJ(j)
	k = updateK(k)
	l = updateL(l)

	// fmt.Println(i, j, k, (*mp)[stringify3D(i, j, k)])
	v, ok := (*mp)[stringify4D(i, j, k, l)]
	if ok && v {
		return 1
	} else if !ok {
		(*mp)[stringify4D(i, j, k, l)] = false
	}

	return 0
}

func doChanges4D(mp *map[string]bool) bool {
	fill := make(map[string]bool)
	empty := make(map[string]bool)

	copy := make(map[string]bool)
	for key, value := range *mp {
		copy[key] = value
	}

	for key, v := range *mp {
		x, y, z, w := intify4D(key)

		occupied := getOccupied4D(x, y, z, w, &copy)

		if v {
			if occupied == 2 || occupied == 3 {
				fill[key] = true
			} else {
				empty[key] = true
			}
		} else {
			if occupied == 3 {
				fill[key] = false
			}
		}
	}

	for s := range fill {
		copy[s] = true
	}
	for s := range empty {
		copy[s] = false
	}

	for key, value := range copy {
		(*mp)[key] = value
	}

	return len(fill)+len(empty) > 0
}
