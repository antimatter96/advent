package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)


func calculate(memory []int) {
	length := len(memory)
	for i := 0; i < length; i+=4 {
		if memory[i] == 99 {
			break
		}
		if memory[i] == 1 {
			memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
		} else if memory[i] == 2 {
			memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
		}

	}
}



// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	str := ""
// 	for scanner.Scan() {
// 		str +=scanner.Text()
// 	}

// 	if scanner.Err() != nil {
// 		// handle error.
// 	}
// 	fmt.Println(str)
// 	strs := strings.Split(str, ",")
// 	memory := make([]int, len(strs))

// 	for i, stringNumber := range strs {
//     intNumber, err := strconv.Atoi(stringNumber)
//     if err != nil {
//       panic(err)
//     }
//     memory[i] = intNumber
// 	}

// 	memory[1] = 12
// 	memory[2] = 2

// 	calculate(memory)
// 	fmt.Println(memory[0])
// }


func main2() {
	scanner := bufio.NewScanner(os.Stdin)
	str := ""
	for scanner.Scan() {
		str +=scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}
	fmt.Println(str)
	strs := strings.Split(str, "-")
	rangeLow,_ := strconv.Atoi(strs[0])
	rangeHigh,_ := strconv.Atoi(strs[1])
	fmt.Println(rangeLow, rangeHigh)
	tots := 0
	for i:= rangeLow; i < rangeHigh; i++ {
		number := strconv.Itoa(i)
		
		adjacent := false
		increasing := true
		for j := 0; j < len(number)-1; j++ {
			if(number[j] == number[j+1]) {
				adjacent = true
			}
			
			if number[j] > number[j+1] {
				increasing = false
				break
			}
		}
		
		if adjacent && increasing {
			tots++
		}

	}
	fmt.Println(tots)
}


func main() {
        scanner := bufio.NewScanner(os.Stdin)
        str := ""
        for scanner.Scan() {
                str +=scanner.Text()
        }

        if scanner.Err() != nil {
                // handle error.
        }
        fmt.Println(str)
        strs := strings.Split(str, "-")
        rangeLow,_ := strconv.Atoi(strs[0])
        rangeHigh,_ := strconv.Atoi(strs[1])
        fmt.Println(rangeLow, rangeHigh)
        tots := 0
        for i:= rangeLow; i < rangeHigh; i++ {
                number := strconv.Itoa(i)

                adjacent := false
                increasing := true
                for j := 0; j < len(number)-1; j++ {
                        if number[j] > number[j+1] {
                                increasing = false
                                break
                        }
                }
		eqlCount := 1;
		for j := 1; j < len(number); j++ {
                        if number[j] == number[j-1] {
                               	eqlCount++
                        } else if eqlCount == 2{
				adjacent = true
				eqlCount = 1
			} else {
				eqlCount = 1
			}
                }
		if eqlCount == 2{
                     adjacent = true
		}
                if adjacent && increasing {
                        tots++
                }

        }
        fmt.Println(tots)
}	
