package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

const preamble = 25

func day(arr []string) {
	mp := make(map[int]int)

	nums := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		nums[i], _ = strconv.Atoi(arr[i])
	}

	for i := 0; i < preamble; i++ {
		mp[nums[i]]++
	}

	target := -1
	for i := preamble; i < len(arr); i++ {
		// fmt.Println(mp)

		if i-preamble-1 > -1 {
			mp[nums[i-preamble-1]]--
		}

		focus := nums[i]

		found := false
		for j := i - preamble; !found && j < i-1; j++ {
			x := nums[j]
			mp[x]--

			y := focus - x
			if mp[y] != 0 {
				found = true
			}
			mp[x]++
		}

		if !found {
			target = focus
			fmt.Println("ERR", focus)
			break
		}
		mp[focus]++
	}

	runningSum := 0
	// found := false
	var i, j int
	for i = 0; i < len(arr); i++ {
		runningSum = nums[i]

		j = i + 1
		for ; runningSum < target && j < len(arr); j++ {
			runningSum += nums[j]
		}

		if runningSum == target {
			fmt.Println(i, j, runningSum, nums[i], nums[j-1])

			temp := 0
			for index := i; index < j; index++ {
				temp += nums[index]
			}

			fmt.Println(temp)

			break
		}
	}

	max := -1
	min := 1 << 60

	for index := i; index < j; index++ {
		if nums[index] > max {
			max = nums[index]
		}

		if nums[index] < min {
			min = nums[index]
		}
	}

	fmt.Println(min, max, min+max)

}
