package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rawInput := takeInput()

	fmt.Println(Run(rawInput))
}

func takeInput() []string {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	return inp
}

func parsePart1(inp []string) []string {
	return inp
}

func parsePart2(inp []string) []string {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(inp []string) int {
	noOfBits := len(inp[0])

	finalBitsGamma := strings.Builder{}
	finalBitsEpsilon := strings.Builder{}

	for i := 0; i < noOfBits; i++ {
		noOfOnesMoreThanZeros := 0

		for _, number := range inp {
			if number[i] == '1' {
				noOfOnesMoreThanZeros++
			} else {
				noOfOnesMoreThanZeros--
			}
		}

		if noOfOnesMoreThanZeros > 0 {
			finalBitsGamma.WriteByte('1')
			finalBitsEpsilon.WriteByte('0')
		} else {
			finalBitsGamma.WriteByte('0')
			finalBitsEpsilon.WriteByte('1')
		}
	}

	fmt.Println(finalBitsGamma.String())
	fmt.Println(finalBitsEpsilon.String())
	var gamma, epsilon int

	fmt.Sscanf(finalBitsGamma.String(), "%b", &gamma)
	fmt.Sscanf(finalBitsEpsilon.String(), "%b", &epsilon)

	ans := gamma * epsilon
	return ans
}

func Part2(inp []string) int {
	noOfBits := len(inp[0])

	pointers := make([]*string, 0)
	for i := 0; i < len(inp); i++ {
		pointers = append(pointers, &inp[i])
	}

	dup := make([]*string, len(inp))
	copy(dup, pointers)

	for i := 0; i < noOfBits; i++ {
		if len(pointers) == 1 {
			break
		}
		noOfOnesMoreThanZeros := 0

		startingWithOnes := make([]*string, 0)
		startingWithZeros := make([]*string, 0)

		for _, number := range pointers {
			if (*number)[i] == '1' {
				noOfOnesMoreThanZeros++
				startingWithOnes = append(startingWithOnes, number)
			} else {
				noOfOnesMoreThanZeros--
				startingWithZeros = append(startingWithZeros, number)
			}
		}

		if noOfOnesMoreThanZeros >= 0 {
			pointers = startingWithOnes
		} else {
			pointers = startingWithZeros
		}
	}

	for i := 0; i < noOfBits; i++ {
		if len(dup) == 1 {
			break
		}
		noOfOnesMoreThanZeros := 0

		startingWithOnes := make([]*string, 0)
		startingWithZeros := make([]*string, 0)

		for _, number := range dup {
			if (*number)[i] == '1' {
				noOfOnesMoreThanZeros++
				startingWithOnes = append(startingWithOnes, number)
			} else {
				noOfOnesMoreThanZeros--
				startingWithZeros = append(startingWithZeros, number)
			}
		}

		if noOfOnesMoreThanZeros >= 0 {
			dup = startingWithZeros
		} else {
			dup = startingWithOnes
		}
	}

	fmt.Println(*pointers[0])
	fmt.Println(*dup[0])

	var gamma, epsilon int
	fmt.Sscanf(*pointers[0], "%b", &gamma)
	fmt.Sscanf(*dup[0], "%b", &epsilon)

	ans := gamma * epsilon
	return ans
}
