package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "github.com/emirpasic/gods/stacks/arraystack"
)

func main() {
	var inp string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp += scanner.Text()
		inp += "\n"
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	day22(inp)
}

func day22(inp string) {
	inp = strings.TrimSpace(inp)

	ss := strings.Split(inp, "")
	r := ring{}
	r.arr = make([]int, len(ss))

	for i, v := range ss {
		temp, _ := strconv.Atoi(v)
		r.arr[i] = temp
	}

	r.temp = make([]int, 3)
	r.tempMap = make(map[int]bool)
	run(&r)
}

type ring struct {
	arr     []int
	temp    []int
	tempMap map[int]bool
}

func run(r *ring) {
	i := 0

	for k := 0; k < 100; k++ {

		// fmt.Println("Current", r.arr[i])
		// fmt.Println("cups:", print(r.arr, r.arr[i]))

		thisWas := r.arr[i]
		r.remove(i, 3)

		// fmt.Println(r.arr)
		// fmt.Println(r)

		jj := r.findNext(r.arr[i])
		theNext := r.arr[jj]

		// fmt.Println(r.arr[i], jj)

		r.j()
		j := r.findThis(theNext)
		// fmt.Println(r.temp)
		// fmt.Println("destination:", r.arr[j])
		r.fill(j)

		i = r.findThis(thisWas)
		i++
		i = i % len(r.arr)

		// fmt.Println()
	}

	fmt.Println(r)
}

func print(arr []int, highlight int) string {
	s := ""

	for _, v := range arr {
		if v == highlight {
			s += fmt.Sprintf("(%d) ", v)
		} else {
			s += fmt.Sprintf("%d ", v)
		}
	}
	return s
}

func (r *ring) remove(afterIndex, n int) {
	var i = afterIndex

	r.tempMap = make(map[int]bool)

	var index int
	for j := 0; j < n; j++ {
		index = (i + 1 + j) % len(r.arr)
		r.temp[j] = r.arr[index]
		r.arr[index] = -1
		r.tempMap[r.temp[j]] = true
	}

	// fmt.Println(r.tempMap)
}

func (r *ring) findNext(n int) int {
	// fmt.Println("next to ", n)
	k := n
	for k-1 > 0 {
		if _, v := r.tempMap[k-1]; !v {
			return r.findThis(k - 1)
		}
		k--
	}
	// fmt.Println("roll over ")
	k = len(r.arr)
	for k > 0 {
		if _, v := r.tempMap[k]; !v {
			return r.findThis(k)
		}
		k--
	}
	return 0
}

func (r *ring) findThis(n int) (index int) {
	index = 0

	for i, v := range r.arr {
		if v == n {
			index = i
			return
		}
	}

	return
}

var temp = make([]int, 9)

func (r *ring) j() {
	i := 0
	for _, v := range r.arr {
		if v == -1 {
			continue
		}
		temp[i] = v
		i++
	}
	r.arr = temp
}

func (r *ring) fill(index int) {
	n := len(r.arr)

	for i := n - 1; i-3 > index && i-3 > -1; i-- {
		r.arr[i] = r.arr[i-3]
		r.arr[i-3] = 0
	}

	for j := 0; j < 3; j++ {
		r.arr[index+1+j] = r.temp[j]
	}
}
