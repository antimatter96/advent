package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antimatter96/advent/2023/common"
)

const (
	FINAL_STATE_ACCEPTED = "A"
	FINAL_STATE_REJECTED = "R"
)

type input struct {
	parts    []part
	ruleSets []ruleSet
}

type part map[string]int

func (p *part) Rating() int {
	sum := 0

	for _, v := range *p {
		sum += v
	}

	return sum
}

type ruleSet struct {
	name string

	rules []rule
}

func (rs *ruleSet) FinalState(part part) string {
	for _, rule := range rs.rules {

		if rule.empty {
			return rule.to
		}

		if rule.greaterThan {
			if part[rule.what] > rule.against {
				return rule.to
			}
		} else {
			if part[rule.what] < rule.against {
				return rule.to
			}
		}
	}

	return ""
}

type rule struct {
	what        string
	against     int
	greaterThan bool

	empty bool

	to string
}

func (r *rule) String() string {
	if r.empty {
		return fmt.Sprintf(" => %s", r.to)
	}
	if r.greaterThan {
		return fmt.Sprintf("%s > %d => %s", r.what, r.against, r.to)
	}

	return fmt.Sprintf("%s < %d => %s", r.what, r.against, r.to)
}

func main() {
	rawInput := common.TakeInputAsString()

	p1, p2 := Run(rawInput)
	common.Log.Info().Int("Answer 1", p1).Send()
	common.Log.Info().Int("Answer 2", p2).Send()
}

func parsePart1(inp string) input {
	inp = strings.TrimSpace(inp)

	groups := strings.Split(inp, "\n\n")

	ruleSetsRaw := strings.Split(groups[0], "\n")

	ruleSets := make([]ruleSet, 0, len(ruleSetsRaw))
	for i := 0; i < len(ruleSetsRaw); i++ {

		ruleSetsRaw[i] = strings.ReplaceAll(ruleSetsRaw[i], "{", " ")
		ruleSetsRaw[i] = strings.ReplaceAll(ruleSetsRaw[i], "}", " ")
		ruleSetsRaw[i] = strings.ReplaceAll(ruleSetsRaw[i], ",", " ")

		ruleSetRaw := strings.Split(ruleSetsRaw[i], " ")

		ruleSet := ruleSet{rules: []rule{}}

		common.Log.Debug().Strs("rules", ruleSetRaw).Send()

		ruleSet.name = ruleSetRaw[0]

		for _, ruleRaw := range ruleSetRaw[1:] {
			if ruleRaw == "" || ruleRaw == " " {
				continue
			}
			common.Log.Debug().Str("ruleStr", ruleRaw).Send()

			rule := rule{}

			if !strings.ContainsRune(ruleRaw, ':') {
				rule.empty = true
				rule.to = ruleRaw
				ruleSet.rules = append(ruleSet.rules, rule)

				common.Log.Debug().Str("rule", rule.String()).Send()

				continue
			}

			ruleRawSplit := strings.Split(ruleRaw, ":")
			rule.to = ruleRawSplit[1]

			rule.what = string(ruleRawSplit[0][0])

			if ruleRawSplit[0][1] == '>' {
				rule.greaterThan = true
			} else {
				rule.greaterThan = false
			}

			against, _ := strconv.Atoi(ruleRawSplit[0][2:])
			rule.against = against

			ruleSet.rules = append(ruleSet.rules, rule)

			common.Log.Debug().Str("rule", rule.String()).Send()
		}

		ruleSets = append(ruleSets, ruleSet)
	}

	partsRaw := strings.Split(groups[1], "\n")

	parts := make([]part, 0, len(partsRaw))

	var x, m, a, s int
	for _, partRaw := range partsRaw {
		part := make(part)
		parts = append(parts, part)

		fmt.Sscanf(partRaw, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)

		part["x"] = x
		part["m"] = m
		part["a"] = a
		part["s"] = s

		common.Log.Debug().Any("", part).Send()
	}

	return input{parts: parts, ruleSets: ruleSets}
}

func parsePart2(inp string) input {
	return parsePart1(inp)
}

func Run(inp string) (int, int) {
	parsedPart1 := parsePart1(inp)
	// parsedPart2 := parsePart2(inp)

	return Part1(parsedPart1), Part2(parsedPart1)
}

func Part1(inp input) int {
	sum := 0

	rulesMap := make(map[string]*ruleSet)
	for i := 0; i < len(inp.ruleSets); i++ {
		rulesMap[inp.ruleSets[i].name] = &inp.ruleSets[i]

		common.Log.Debug().Str("rule", inp.ruleSets[i].name).Send()
	}

	common.Log.Debug().Any("", rulesMap).Send()

	for i, part := range inp.parts {
		state := "in"

		for state != FINAL_STATE_ACCEPTED && state != FINAL_STATE_REJECTED {
			nextState := rulesMap[state].FinalState(part)

			common.Log.Debug().Str("current", state).Str("next", nextState).Send()

			state = nextState
		}

		if state == FINAL_STATE_ACCEPTED {
			common.Log.Debug().Int("[i]", i).Int("rating", part.Rating()).Send()

			sum += part.Rating()
		}
	}

	return sum
}

type partRange struct {
	min int
	max int
}

type testRange map[string][]partRange

func deepCopyTestRange(given testRange) testRange {
	newTestRange := make(testRange, 0)

	for k, v := range given {
		newTestRange[k] = make([]partRange, 0, len(v))

		for i := 0; i < len(v); i++ {
			newTestRange[k][i] = partRange{min: v[i].min, max: v[i].max}
		}
	}

	return newTestRange
}

func newDefaultRange() testRange {
	newTestRange := make(testRange, 0)

	newTestRange["x"] = []partRange{{min: 0, max: 4000}}
	newTestRange["m"] = []partRange{{min: 0, max: 4000}}
	newTestRange["a"] = []partRange{{min: 0, max: 4000}}
	newTestRange["s"] = []partRange{{min: 0, max: 4000}}

	return newTestRange
}

func Part2(inp input) int {

	// global := newDefaultRange()

	return 0
}
