package main

import (
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

type handBid struct {
	hand     []rune
	bid      int
	count    map[rune]int
	maxCount int

	cardScore int
}

func main() {
	rawInput := common.TakeInputAsStringArray()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp []string) []handBid {
	handBids := make([]handBid, 0)

	for _, handBidStr := range inp {
		split := strings.Split(strings.TrimSpace(handBidStr), " ")

		bid, _ := strconv.Atoi(split[1])
		handBids = append(handBids, handBid{
			bid:  bid,
			hand: []rune(strings.TrimSpace(split[0])),
		})
	}

	return handBids
}

func parsePart2(inp []string) []handBid {
	return parsePart1(inp)
}

func Run(inp []string) (int, int) {
	parsedPart1 := parsePart1(inp)
	parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart2)
}

func Part1(handBids []handBid) int {
	var cardStrengthOne = map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	winnings := 0

	for i := 0; i < len(handBids); i++ {

		handBids[i].count = make(map[rune]int)

		for j := 0; j < len(handBids[i].hand); j++ {
			handBids[i].count[handBids[i].hand[j]]++
		}

		count := make([]int, 0)

		for _, v := range handBids[i].count {
			count = append(count, -v)
		}

		slices.Sort[[]int, int](count)

		handBids[i].cardScore = determineScore(count)

		common.Log.Debug().Str("Hand", string(handBids[i].hand)).Int("Max", handBids[i].maxCount).Ints("Counts", count).Send()
	}

	sort.Slice(handBids, func(i, j int) bool {

		if handBids[i].cardScore != handBids[j].cardScore {
			return handBids[i].cardScore < handBids[j].cardScore
		}

		return compareCards(handBids[i].hand, handBids[j].hand, cardStrengthOne)
	})

	for i, habdhandBid := range handBids {
		// common.Log.Debug().Str("hand", string(habdhandBid.hand)).Int("rank", i+1).Int("bid", habdhandBid.bid).Send()

		winnings += (i + 1) * (habdhandBid.bid)
	}

	return winnings
}

func compareCards(hand1, hand2 []rune, Strength map[rune]int) bool {

	for i := 0; i < len(hand1); i++ {
		if hand1[i] == hand2[i] {
			continue
		}

		return Strength[hand1[i]] < Strength[hand2[i]]
	}

	return false
}

func Part2(handBids []handBid) int {
	var cardStrengthOne = map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': -1,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	winnings := 0

	for i := 0; i < len(handBids); i++ {

		handBids[i].count = make(map[rune]int)

		countJ := 0
		for j := 0; j < len(handBids[i].hand); j++ {
			if handBids[i].hand[j] == 'J' {
				countJ++
				continue
			}
			handBids[i].count[handBids[i].hand[j]]++
		}

		count := make([]int, 0)

		for _, v := range handBids[i].count {
			count = append(count, -v)
		}

		slices.Sort[[]int, int](count)

		if len(count) == 0 {
			count = append(count, -countJ)
		} else {
			count[0] -= countJ
		}

		handBids[i].cardScore = determineScore(count)

		common.Log.Debug().Str("Hand", string(handBids[i].hand)).Int("Score", handBids[i].cardScore).Ints("Counts", count).Send()
	}

	sort.Slice(handBids, func(i, j int) bool {

		if handBids[i].cardScore != handBids[j].cardScore {
			return handBids[i].cardScore < handBids[j].cardScore
		}

		return compareCards(handBids[i].hand, handBids[j].hand, cardStrengthOne)
	})

	for i, habdhandBid := range handBids {
		// common.Log.Debug().Str("hand", string(habdhandBid.hand)).Int("rank", i+1).Int("bid", habdhandBid.bid).Send()

		winnings += (i + 1) * (habdhandBid.bid)
	}

	return winnings
}

func determineScore(count []int) int {
	if count[0] == -5 {
		return 6
	}
	if count[0] == -4 {
		return 5
	}
	if count[0] == -3 && count[1] == -2 {
		return 4
	}
	if count[0] == -3 && count[1] == -1 {
		return 3
	}
	if count[0] == -2 && count[1] == -2 {
		return 2
	}
	if count[0] == -2 && count[1] == -1 {
		return 1
	}

	return 0
}
