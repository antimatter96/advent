package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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

	day(inp)
}

func day(arr []string) {
	mp := make(map[int][]int)

	// [1518-11-01 00:00] Guard #10 begins shift
	// [1518-11-01 00:05] falls asleep
	// [1518-11-01 00:25] wakes up

	sort.Strings(arr)

	var month, day, hour, minutes int
	var gaurdID int
	var op int
	for _, s := range arr {
		s = strings.ReplaceAll(s, "-", " ")
		s = strings.ReplaceAll(s, ":", " ")
		s = strings.ReplaceAll(s, "#", "")

		// fmt.Println(s)
		if n, _ := fmt.Sscanf(s, "[1518 %d %d %d %d] Guard %d begins shift", &month, &day, &hour, &minutes, &gaurdID); n == 5 {
			op = 1
		} else if n, _ := fmt.Sscanf(s, "[1518 %d %d %d %d] falls asleep", &month, &day, &hour, &minutes); n == 4 {
			op = 2
		} else if n, _ := fmt.Sscanf(s, "[1518 %d %d %d %d] wakes up", &month, &day, &hour, &minutes); n == 4 {
			op = 3
		}

		if op == 2 && s[19] == 'w' {
			op = 3
		}

		if _, ok := mp[gaurdID]; !ok {
			mp[gaurdID] = make([]int, 60)
		}
		if op == 2 {
			// fmt.Println(gaurdID, "SLEEPS", hour, minutes)
			mp[gaurdID][minutes]++
		} else if op == 3 {
			// fmt.Println(gaurdID, "WAKES", hour, minutes)
			mp[gaurdID][minutes]--
		}

		// fmt.Println(op, gaurdID, month, day, hour, minutes)
		//fmt.Println(s, a, b, c, d, e)
	}

	// fmt.Println(mp)
	startergy2(mp)
}

func startergy1(mp map[int][]int) {
	maxID := 0
	maxSum := 0
	for id, arr := range mp {
		sum := 0

		for i := 0; i < 60; i++ {
			sum += arr[i]
			arr[i] = sum
		}

		sum = 0
		for i := 0; i < 60; i++ {
			sum += arr[i]
		}

		if sum > maxSum {
			maxSum = sum
			maxID = id
		}

		// fmt.Println(arr)
	}

	// fmt.Println(mp)

	//fmt.Println(maxID, maxSum)

	maxNights := 0
	maxMinute := 0
	for i, nights := range mp[maxID] {
		if nights > maxNights {
			maxNights = nights
			maxMinute = i
		}
	}

	//fmt.Println(maxNights, maxMinute)

	fmt.Println(maxID * maxMinute)
}

func startergy2(mp map[int][]int) {
	maxID := 0
	maxNights := 0
	maxMinute := 0

	for id, arr := range mp {
		sum := 0

		for i := 0; i < 60; i++ {
			sum += arr[i]
			arr[i] = sum
		}

		for i := 0; i < 60; i++ {
			if arr[i] > maxNights {
				maxNights = arr[i]
				maxID = id
				maxMinute = i
			}

		}
	}

	fmt.Println(maxID * maxMinute)
}
