package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inp := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp += scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}

	day(inp)
}

func day(input string) {
	min := len(input)

	var x int

	for i := 0; i < 26; i++ {
		arr := remove([]byte(input), byte(65+i))

		x = getLength(arr)

		if x < min {
			min = x
		}
	}

	fmt.Println(min)

}

func getLength(s []byte) int {
	var b bool
	var arr []byte

	for arr, b = doOnce(s); b; {
		arr, b = doOnce(arr)
	}

	return len(arr)
}

func remove(arr []byte, b byte) []byte {
	// fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == b || arr[i] == b+32 {
			// fmt.Println(arr[:i], arr[i+1:])
			arr = append(arr[:i], arr[i+1:]...)

			i--
		}
	}
	return arr
}

func doOnce(arr []byte) ([]byte, bool) {
	var toRet bool
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+32 == arr[i+1] || arr[i] == arr[i+1]+32 {
			toRet = true
			//fmt.Println(arr[:i], arr[i+2:])
			arr = append(arr[:i], arr[i+2:]...)

			i--
		}
	}

	// fmt.Println(arr)

	return arr, toRet
}
