package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func Calculate(memory []int, inputs []int) {
// 	inputCounter := 0
// 	length := len(memory)
// 	for i := 0; i < length; {
// 		//fmt.Println(i, "===============================")
// 		opcode := fmt.Sprintf("%05d", memory[i])
// 		fmt.Println(i, opcode)
// 		if memory[i] == 99 {
// 			break
// 		}
// 		if memory[i] == 3 {
// 			//fmt.Println("Input", inputs[inputCounter], "at", memory[i+1], "replacing", memory[memory[i+1]])
// 			memory[memory[i+1]] = inputs[inputCounter]
// 			inputCounter++
// 			i += 2
// 		} else if memory[i] == 4 {
// 			for j := 0; j < len(memory); j++ {
// 				if memory[j] == 2808771 {
// 					fmt.Println(j)
// 				}
// 			}
// 			fmt.Println(i, memory[i], memory[i+1], memory[memory[i+1]])
// 			fmt.Println("output 1", "==>", memory[memory[i+1]])
// 			i += 2
// 		} else {
// 			if opcode[3:] == "04" {
// 				fmt.Println("output 2", "==>", memory[memory[i+1]])
// 				i += 2
// 				continue
// 			}

// 			//fmt.Println(opcode, memory[i+1], memory[i+2], memory[i+3])
// 			var op1, op2, ans int
// 			if opcode[2] == '0' {
// 				//fmt.Println("position mode", i, "firstPara")
// 				op1 = memory[memory[i+1]]
// 			} else {
// 				op1 = memory[i+1]
// 			}

// 			if opcode[1] == '0' {
// 				//fmt.Println("position mode", i, "secondPara")
// 				op2 = memory[memory[i+2]]
// 			} else {
// 				op2 = memory[i+2]
// 			}

// 			if opcode[3:] == "01" {
// 				ans = op1 + op2
// 				//fmt.Println("opcode", "01")
// 				memory[memory[i+3]] = ans
// 				i += 4
// 			} else if opcode[3:] == "02" {
// 				ans = op1 * op2
// 				memory[memory[i+3]] = ans
// 				i += 4
// 			} else if opcode[3:] == "05" {
// 				if op1 != 0 {
// 					i = op2
// 				} else {
// 					i += 3
// 				}

// 			} else if opcode[3:] == "06" {
// 				if op1 == 0 {
// 					i = op2
// 				} else {
// 					i += 3
// 				}
// 			} else if opcode[3:] == "07" {
// 				if op1 < op2 {
// 					memory[memory[i+3]] = 1
// 				} else {
// 					memory[memory[i+3]] = 0
// 				}
// 				i += 4
// 			} else if opcode[3:] == "08" {
// 				if op1 == op2 {
// 					memory[memory[i+3]] = 1
// 				} else {
// 					memory[memory[i+3]] = 0
// 				}
// 				i += 4
// 			} else {
// 				fmt.Println("ERROR", opcode, memory[i], memory, i)
// 				break
// 			}

// 		}

// 	}
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	str := ""
// 	for scanner.Scan() {
// 		str += scanner.Text()
// 	}

// 	if scanner.Err() != nil {
// 		// handle error.
// 	}
// 	//fmt.Println(str)
// 	strs := strings.Split(str, ",")
// 	memory := make([]int, len(strs))

// 	for i, stringNumber := range strs {
// 		intNumber, err := strconv.Atoi(stringNumber)
// 		if err != nil {
// 			panic(err)
// 		}
// 		memory[i] = intNumber
// 	}

// 	Calculate(memory, []int{5})
// }
