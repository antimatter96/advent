package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
	}

	tester(inp)
}

// const loops = 50

func tester(inps []string) {
	for _, v := range inps {
		fmt.Println(v, isValidByte([]byte(v)))
	}
}

func day1(inps []string) {
	inp := inps[0]

	for i := 0; i < len(inps); i++ {
		// inp = expand(inp)
	}

	fmt.Println(len(inp))
}

func next26(password []byte, n int) {

}

var intI = byte('i')
var intO = byte('o')
var intL = byte('l')

func isValidByte(password []byte) bool {
	if len(password) < 3 {
		return false
	}

	for i := 0; i < len(password); i++ {
		fmt.Printf("%d ", password[i])
		if password[i] == intO || password[i] == intI || password[i] == intL {
			fmt.Println("Has i o l")
			return false
		}
	}

	containsSuccessive := false
	for i := 0; i < len(password)-3; i++ {
		if password[i+1]-password[i] == 1 && password[i+2]-password[i+1] == 1 {
			containsSuccessive = true
			break
		}
	}
	if !containsSuccessive {
		fmt.Println("No Successive")
		return false
	}

	mp := make(map[byte]int)
	for i := 0; i < len(password)-1; i++ {
		if password[i+1] == password[i] {
			mp[password[i+1]]++
		}
	}

	if len(mp) < 2 {
		fmt.Println("No twiced")
		return false
	}

	return true
}

// 97 - 122
