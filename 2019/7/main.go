package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculate(memory []int, inputs []int, start int) (int, int, bool) {

	inputCounter := 0
	outputCounter := 0
	var output int
	length := len(memory)

	for i := start; i < length; {
		//<-timer1.C
		opcode := fmt.Sprintf("%05d", memory[i])
		//fmt.Println(i, opcode)
		switch opcode[3:] {
		case "99":
			i++
			return -1, i, true
		case "03":
			memory[memory[i+1]] = inputs[inputCounter]
			//fmt.Println("Input received", memory[memory[i+1]])
			inputCounter++
			i += 2
			continue
		case "04":
			outputCounter++
			output = memory[memory[i+1]]
			i += 2
			return output, i, false
			continue
		}

		var op1, op2 int
		if opcode[2] == '0' {
			op1 = memory[memory[i+1]]
		} else {
			op1 = memory[i+1]
		}

		if opcode[1] == '0' {
			op2 = memory[memory[i+2]]
		} else {
			op2 = memory[i+2]
		}

		switch opcode[3:] {
		case "01":
			memory[memory[i+3]] = op1 + op2
			i += 4
		case "02":
			memory[memory[i+3]] = op1 * op2
			i += 4
		case "05":
			if op1 != 0 {
				i = op2
			} else {
				i += 3
			}
		case "06":
			if op1 == 0 {
				i = op2
			} else {
				i += 3
			}
		case "07":
			if op1 < op2 {
				memory[memory[i+3]] = 1
			} else {
				memory[memory[i+3]] = 0
			}
			i += 4
		case "08":
			if op1 == op2 {
				memory[memory[i+3]] = 1
			} else {
				memory[memory[i+3]] = 0
			}
			i += 4
		}
	}

	fmt.Println("LOOP END")
	return output, length, false
}

func main() {
	//timer1 := time.NewTimer(10 * time.Second)
	scanner := bufio.NewScanner(os.Stdin)
	str := ""
	for scanner.Scan() {
		str += scanner.Text()
	}

	if scanner.Err() != nil {
		// handle error.
	}
	//fmt.Println(str)
	strs := strings.Split(str, ",")
	memory := make([]int, len(strs))

	zero := make([]int, len(strs))

	for i, stringNumber := range strs {
		intNumber, err := strconv.Atoi(stringNumber)
		if err != nil {
			panic(err)
		}
		zero[i] = intNumber
	}

	reset(&memory, zero)
	maxSetting := ""

	maxS := 0

	// for a := 0; a < 5; a++ {
	// 	reset(&memory, zero)
	// 	out1 := Calculate(memory, []int{a, 0})
	// 	for b := 0; b < 5; b++ {
	// 		reset(&memory, zero)
	// 		out2 := Calculate(memory, []int{b, out1})
	// 		for c := 0; c < 5; c++ {
	// 			reset(&memory, zero)
	// 			out3 := Calculate(memory, []int{c, out2})
	// 			for d := 0; d < 5; d++ {
	// 				reset(&memory, zero)
	// 				out4 := Calculate(memory, []int{d, out3})
	// 				for e := 0; e < 5; e++ {
	// 					reset(&memory, zero)
	// 					out5 := Calculate(memory, []int{e, out4})
	// 					if out5 >= maxS {
	// 						maxS = out5
	// 						maxSetting = fmt.Sprintf("%d,%d,%d,%d,%d", a, b, c, d, e)
	// 					}
	// 				}
	// 			}

	// 		}

	// 	}

	// }
	outp := make(chan []int, 1)
	go mai2n(outp)

	for patt := range outp {
		fmt.Println(patt)

		pp := 0
		x := true

		rr := make([][]int, len(patt))
		zz := make([]int, len(patt))
		for i := 0; i < len(patt); i++ {
			rr[i] = make([]int, len(zero))
			copy(rr[i], zero)
		}
		asd := 0
		j := 0
		var temp int
	Loop:
		for {
			for i, set := range patt {
				y := pp
				if j == 0 {
					temp, asd, x = Calculate(rr[i], []int{set, y}, zz[i])
					fmt.Println([]int{set, y}, pp, x, asd)
				} else {
					temp, asd, x = Calculate(rr[i], []int{y}, zz[i])
					fmt.Println(y, pp, x, asd, i)
				}

				zz[i] = asd
				//fmt.Println(x, pp)
				if x {
					fmt.Println("Arpit")
					break Loop
				}
				pp = temp

			}
			j++
		}

		if pp > maxS {
			maxS = pp
		}
	}

	fmt.Println(maxS, maxSetting)
	//<-timer1.C
}

func reset(memory *[]int, zero []int) {
	for i, v := range zero {
		(*memory)[i] = v
	}
}

/*
procedure generate(n : integer, A : array of any):
    //c is an encoding of the stack state. c[k] encodes the for-loop counter for when generate(k+1, A) is called
    c : array of int

    for i := 0; i < n; i += 1 do
        c[i] := 0
    end for

    output(A)

    //i acts similarly to the stack pointer
    i := 0;
    while i < n do
        if  c[i] < i then
            if i is even then
                swap(A[0], A[i])
            else
                swap(A[c[i]], A[i])
            end if
            output(A)
            //Swap has occurred ending the for-loop. Simulate the increment of the for-loop counter
            c[i] += 1
            //Simulate recursive call reaching the base case by bringing the pointer to the base case analog in the array
            i := 0
        else
            //Calling generate(i+1, A) has ended as the for-loop terminated. Reset the state and simulate popping the stack by incrementing the pointer.
            c[i] := 0
            i += 1
        end if
		end while
*/

func mai2n(out chan []int) {
	var a = []int{5, 6, 7, 8, 9}

	var n = len(a) - 1
	var i, j int
	outp := make([]int, n+1)
	copy(outp, a)
	out <- outp
	for c := 1; c < 120; c++ { // 3! = 6:
		i = n - 1
		j = n
		for a[i] > a[i+1] {
			i--
		}
		for a[j] < a[i] {
			j--
		}
		a[i], a[j] = a[j], a[i]
		j = n
		i++
		for i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
		outp := make([]int, n+1)
		copy(outp, a)
		out <- outp
	}
	close(out)
}
