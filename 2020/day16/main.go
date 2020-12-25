package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getickets(unprocessed []string) [][]int {
	var tickets [][]int

	var temp []int
	for i := 1; i < len(unprocessed); i++ {
		temp = geticket(unprocessed[i])

		if len(temp) > 0 {
			tickets = append(tickets, temp)
		}
	}

	return tickets
}

func geticket(unprocessed string) []int {
	if len(unprocessed) == 0 {
		return []int{}
	}
	s := strings.Split(unprocessed, ",")

	ticket := make([]int, len(s))

	var temp int
	for i, v := range s {
		temp, _ = strconv.Atoi(v)
		ticket[i] = temp
	}

	return ticket
}

// [lo, high]  [lo high]

type fieldStruct struct {
	name string
	range1Low, range1High,
	range2Low, range2High int
}

func getFieldRanges(unparsed []string) []fieldStruct {
	var fields []fieldStruct

	for _, v := range unparsed {
		temp := fieldStruct{}

		x := strings.Split(v, ":")
		_, err := fmt.Sscanf(x[1], "%d-%d or %d-%d ", &temp.range1Low, &temp.range1High, &temp.range2Low, &temp.range2High)
		if err != nil {
			fmt.Println(err)
		}
		temp.name = x[0]

		fields = append(fields, temp)
	}

	return fields
}

func isValid(value int, field fieldStruct) bool {
	return (value >= field.range1Low && value <= field.range1High) ||
		(value >= field.range2Low && value <= field.range2High)
}

func main() {
	var inp string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp += scanner.Text()
		inp += "\n"
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	day1(inp)
}

func day1(unparsed string) {
	raw := strings.Split(unparsed, "\n\n")

	myTicketString := strings.Split(raw[1], "your ticket:\n")[1]
	fieldsStrings := strings.Split(raw[0], "\n")
	ticketsStrings := strings.Split(raw[2], "\n")

	tickets := (getickets(ticketsStrings))
	fields := (getFieldRanges(fieldsStrings))
	myTicket := geticket(myTicketString)

	var falseValues []int
	var goodTickets [][]int

	for _, ticket := range tickets {
		allClear := true
		for _, fieldValue := range ticket {
			atLeastOne := false
			for _, f := range fields {
				if isValid(fieldValue, f) {
					atLeastOne = true
					break
				}
			}

			if !atLeastOne {
				falseValues = append(falseValues, fieldValue)
				allClear = false
				break
			}
		}

		if allClear {
			goodTickets = append(goodTickets, ticket)
		}
	}

	sum := 0

	for _, v := range falseValues {
		sum += v
	}

	fmt.Println(sum)
	fmt.Println(len(goodTickets), len(tickets))

	orderOfField := process2(goodTickets, fields)

	tots := uint64(1)

	// fmt.Println(orderOfField)
	// fmt.Println(myTicket)

	// for k, v := range orderOfField {
	// 	fmt.Println(fields[k].name, v)
	// }

	fmt.Println()

	// for _, v := range orderOfField {
	// 	// field := fields[k]
	// 	// fmt.Println(myTicket[v], "(", field.range1Low, field.range1High, ") (", field.range2Low, field.range2High, ")")
	// 	tots *= uint64(myTicket[v])
	// 	fmt.Println(tots)
	// }
	// fmt.Print("tots", tots)

	tots = 1
	for i := 0; i < 6; i++ {
		tots *= uint64(myTicket[orderOfField[i]])
	}

	fmt.Println("==>", tots, "<==")

	// VERIFY
	for k, v := range orderOfField {
		field := fields[k]
		for i, ticket := range goodTickets {
			val := ticket[v]

			if isValid(val, field) {
			} else {
				panic(fmt.Sprintf("ticket %d field %d  %d %v ", i, v, val, field))
			}
		}
	}

}

func printThis(thisFieldCanTakeIndex [][]bool, fields []fieldStruct, extraCheck bool) {
	fmt.Println("========================================================")
	var s string
	for i, v := range thisFieldCanTakeIndex {
		fmt.Printf("%20s ", fields[i].name)

		for _, vv := range v {

			if vv {
				s = "X"
			} else {
				s = " "
			}

			fmt.Printf("%-2v ", s)
		}

		fmt.Println()
	}
	fmt.Println("========================================================")
	if extraCheck {
		mp := make(map[int]bool)
		for i, v := range thisFieldCanTakeIndex {
			for j, vv := range v {
				if vv {
					if mp[j] {
						panic(fmt.Sprintf("%s  %d ", fields[i].name, j))
					}
					mp[j] = true
				}
			}
		}
	}

	fmt.Println("========================================================")
}

func process2(tickets [][]int, fields []fieldStruct) []int {
	n := len(fields)

	// [i][j]
	// => i is order in fieldsArray
	// => j is order in the ticket

	isPossible := make([][]bool, len(fields))
	for i := 0; i < n; i++ {
		isPossible[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			isPossible[i][j] = true
		}
	}

	// printThis(isPossible, fields, false)

	// For all tickets, see all values
	// for each value [j], if out of range for a field [i]
	// isPossible[i][j] = false

	for _, ticket := range tickets {
		for j, value := range ticket {
			for i, field := range fields {
				if !isValid(value, field) {
					isPossible[i][j] = false
				}
			}
		}
	}

	// printThis(isPossible, fields, false)

	// For all fields [i]
	// Check specific value of [j] tickets  =>> if out of range for a field [i]
	// break and
	// isPossible[i][j] = false

	for i, field := range fields {
		for j := 0; j < n; j++ {
			//broke := false
			for _, ticket := range tickets {
				if !isValid(ticket[j], field) {
					isPossible[i][j] = false
				}
			}
		}
	}

	// printThis(isPossible, fields, false)

	duplicate := make([][]bool, len(isPossible))
	for i := range isPossible {
		duplicate[i] = make([]bool, len(isPossible[i]))
		copy(duplicate[i], isPossible[i])
	}

	mp := make(map[int]bool)

	reduce(isPossible, n, mp)

	// printThis(isPossible, fields, true)

	mp = make(map[int]bool)

	reduce(duplicate, n, mp)

	// printThis(duplicate, fields, true)

	//
	//
	//

	order := make([]int, n)

	for i, vv := range isPossible {
		for j, v := range vv {
			if v {
				order[i] = j
				break
			}
		}
	}

	fmt.Println(order)

	return order
}

// func process(tickets [][]int, fields []fieldStruct) []int {
// 	n := len(fields)
// 	thisFieldCanTakeIndex := make([][]bool, n)

// 	for i := 0; i < n; i++ {
// 		thisFieldCanTakeIndex[i] = make([]bool, n)
// 		for j := 0; j < n; j++ {
// 			thisFieldCanTakeIndex[i][j] = true
// 		}
// 	}

// 	// printThis(thisFieldCanTakeIndex, fields, false)

// 	for _, ticket := range tickets {
// 		for index, fieldValue := range ticket {
// 			for fieldIndex, f := range fields {
// 				if isValid(fieldValue, f) {
// 				} else {
// 					// this field can't take this value pakka se
// 					thisFieldCanTakeIndex[fieldIndex][index] = false
// 				}
// 			}
// 		}
// 		// fmt.Println(ticket, thisTicketHasThisTicketValidForThisNoOfFields)
// 	}

// 	// printThis(thisFieldCanTakeIndex, fields, false)

// 	for indexOfCurrentField, f := range fields {
// 		thisFieldIsValidForTheseIndexesForThisNoOfTickets := make(map[int]int)

// 		for i := 0; i < len(tickets[0]); i++ {
// 			for _, ticket := range tickets {
// 				fieldValue := ticket[i]

// 				if isValid(fieldValue, f) {
// 					thisFieldIsValidForTheseIndexesForThisNoOfTickets[i]++
// 				}
// 			}
// 		}

// 		for k, v := range thisFieldIsValidForTheseIndexesForThisNoOfTickets {
// 			if v != len(tickets) {
// 				thisFieldCanTakeIndex[indexOfCurrentField][k] = false
// 			}
// 		}

// 		// fmt.Println(f.name, thisFieldIsValidForTheseIndexesForThisNoOfTickets)
// 	}

// 	// printThis(thisFieldCanTakeIndex, fields, false)

// 	mp := make(map[int]bool)

// 	reduce(thisFieldCanTakeIndex, n, mp)

// 	// printThis(thisFieldCanTakeIndex, fields, false)

// 	order := make([]int, n)

// 	tota := 0
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			if thisFieldCanTakeIndex[i][j] {
// 				tota++
// 				order[i] = j
// 				break
// 			}
// 		}
// 	}

// 	fmt.Println(">>>>>", tota)
// 	return order
// }

// func reduce(arr [][]bool, n int, thisIsFinal map[int]bool) {
// 	// check if this field fit where all

// 	for i := 0; i < n; i++ {
// 		if thisIsFinal[i] {
// 			continue
// 		}

// 		noOfIndexWhereThisIsValid := make(map[int]int)
// 		for j := 0; j < n; j++ {
// 			if arr[i][j] {
// 				noOfIndexWhereThisIsValid[j]++
// 			}
// 		}

// 		// REMOVE THIS FROM ALL OTHERS
// 		if len(noOfIndexWhereThisIsValid) == 1 {
// 			for indexWhichThisFieldCanTake := range noOfIndexWhereThisIsValid {
// 				for i2 := 0; i2 < n; i2++ {
// 					if i2 != i {
// 						arr[i2][indexWhichThisFieldCanTake] = false
// 					}
// 				}
// 			}
// 			thisIsFinal[i] = true

// 			reduce(arr, n, thisIsFinal)
// 		} else {
// 			// fmt.Println("haWw")
// 		}
// 	}
// }

func reduce(isPossible [][]bool, n int, thisIsFinal map[int]bool) {
	// For each field [i] see if only 1 [j] is true
	// It it is => mark i as true
	// For all other fields, mark j as false

	for i := 0; i < n; i++ {
		if thisIsFinal[i] {
			continue
		}

		var temp []int
		for j := 0; j < n; j++ {
			if isPossible[i][j] {
				temp = append(temp, j)
			}
		}

		if len(temp) == 1 {
			thisIsFinal[i] = true

			for i2 := 0; i2 < n; i2++ {
				isPossible[i2][temp[0]] = false
			}

			isPossible[i][temp[0]] = true

			reduce(isPossible, n, thisIsFinal)
		}

	}

}
