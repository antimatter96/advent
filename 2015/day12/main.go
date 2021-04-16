package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err().Error())
	}

	day1(inp)
}

const loops = 50

func day1(inps []string) {

	for _, s := range inps {
		count2(s)
	}
}

func count(s string) {
	// fmt.Println("=========")

	var re = regexp.MustCompile(`[a-zA-Z\[\]\{\},":]`)

	s = re.ReplaceAllString(s, " ")

	//fmt.Println(s)

	var re2 = regexp.MustCompile(` +`)

	s = re2.ReplaceAllString(s, " ")

	//fmt.Println(s)

	arr := strings.Split(s, " ")
	// fmt.Println(arr)

	tots := 0
	for _, v := range arr {
		tt, _ := strconv.Atoi(v)

		tots += tt
	}

	fmt.Println(tots)
}

func count2(s string) {
	fmt.Println("=========")

	var tots float64 = 0
	dec := json.NewDecoder(strings.NewReader(s))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%T: %v", t, t)
		switch t.(type) {
		case float64:
			temp, _ := t.(float64)
			tots += temp
		}
		// fmt.Printf("\n")
	}

	fmt.Println(tots)

}
