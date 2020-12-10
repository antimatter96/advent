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

	day(inp)
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
