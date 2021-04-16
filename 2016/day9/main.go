package main

import (
	"bufio"
	"fmt"
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
		panic(scanner.Err().Error())
	}

	day1(inp)
}

func day1(inp []string) {
	for _, s := range inp {
		_ = expand2(s)
	}
}

var rer strings.Builder
var next, repeat int
var c int

var i, j int
var sss []string
var n int

func expand(s string) string {
	rer.Reset()
	n = len(s)
	j = 0
	for i = 0; i < n; i++ {
		if s[i] == '(' {
			for j = i + 1; j < n; j++ {
				if s[j] == ')' {
					break
				}
			}
			sss = strings.Split(s[i+1:j], "x")
			next, _ = strconv.Atoi(sss[0])
			repeat, _ = strconv.Atoi(sss[1])

			for c = 0; c < repeat; c++ {
				rer.WriteString(s[j+1 : j+1+next])
			}
			i = j + next
		} else {
			rer.WriteByte(s[i])
		}
	}
	return rer.String()
}

var re = regexp.MustCompile(`\(\d+x\d+\)`)

type inss struct {
	low, high, next, repeat int
}

func expand2(s string) string {
	// fmt.Println("===============")
	// fmt.Println(s)

	insMarkers := re.FindAllStringSubmatchIndex(s, -1)
	if len(insMarkers) == 0 {
		return s
	}

	var aaa []inss

	for _, marker := range insMarkers {
		ins := s[marker[0]+1 : marker[1]-1]
		nextRepeat := strings.Split(ins, "x")
		next, _ := strconv.Atoi(nextRepeat[0])
		repeat, _ := strconv.Atoi(nextRepeat[1])

		aaa = append(aaa, inss{low: marker[0], high: marker[1], next: next, repeat: repeat})
	}

	//fmt.Println(aaa)

	//sss := ss[0]
	//pp := &stack{}
	//var whatShouldITake inss

	var aaw2 []int

	aa := 0

	marker := 0

	for i := 0; i < len(s); i++ {
		if marker < len(aaa) && i == aaa[marker].low {
			aaw2 = append(aaw2, aa)
			aa = 0

			aaw2 = append(aaw2, -(marker + 1))

			i = aaa[marker].high - 1

			marker++
		} else {
			aa++
		}
	}
	aaw2 = append(aaw2, aa)

	//fmt.Println(aaa)

	for neg(aaw2) {
		//fmt.Println(len(aaw2))
		aaw2 = removeZero(aaw2)
		//fmt.Println(len(aaw2))
		// fmt.Println("==>")
		// fmt.Println(aaw2)
		// fmt.Println("<==")
		running := 0
		howManyINeed := -1
		marker := 0
		var inssItook []int
		iTook := 0

		for i := 0; i < len(aaw2); i++ {
			//fmt.Println(aaw2)
			x := aaw2[i]
			if x == 0 {
				continue
			}
			if howManyINeed > -1 {
				if x < 0 {
					iTook++
					//fmt.Println("===",marker, iTook, marker+iTook, "===")
					//fmt.Println(aaa[marker+iTook], "length of", aaa[marker+iTook].high-aaa[marker+iTook].low)
	
					running += aaa[marker+iTook].high - aaa[marker+iTook].low

					inssItook = append(inssItook, x)
	
					aaw2[i] = 0
					//inssItook = append(inssItook, x)
					continue

				}
				//fmt.Println(">>>>", running, x, howManyINeed)
				
				if running+x >= howManyINeed {
					
					aaw2[i] = (running+x-howManyINeed)
					inssItook = append(inssItook, x-aaw2[i])
					//running = howManyINeed
					//fmt.Println(">>", running, aaw2[i], howManyINeed)
	
					//fmt.Println(inssItook)
					if neg(inssItook) {
						//fmt.Println("==>\n", aaw2, "\n<==")
						var wer []int
						wer = append(wer, aaw2[:i-1]...)
						for ve := 0; ve < aaa[marker].repeat; ve++ {
							wer = append(wer, inssItook...)
						}
						wer = append(wer, aaw2[i:]...)
						//fmt.Println("==>\n", wer, "\n<==")
						aaw2 = wer
	
						//break
						// inssItook = []int{}
						// howManyINeed = -1
						// i--
						// marker = 0
						// iTook = 0
						// running = 0
						//fmt.Println(aaw2[:i], inssItook, aaw2[i:])
					} else {
						aaw2[i] += (howManyINeed * aaa[marker].repeat)
						//break
					}
					break
					//fmt.Println(">", aaw2)
				} else {
					inssItook = append(inssItook, x)
					running += x
					aaw2[i] = 0
				}
				continue
			}
			if x < 0 {
				marker = -1 - x
				//fmt.Println(-1 - x)
				howManyINeed = aaa[marker].next
				aaw2[i] = 0
				running = 0
				iTook = 0
			}
		}
	}


	fmt.Println(sum(aaw2))
	return ""
}

func sum(arr []int) int {
	tots := 0
	for i := 0; i < len(arr); i++ {
		tots += arr[i]
	}
	return tots
}

func neg(arr []int) bool {
	for _, ele := range arr {
		if ele < 0 {
			return true
		}
	}
	return false
}

func removeZero(arr []int) []int {
	var filtered []int
	for _, ele := range arr {
		if ele != 0 {
			filtered = append(filtered, ele)
		}
	}
	arr = []int{}
	return filtered
}
