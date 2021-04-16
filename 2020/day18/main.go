package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/stacks/arraystack"
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

	day22(inp)
}

func day22(inp []string) {
	tots := 0
	for _, v := range inp {
		tots += process(v)
	}
	fmt.Println(tots)
}

func process(inp string) int {
	inp = strings.Replace(inp, "(", " ( ", -1)
	inp = strings.Replace(inp, ")", " ) ", -1)

	inpArray := strings.Split(inp, " ")
	fmt.Println("=========\n", inpArray)

	stack := arraystack.New()

	for _, v := range inpArray {
		if v == "" {
			continue
		}
		// fmt.Println("--")
		// fmt.Println(v, len(v))
		if v == "(" {
			stack.Push(v)
		} else if v == "*" || v == "+" {
			stack.Push(v)
		} else if v == ")" {

			stack.Push(v)
			pop3AndPush(stack)
			// if !x {

			// }
		} else {
			if stack.Empty() {
				stack.Push(v)
			} else {
				k, vv := stack.Peek()
				if vv {
					if k == "+" || k == "*" {
						pop2AndPush(stack, v)
					} else if k == "(" {
						stack.Push(v)
					}

					// fmt.Println(k)
				} else {
					stack.Push(v)
				}
			}

		}
		// fmt.Printf("%s => [%s]\n", v, strings.Replace(fmt.Sprintf("%s", stack), "ArrayStack\n", "", -1))
		reduce(stack)
		// fmt.Printf("%s => [%s]\n", v, strings.Replace(fmt.Sprintf("%s", stack), "ArrayStack\n", "", -1))
	}

	if stack.Size() != 1 {
		panic(fmt.Sprintf("=> [%s]\n\n", strings.Replace(fmt.Sprintf("%s", stack), "ArrayStack\n", "", -1)))
	}

	finalInterface, _ := stack.Peek()
	finalString, _ := finalInterface.(string)

	finalInt, _ := strconv.Atoi(finalString)

	return finalInt
	// fmt.Printf("%v\n\n", stack)
}

var specialCharacters = map[string]bool{
	"+": true,
	"*": true,
	"(": true,
	")": true,
}

func reduce(stack *arraystack.Stack) {
	for stack.Size() > 2 {
		topInterface, _ := stack.Peek()
		topString, _ := topInterface.(string)

		if !specialCharacters[topString] {
			changed := pop3AndPush(stack)
			if !changed {
				break
			}
		} else {
			break
		}

	}
}

func pop2AndPush(stack *arraystack.Stack, op2String string) {
	operation, _ := stack.Pop()
	operationString, _ := operation.(string)

	op1Interface, _ := stack.Pop()

	op2, _ := strconv.Atoi(op2String)

	op1string, _ := op1Interface.(string)
	op1, _ := strconv.Atoi(op1string)

	if operationString == "+" {
		stack.Push(fmt.Sprintf("%d", op1+op2))
	} else {
		stack.Push(fmt.Sprintf("%d", op1*op2))
	}
}

func pop3AndPush(stack *arraystack.Stack) bool {
	// fmt.Println("pop3AndPush")
	op2Interface, _ := stack.Pop()
	op2String, _ := op2Interface.(string)

	operation, _ := stack.Pop()
	operationString, _ := operation.(string)

	if operationString == "(" {
		stack.Push(operationString)
		stack.Push(op2String)
		return false
	}

	op2, _ := strconv.Atoi(op2String)
	op1Interface, _ := stack.Pop()
	op1string, _ := op1Interface.(string)
	op1, _ := strconv.Atoi(op1string)

	if op2String == ")" && op1string == "(" {
		stack.Push(operationString)
		return true
	}

	// fmt.Printf("%d %s %d\n", op1, operationString, op2)
	if operationString == "+" {
		stack.Push(fmt.Sprintf("%d", op1+op2))
	} else {
		stack.Push(fmt.Sprintf("%d", op1*op2))
	}

	return true
}
