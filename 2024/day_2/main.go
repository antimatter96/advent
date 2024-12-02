package main

import (
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2024/common"
)

const (
	ORDER_INC = 1
	ORDER_DEC = 2
	//
	ORDER_IND          = 0
	ORDER_NOT_IN_ORDER = -1
)

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) [][]int {
	common.Log.Debug().Int("Input Lengtg", len(inp)).Send()

	reports := make([][]int, len(inp))

	for i, reportS := range inp {
		reports[i] = make([]int, 0, len(inp))
		reportValuesS := strings.Split(reportS, " ")
		for _, reportValueS := range reportValuesS {
			value, _ := strconv.Atoi(reportValueS)
			reports[i] = append(reports[i], value)
		}
	}

	return reports
}

func parsePart2(inp []string) [][]int {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(nums [][]int) int {
	sum := len(nums)

	for i := 0; i < len(nums); i++ {
		report := nums[i]
		if getOrder(report) == ORDER_NOT_IN_ORDER {
			sum--
		}
	}

	return sum
}

func Part2(nums [][]int) int {
	var sum int

	for i := 0; i < len(nums); i++ {
		report := nums[i]
		if getOrder(report) != ORDER_NOT_IN_ORDER {
			sum++
			continue
		}

		possible := false
		common.Log.Debug().Ints("ORIGINAL", report).Send()
		for j := 0; j < len(report); j++ {
			left := report[:j]
			right := report[j+1:]

			common.Log.Debug().Ints("left", left).Ints("right", right).Send()
			common.Log.Debug().Int("j", j).Int("report[j]", report[j]).Send()

			l := getOrder(left)
			if l == ORDER_NOT_IN_ORDER {
				continue
			}
			r := getOrder(right)
			if r == ORDER_NOT_IN_ORDER {
				continue
			}
			if (l == ORDER_INC && r == ORDER_DEC) || (l == ORDER_DEC && r == ORDER_INC) {
				continue
			}

			if len(left) == 0 || len(right) == 0 {
				possible = true
				break
			}

			if l == ORDER_IND {
				if r == ORDER_INC {
					if left[0] < right[0] && right[0]-left[0] <= 3 {
						common.Log.Debug().Ints("_left", left).Ints("_right", right).Int("a < ", left[0]).Int("b", right[0]).Send()
						possible = true
					}
				} else if r == ORDER_DEC {
					if left[0] > right[0] && left[0]-right[0] <= 3 {
						common.Log.Debug().Ints("_left", left).Ints("_right", right).Int("a > ", left[0]).Int("b", right[0]).Send()
						possible = true
					}
				}
			} else if l == ORDER_INC {
				if left[len(left)-1] < right[0] && right[0]-left[len(left)-1] <= 3 {
					common.Log.Debug().Ints("_left", left).Ints("_right", right).Int("a < ", left[len(left)-1]).Int("b", right[0]).Send()
					possible = true
				}
			} else if l == ORDER_DEC {
				if left[len(left)-1] > right[0] && left[len(left)-1]-right[0] <= 3 {
					common.Log.Debug().Ints("_left", left).Ints("_right", right).Int("a > ", left[len(left)-1]).Int("b", right[0]).Send()
					possible = true
				}
			}

			if possible {
				break
			}
		}
		if possible {
			sum++
		}
	}

	return sum
}

func getOrder(report []int) int {
	if len(report) < 2 {
		return ORDER_IND
	}

	if report[1] > report[0] {
		for j := 1; j < len(report); j++ {
			diff := report[j] - report[j-1]
			if diff <= 0 || diff >= 4 {
				return ORDER_NOT_IN_ORDER
			}
		}

		return ORDER_INC
	} else {
		for j := 1; j < len(report); j++ {
			diff := report[j-1] - report[j]
			if diff <= 0 || diff >= 4 {
				return ORDER_NOT_IN_ORDER
			}
		}
		return ORDER_DEC
	}

}
