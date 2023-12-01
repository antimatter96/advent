package main

import (
	"fmt"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	fmt.Println("Answer 1", p1)
	fmt.Println("Answer 2", p2)
}

func parsePart1(inp []string) []string {
	same := make([]string, len(inp))
	copy(same, inp)
	return same
}

func parsePart2(inp []string) []string {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func extractActualValue(garbledCalibrationValue string) int {
	var actualCalibratedValue int

	for i := 0; i < len(garbledCalibrationValue); i++ {
		fmt.Println("[DEBUG]", garbledCalibrationValue[i])
		if garbledCalibrationValue[i] > 47 && garbledCalibrationValue[i] < 58 {
			actualCalibratedValue = (int(garbledCalibrationValue[i]) - 48) * 10
			break
		}
	}

	for i := len(garbledCalibrationValue) - 1; i > -1; i-- {
		fmt.Println("[DEBUG]", garbledCalibrationValue[i])
		if garbledCalibrationValue[i] > 47 && garbledCalibrationValue[i] < 58 {
			actualCalibratedValue += (int(garbledCalibrationValue[i]) - 48)
			break
		}
	}

	return actualCalibratedValue
}

func Part1(garbledCalibrationValues []string) int {
	var sum int

	for i := 0; i < len(garbledCalibrationValues); i++ {
		fmt.Println("[DEBUG]", "garbledCalibrationValues", garbledCalibrationValues[i])

		actualCalibratedValue := extractActualValue(garbledCalibrationValues[i])

		fmt.Println("[DEBUG]", "actualCalibratedValue", actualCalibratedValue)

		sum += actualCalibratedValue
	}

	return sum
}

func Part2(garbledCalibrationValues []string) int {
	replaceMap := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for i := 0; i < len(garbledCalibrationValues); i++ {

		for old, new := range replaceMap {
			garbledCalibrationValues[i] = strings.ReplaceAll(garbledCalibrationValues[i], old, old+new+old)
		}

	}

	return Part1(garbledCalibrationValues)
}
