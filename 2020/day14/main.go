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

	day2(inp)
}

func day1(unparsed []string) {
	var mask string
	var maskArr []byte

	var memIndex int
	var value int64

	mp := make(map[int]int64)

	for _, u := range unparsed {
		if n, _ := fmt.Sscanf(u, "mask = %s", &mask); n == 1 {
			maskArr = []byte(mask)
		} else if n, _ := fmt.Sscanf(u, "mem[%d] = %d", &memIndex, &value); n == 2 {
			fmt.Println(memIndex, value)
			mp[memIndex] = getResult(fmt.Sprintf("%036b", value), maskArr)
		}
	}

	tots := int64(0)

	for _, v := range mp {
		tots += v
	}
	fmt.Println(tots)
}

func day2(unparsed []string) {
	var mask string
	var maskArr []byte

	var memIndex int64
	var value int64

	mp := make(map[int64]int64)

	for _, u := range unparsed {
		if n, _ := fmt.Sscanf(u, "mask = %s", &mask); n == 1 {
			maskArr = []byte(mask)
		} else if n, _ := fmt.Sscanf(u, "mem[%d] = %d", &memIndex, &value); n == 2 {
			tempAddress := getTempAddress(fmt.Sprintf("%036b", memIndex), maskArr)

			saveThisAtThis(value, tempAddress, mp)
		}
	}

	tots := int64(0)

	for _, v := range mp {
		tots += v
	}
	fmt.Println(len(mp), tots)
}

func getResult(s string, maskArr []byte) int64 {
	sArr := []byte(s)
	for i := 0; i < len(sArr); i++ {
		if maskArr[i] != 'X' {
			sArr[i] = maskArr[i]
		}
	}

	ret, _ := strconv.ParseInt(string(sArr), 2, 64)
	return ret
}

func getTempAddress(s string, maskArr []byte) []byte {
	sArr := []byte(s)
	for i := 0; i < len(sArr); i++ {
		if maskArr[i] != '0' {
			sArr[i] = maskArr[i]
		}
	}
	return sArr
}

func saveThisAtThis(value int64, tempAddress []byte, mp map[int64]int64) {
	var indices []int
	for i, v := range tempAddress {
		if v == 'X' {
			indices = append(indices, i)
		}
	}

	k := len(indices)

	max := 1 << k

	var ret int64
	for i := 0; i < max; i++ {
		s := fmt.Sprintf(fmt.Sprintf("%%0%db", k), i)

		for index, v := range indices {
			tempAddress[v] = s[index]
		}

		ret, _ = strconv.ParseInt(string(tempAddress), 2, 64)
		mp[ret] = value
	}
}
