package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
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

	day1(inp)
}

func day1(inp []string) {
	total := 0

	for _, s := range inp {
		frequencyString, id, checksum := extract(s)
		//fmt.Println(frequencyString, ">> ", id, "<<", checksum)
		if isItTheChecksum(frequencyString, checksum) {
			total += id
		}

		fmt.Println(id, decryptString(frequencyString, id))
	}
	fmt.Println(total)
}

func extract(s string) (string, int, string) {
	//fmt.Println(s)
	var id int

	//ss := regexp.MustCompile(`([a-z\-]+)-`)
	ss := regexp.MustCompile(`([a-z\-]+)-([0-9]+)\[([a-z]+)\]`)
	//ss := regexp.MustCompile(`([a-z\-]+)-([\d]+)\[([a-z]+)\]`)

	sss := ss.FindAllStringSubmatch(s, -1)
	//fmt.Sscanf(s, "%s-%d[%s]", frequencyString, id, checksum)
	id, _ = strconv.Atoi(sss[0][2])
	return sss[0][1], id, sss[0][3]
}

func isItTheChecksum(freq, checksum string) bool {
	// Generate the freq map
	// Check if checksum's characters are in this map
	// If false :)
	// If true :( Calculate the actual checksum

	freqMp := make(map[rune]int)
	for _, r := range freq {
		if r == '-' {
			continue
		}
		freqMp[r]++
	}

	for _, r := range checksum {
		if _, present := freqMp[r]; !present {
			return false
		}
	}

	// :(

	return string(rankMapStringInt(freqMp)[0:5]) == checksum
}

func rankMapStringInt(values map[rune]int) []rune {
	type kv struct {
		Key   rune
		Value int
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value != ss[j].Value {
			return ss[i].Value > ss[j].Value
		}
		return ss[i].Key < ss[j].Key
	})
	ranked := make([]rune, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}

func decryptString(enc string, inc int) string {
	inc = (inc % 26)

	dec := make([]rune, len(enc))

	for i, r := range enc {
		dec[i] = nextRune(r, inc)
	}

	return string(dec)
}

var lowerLimit = int('a') - 'a'
var upperLimit = int('z') - 'a'

func nextRune(r rune, i int) rune {
	if r == '-' {
		return ' '
	}

	next := (int(r) - 'a') + i
	//fmt.Println(next, r, rune(next+int('a')), string(r), string(next+int('a')))
	if next > upperLimit {
		next -= 26
	}

	//fmt.Println(next, r, rune(next+int('a')), string(r), string(next+int('a')))

	return rune(next + int('a'))
}
