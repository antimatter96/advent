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

const preamble = 25

func day(arr []string) {
	mp := make(map[int]int)

	nums := make([]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		nums[i], _ = strconv.Atoi(arr[i])
	}

	sort.Ints(nums)

	nums = append(nums, nums[len(arr)]+3)

	//fmt.Println(nums)

	for i := 0; i < len(arr)+1; i++ {
		mp[nums[i+1]-nums[i]]++
	}

	//fmt.Println(mp, mp[1]*mp[3])
	fmt.Println(mp[1] * mp[3])
}

func day2(arr []string) {
	mp := make(map[int]int)

	nums := make([]int, len(arr)+1)

	for i := 0; i < len(arr); i++ {
		nums[i], _ = strconv.Atoi(arr[i])
	}

	sort.Ints(nums)

	howMany(0, nums, mp)
	fmt.Println(mp[0])
}

func howMany(index int, nums []int, cache map[int]int) int {
	// fmt.Println("checking", index, nums[index])
	if index == len(nums)-1 {
		// fmt.Println("Reached", index, nums[index])
		return 1
	}
	if res, ok := cache[index]; ok {
		// fmt.Println("Hit", nums[index], res)
		return res
	}

	temp := 0

	if index+3 < len(nums) && nums[index+3]-nums[index] <= 3 {
		temp += howMany(index+3, nums, cache)
	}
	if index+2 < len(nums) && nums[index+2]-nums[index] <= 3 {
		temp += howMany(index+2, nums, cache)
	}
	if index+1 < len(nums) && nums[index+1]-nums[index] <= 3 {
		temp += howMany(index+1, nums, cache)
	}

	cache[index] = temp

	return temp
}
