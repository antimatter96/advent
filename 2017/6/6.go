package main

import (
	"fmt"
	"strconv"
)

var input []int = []int{4, 10, 4, 1, 8, 4, 9, 14, 5, 1, 14, 15, 0, 15, 3, 5}

func main() {
	res := countCycles2(input)
	fmt.Println(res)
}

var mp = make(map[string]bool)

func countCycles(arr []int) int {
	n := 0

	length := len(arr)
	for true {
		uniq := toString(arr)
		if mp[uniq] {
			break
		}
		mp[uniq] = true
		var high, pos int = -1, -1
		for i, v := range arr {
			if v > high {
				high = v
				pos = i
			}
		}

		arr[pos] = 0
		for i := 0; i < high; i++ {
			pos += 1
			pos %= length
			arr[pos]++
		}
		n++
	}

	return n
}

func toString(arr []int) string {
	var s string
	for _, v := range arr {
		s += strconv.Itoa(v)
		s += "_"
	}
	return s
}

var mp2 = make(map[string]int)

func countCycles2(arr []int) int {
	n := 1

	length := len(arr)
	for true {
		uniq := toString(arr)
		if mp2[uniq] > 0 {
			return n - mp2[uniq]
		}
		mp2[uniq] = n
		var high, pos int = -1, -1
		for i, v := range arr {
			if v > high {
				high = v
				pos = i
			}
		}

		arr[pos] = 0
		for i := 0; i < high; i++ {
			pos += 1
			pos %= length
			arr[pos]++
		}
		n++
	}

	return n
}
