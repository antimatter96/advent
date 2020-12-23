package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	// "github.com/emirpasic/gods/stacks/arraystack"
	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
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

	day44(inp)
}

func day22(inp string) {
	// tots := 0

	ss := strings.Split(inp, "\n\n")
	var temp int

	s1 := strings.Split(ss[0], "\n")
	// fmt.Println(s1[1:])

	l1 := dll.New()

	for _, v := range s1[1:] {
		temp, _ = strconv.Atoi(v)
		l1.Add(temp)
	}

	fmt.Println(l1)

	s2 := strings.Split(ss[1], "\n")
	// fmt.Println(s2[1:])
	l2 := dll.New()

	for _, v := range s2[1 : len(s2)-1] {
		temp, _ = strconv.Atoi(v)
		l2.Add(temp)
	}

	fmt.Println(l2)

	run(l1, l2)

	var nonEmpty *dll.List

	if l1.Empty() {
		nonEmpty = l2
	} else {
		nonEmpty = l1
	}

	tots := 0

	n := nonEmpty.Size()
	it := nonEmpty.Iterator()
	for it.Next() {
		_, value := it.Index(), it.Value()

		valueInt, _ := value.(int)
		tots += (valueInt * n)
		n--

	}

	fmt.Println(tots)
}

func day44(inp string) {
	// tots := 0

	ss := strings.Split(inp, "\n\n")
	var temp int

	s1 := strings.Split(ss[0], "\n")
	// fmt.Println(s1[1:])

	l1 := dll.New()

	for _, v := range s1[1:] {
		temp, _ = strconv.Atoi(v)
		l1.Add(temp)
	}

	fmt.Println(l1)

	s2 := strings.Split(ss[1], "\n")
	// fmt.Println(s2[1:])
	l2 := dll.New()

	for _, v := range s2[1 : len(s2)-1] {
		temp, _ = strconv.Atoi(v)
		l2.Add(temp)
	}

	fmt.Println(l2)

	game(1, l1, l2)

	var nonEmpty *dll.List

	if l1.Empty() {
		nonEmpty = l2
	} else {
		nonEmpty = l1
	}

	tots := 0

	n := nonEmpty.Size()
	it := nonEmpty.Iterator()
	for it.Next() {
		_, value := it.Index(), it.Value()

		valueInt, _ := value.(int)
		tots += (valueInt * n)
		n--

	}

	fmt.Println(tots)
}

func run(l1, l2 *dll.List) {
	var addToThis *dll.List

	var bigger, smaller *int

	for l1.Empty() && !l2.Empty() {
		m, _ := l1.Get(0)
		n, _ := l2.Get(0)

		l1.Remove(0)
		l2.Remove(0)

		mInt := m.(int)
		nInt := n.(int)

		if mInt > nInt {
			addToThis = l1
			bigger = &mInt
			smaller = &nInt
		} else {
			addToThis = l2
			bigger = &nInt
			smaller = &mInt
		}

		addToThis.Add(*bigger, *smaller)
	}
}

func recursiveRound(l1, l2 *dll.List, mp1, mp2 map[string]bool, roundNo, gameNo int) int {
	s1 := fmt.Sprint(l1)
	if mp1[s1] {
		return -1
	}
	mp1[s1] = true

	s2 := fmt.Sprint(l2)
	if mp2[s2] {
		return -1
	}
	mp2[s2] = true

	m, _ := l1.Get(0)
	n, _ := l2.Get(0)

	l1.Remove(0)
	l2.Remove(0)

	mInt := m.(int)
	nInt := n.(int)

	if mInt <= l1.Size() && nInt <= l2.Size() {
		newL1 := extractNewList(l1, mInt)
		newL2 := extractNewList(l2, nInt)

		x := game(gameNo+1, newL1, newL2)
		if x == -1 || x == 1 {
			l1.Add(mInt, nInt)
		} else {
			l2.Add(nInt, mInt)
		}

		return 0
	}

	var toRet = 0
	var addToThis *dll.List
	var bigger, smaller *int

	if mInt > nInt {
		addToThis = l1
		bigger = &mInt
		smaller = &nInt
		toRet = 1
	} else {
		addToThis = l2
		bigger = &nInt
		smaller = &mInt
		toRet = 2
	}

	addToThis.Add(*bigger, *smaller)

	return toRet
}

func game(gameNo int, l1, l2 *dll.List) int {
	mp1 := make(map[string]bool)
	mp2 := make(map[string]bool)

	roundNo := 1

	var x int

	// fmt.Printf("\n\n=== Game %d ===\n\n", gameNo)
	for !l1.Empty() && !l2.Empty() {
		// fmt.Printf("-- Round %d (Game %d) --\n", roundNo, gameNo)
		// fmt.Println(l1)
		// fmt.Println(l2)
		x = recursiveRound(l1, l2, mp1, mp2, roundNo, gameNo)
		if x < 0 {
			break
		}
		// fmt.Println("\n")

		roundNo++
	}

	return x
}

func extractNewList(l1 *dll.List, n int) *dll.List {
	newList := dll.New()
	it := l1.Iterator()
	for it.Next() {
		i, value := it.Index(), it.Value()

		if i+1 > n {
			break
		}

		valueInt, _ := value.(int)
		newList.Add(valueInt)

	}
	return newList
}
