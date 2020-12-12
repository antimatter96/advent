package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	//inp := ""
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	inp += scanner.Text()
	// }

	// if scanner.Err() != nil {
	// 	// handle error.
	// }

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
	}
	day1_2(string(b))
}

func day1(inp string) {
	total := 0

	xx := extract(inp)

	for _, s := range xx {

		fields := len(s) / 2
		fmt.Println(fields, s)

		if fields == 8 {
			total++
		} else if fields == 7 && !hasCid(s) {
			total++
		}
	}

	fmt.Println(total)
}

func day1_2(inp string) {
	total := 0

	xx := extract(inp)

	fmt.Println(len(xx))

	for _, s := range xx {
		if isValidPassport(s) {
			total++
		}
	}

	fmt.Println(total)
}

/*
byr (Birth Year)
iyr (Issue Year)
eyr (Expiration Year)
hgt (Height)
hcl (Hair Color)
ecl (Eye Color)
pid (Passport ID)
cid (Country ID)
*/

func extract(inp string) [][]string {
	inp = strings.TrimSpace(inp)
	var m [][]string
	ss := strings.Split(inp, "\n\n")

	for _, s := range ss {
		sss := regexp.MustCompile(`[\s:]`).Split(s, -1)
		m = append(m, sss)
	}

	return m
}

func hasCid(ss []string) bool {
	for i := 0; i < len(ss); i += 2 {
		if ss[i] == "cid" {
			return true
		}
	}

	return false
}

func isValidPassport(ss []string) bool {
	//fmt.Println("checking", ss)
	if len(ss)/2 < 7 {
		//fmt.Println("less than 7 fields")
		return false
	}
	// for i := 0; i < len(ss); i += 2 {
	// 	fmt.Print(ss[i], " ")
	// }
	// fmt.Println()

	for i := 0; i < len(ss); i += 2 {
		if ss[i] == "cid" {
			if len(ss) == 14 {
				return false
			}
			continue
		}

		//fmt.Printf("|%s|%s|\n", ss[i])
		x := isValidField(ss[i], ss[i+1])

		if !x {
			//fmt.Println(ss[i], ss[i+1], x)
			return false
		}
	}

	return true
}

/*
byr (Birth Year) - four digits; at least 1920 and at most 2002.
iyr (Issue Year) - four digits; at least 2010 and at most 2020.
eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
hgt (Height) - a number followed by either cm or in:
	If cm, the number must be at least 150 and at most 193.
	If in, the number must be at least 59 and at most 76.
hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
ecl (Eye Color) - exactly one of:
	amb
	blu
	brn
	gry
	grn
	hzl
	oth.
pid (Passport ID) - a nine-digit number, including leading zeroes.
--cid (Country ID) - ignored, missing or not.
*/

var lengths = map[string]int{
	"byr": 4,
	"iyr": 4,
	"eyr": 4,
	"hcl": 7,
	"ecl": 3,
	"pid": 9,
}

var accpetedEyeColors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func isValidField(field string, value string) bool {
	//fmt.Println("checking", field, value)
	if l, ok := lengths[field]; ok {
		if l != len(value) {
			//fmt.Println("length of", field, "should be", l, "is", len(value))
			return false
		}
	}

	switch field {
	case "byr", "iyr", "eyr":
		val, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		switch field {
		case "byr":
			if val >= 1920 && val <= 2002 {
				return true
			}
		case "iyr":
			if val >= 2010 && val <= 2020 {
				return true
			}
		case "eyr":
			if val >= 2020 && val <= 2030 {
				return true
			}
		}
	case "ecl":
		_, ok := accpetedEyeColors[value]
		return ok
	case "hcl":
		if value[0] != '#' {
			return false
		}

		for _, r := range value[1:] {
			if unicode.IsDigit(r) || (r >= 'a' && r <= 'f') {

			} else {
				return false
			}
		}

		return true
	case "pid":
		if _, err := strconv.Atoi(value); err != nil {
			fmt.Println("Can't convert", err)
			return false
		}
		return true
	case "hgt":
		n := len(value)
		if n > 5 || n < 4 {
			return false
		}

		if n == 5 && value[3:] == "cm" {
			val, err := strconv.Atoi(value[:3])
			if err != nil {
				fmt.Println("Can't convert", err)
				return false
			}
			if val >= 150 && val <= 193 {
				return true
			}
			return false
		} else if n == 4 && value[2:] == "in" {
			val, err := strconv.Atoi(value[:2])
			if err != nil {
				fmt.Println("Can't convert", err)
				return false
			}
			if val >= 59 && val <= 76 {
				return true
			}
			return false
		}
		return false
	default:
		panic(fmt.Errorf("%s %s", field, value))
	}

	return false
}
