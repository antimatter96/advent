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
		fmt.Errorf(scanner.Err().Error())
	}

	day22(inp)
}

var ROUNDS = 10000000
var NUMBER = 1000000

func day22(inp string) {
	inp = strings.TrimSpace(inp)

	ss := strings.Split(inp, "")

	temp, _ := strconv.Atoi(ss[0])
	head := &node{val: temp}
	mainMap[temp] = head
	ptr := head

	for i := 1; i < len(ss); i++ {
		temp, _ := strconv.Atoi(ss[i])
		ptr.next = &node{val: temp}
		mainMap[temp] = ptr.next
		ptr = ptr.next
	}

	for i := len(ss); i < NUMBER; i++ {
		ptr.next = &node{val: i + 1}
		mainMap[i+1] = ptr.next
		ptr = ptr.next
	}

	ptr.next = head

	run(head)
}

type node struct {
	val  int
	next *node
}

type ring struct {
	temp    []int
	tempMap map[int]bool
}

func print(head *node, highlight int) {
	ptr := head
	mp := make(map[int]bool)

	s := ""
	for ptr != nil {
		if mp[ptr.val] {
			break
		}

		if ptr.val == highlight {
			s += fmt.Sprintf("(%d) ", ptr.val)
		} else {
			s += fmt.Sprintf("%d ", ptr.val)
		}

		mp[ptr.val] = true

		ptr = ptr.next
	}

	fmt.Println(s)
}

func run(head *node) {
	temp := head

	x := ROUNDS / 20
	x++

	for k := 0; k < ROUNDS; k++ {
		if k%(x) == 0 {
			fmt.Println("-- move", k+1, "--")
		}

		startOfRemoved := remove(temp, 3)

		afterThis := findNext(temp, temp.val)

		add(afterThis, startOfRemoved, 3)

		// print(temp, -1)

		temp = temp.next
	}

	tt := mainMap[1]

	fmt.Println(tt.next.val * tt.next.next.val)
	// print(head, -1)
}

var tempMap = make(map[int]bool)
var tempArr = []int{1, 1, 1}

var mainMap = make(map[int]*node)

func remove(ptr *node, n int) *node {
	temp := ptr.next
	temp2 := ptr.next
	// temp := z

	tempMap = make(map[int]bool)

	for i := 0; i < n; i++ {
		tempArr[i] = temp.val
		tempMap[temp.val] = true

		temp = temp.next
	}

	ptr.next = temp

	return temp2
}

func findNext(head *node, n int) *node {
	// fmt.Println("next to ", n)
	k := n
	for k-1 > 0 {
		if _, v := tempMap[k-1]; !v {
			return mainMap[k-1]
		}
		k--
	}

	k = NUMBER
	for k > 0 {
		if _, v := tempMap[k]; !v {
			return mainMap[k]
		}
		k--
	}
	return nil
}

func findThis(head *node, n int) (ptr *node) {
	fmt.Println("finding", n)
	ptr = head

	for ptr != nil {
		if ptr.val == n {
			return ptr
		}
		ptr = ptr.next
	}

	return ptr
}

func add(where, addThese *node, n int) {
	temp := where.next

	where.next = addThese

	ptr := addThese
	for i := 0; i < n-1; i++ {
		ptr = ptr.next
	}
	ptr.next = temp
}
