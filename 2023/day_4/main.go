package main

import (
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

type Card struct {
	have    []int
	winning []int
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(cardStrings []string) []Card {
	cards := make([]Card, 0)

	for _, cardString := range cardStrings {
		card := Card{winning: make([]int, 0), have: make([]int, 0)}

		numbers := strings.Split(cardString, ":")[1]

		winningNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[0]), " ")
		haveNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[1]), " ")

		for _, winningNumber := range winningNumbers {
			if winningNumber == "" || winningNumber == " " {
				continue
			}
			num, _ := strconv.Atoi(winningNumber)
			card.winning = append(card.winning, num)
		}

		for _, haveNumber := range haveNumbers {
			if haveNumber == "" || haveNumber == " " {
				continue
			}
			num, _ := strconv.Atoi(haveNumber)
			card.have = append(card.have, num)
		}

		cards = append(cards, card)
	}

	return cards
}

func parsePart2(inp []string) []Card {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(cards []Card) int {
	var sum int

	for cardI, card := range cards {
		common.Log.Debug().Int("Card", cardI+1).Ints("winning", card.winning).Send()
		common.Log.Debug().Int("Card", cardI+1).Ints("have", card.have).Send()

		winning := make(map[int]struct{})

		for i := 0; i < len(card.winning); i++ {
			winning[card.winning[i]] = struct{}{}
		}

		common.Log.Debug().Int("Card", cardI+1).Ints("diff", []int{len(winning), len(card.winning)}).Send()

		var totalHave int

		for i := 0; i < len(card.have); i++ {
			if _, ok := winning[card.have[i]]; ok {
				common.Log.Debug().Int("Card", cardI+1).Int("match", card.have[i]).Send()
				totalHave++
			}
		}

		common.Log.Debug().Int("Card", cardI+1).Int("totalHave", totalHave).Send()

		if totalHave > 0 {
			common.Log.Debug().Int("Card", cardI+1).Int("score", 1<<(totalHave-1)).Send()

			sum += 1 << (totalHave - 1)
		}
	}

	return sum
}

func Part2(cards []Card) int {

	var sum int

	cardCount := make(map[int]int)

	for cardI := range cards {
		cardCount[cardI] = 1
	}

	for cardI := 0; cardI < len(cards); cardI++ {
		card := cards[cardI]

		winning := make(map[int]struct{})

		for i := 0; i < len(card.winning); i++ {
			winning[card.winning[i]] = struct{}{}
		}

		var totalHave int64

		for i := 0; i < len(card.have); i++ {
			if _, ok := winning[card.have[i]]; ok {
				totalHave++
			}
		}

		common.Log.Debug().Int("Card", cardI+1).Int64("totalHave", totalHave).Send()
		if totalHave > 0 {
			for j := 0; j < int(totalHave); j++ {
				common.Log.Debug().Int("Card", cardI+1).Int("add", cardI+j+1+1).Int("times", cardCount[cardI]).Send()
				cardCount[cardI+j+1] += cardCount[cardI]
			}
		}
	}

	for _, v := range cardCount {
		sum += v
	}

	return sum
}
