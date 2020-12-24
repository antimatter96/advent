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

	day2(inp)
}

// N is
const N = 1000

func day1(ss []string) {
	arr = make([][]int, N)
	for i := 0; i < N; i++ {
		arr[i] = make([]int, N)
	}

	for _, s := range ss {
		var a, b, c, d, e int
		fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &a, &b, &c, &d, &e)
		mark(b, c, d, e)
		//fmt.Println(s, a, b, c, d, e)
	}

	fmt.Println(count(2))
	//printArr()

	//fmt.Println(twos * threes)
}

func printArr() {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}

var arr [][]int

func mark(y, x, width, height int) {
	for i := x; i < x+height; i++ {
		for j := y; j < y+width; j++ {
			arr[i][j]++
		}
	}
}

func count(limit int) int {
	tots := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if arr[i][j] >= limit {
				tots++
			}
		}
	}
	return tots
}

var arr2 [][]map[int]bool

func day2(ss []string) {
	arr2 = make([][]map[int]bool, N)

	for i := 0; i < N; i++ {
		arr2[i] = make([]map[int]bool, N)
		for j := 0; j < N; j++ {
			arr2[i][j] = make(map[int]bool, N)
		}
	}

	for _, s := range ss {
		var a, b, c, d, e int
		fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &a, &b, &c, &d, &e)
		mark2(a, b, c, d, e)
		//fmt.Println(s, a, b, c, d, e)
	}

	fmt.Println(count2(len(ss)))
	//printArr()

	//fmt.Println(twos * threes)
}

func mark2(id, y, x, width, height int) {
	for i := x; i < x+height; i++ {
		for j := y; j < y+width; j++ {
			arr2[i][j][id] = true
		}
	}
}

func count2(limit int) int {
	cnt := make(map[int]bool)

	for i := 1; i < limit+1; i++ {
		cnt[i] = true
	}

	tots := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			mp := arr2[i][j]
			if len(mp) > 1 {
				for k, _ := range mp {
					cnt[k] = false
				}
			}
		}
	}

	for i := 1; i < limit; i++ {
		if cnt[i] {
			return i
		}
	}

	//fmt.Println(cnt)
	return tots
}
