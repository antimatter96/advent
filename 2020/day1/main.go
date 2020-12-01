package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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

func day1(inp []string) {
	var ints []int
	var temp int
	for _, s := range inp {
		temp, _ = strconv.Atoi(s)
		ints = append(ints, temp)
	}

	x, y := sumTo2020(ints)

	fmt.Println(x, y, x*y, "")
}

func sumTo2020(nums []int) (int, int) {
	sort.Ints(nums)

	var sum int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum = nums[i] + nums[j]
			if sum == 2020 {
				return nums[i], nums[j]
			}
			if sum > 2020 {
				break
			}
		}
	}

	return -1, -1
}

func day2(inp []string) {
	var ints []int
	var temp int
	for _, s := range inp {
		temp, _ = strconv.Atoi(s)
		ints = append(ints, temp)
	}

	x, y, z := sumTo2020_2(ints)

	fmt.Println(x, y, z, x*y*z, "")
}
func sumTo2020_2(nums []int) (int, int, int) {
	sort.Ints(nums)

	var sum int

	var x, y int

	for i := 0; i < len(nums); i++ {
		x = nums[i]
		for j := i + 1; j < len(nums); j++ {
			y = nums[j]
			if x+y > 2020 {
				break
			}
			for k := j + 1; k < len(nums); k++ {
				sum = x + y + nums[k]
				if sum == 2020 {
					return x, y, nums[k]
				}
				if sum > 2020 {
					break
				}
			}
		}
	}

	return -1, -1, -1
}
