package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
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

	day2(inp)
}

func day1(inp []string) {
	start, _ := strconv.Atoi(inp[0])

	fmt.Println(start)

	// var possible []int

	inps := strings.Split(inp[1], ",")

	var temp int
	min, minID := start*10, -1

	var meh int
	for _, inp := range inps {
		if inp != "x" {
			temp, _ = strconv.Atoi(inp)
			// fmt.Println(temp)

			meh = findNearest(start, temp)

			if meh < min {
				min = meh
				minID = temp
			}
			// possible = append(possible, temp)
		}
	}

	fmt.Println((min - start) * minID)
}

func findNearest(target, interval int) int {
	return target + interval - (target % interval)
}

func day2(inp []string) {
	for i := 1; i < len(inp); i++ {
		solve3(strings.Split(inp[i], ","))
	}
}

func solve(inps []string) {
	var temp int

	mp := make(map[int]int)

	// var meh int
	start := -1
	for i, inp := range inps {
		if inp != "x" {
			temp, _ = strconv.Atoi(inp)
			mp[temp] = i
		}
		if start == -1 {
			start = temp
		}
	}

	fmt.Println(start, mp)

	max, offSetOfMax := -1, -1

	for k, v := range mp {
		if k > max {
			max = k
			offSetOfMax = v
		}
	}

	for k, v := range mp {
		mp[k] = v - offSetOfMax
	}

	fmt.Println(start, mp)

	for i := 0; i < 1<<62; i += max {
		if what(i, mp) {
			fmt.Println(i - offSetOfMax)
			break
		}
	}
}

func what(i int, mp map[int]int) bool {
	for k, v := range mp {
		if (i+v)%k != 0 {
			return false
		}
	}

	return true
}

func solve2(inps []string) {
	var temp int

	mp := make(map[int]int)

	// var meh int
	start := -1
	for i, inp := range inps {
		if inp != "x" {
			temp, _ = strconv.Atoi(inp)
			mp[temp] = i
		}
		if start == -1 {
			start = temp
		}
	}

	fmt.Println(start, mp)

	max, offSetOfMax := -1, -1

	for k, v := range mp {
		if k > max {
			max = k
			offSetOfMax = v
		}
	}

	for k, v := range mp {
		mp[k] = v - offSetOfMax
	}

	intervals, offsets := make([]int, len(mp)-1), make([]int, len(mp)-1)
	temp = 0

	nh := make([]*big.Int, len(mp))
	ah := make([]*big.Int, len(mp))

	for k, v := range mp {
		
		nh[temp] = big.NewInt(int64(k))
		ah[temp] = big.NewInt(int64(v))
		// mp[k] = v - offSetOfMax

		temp++
	}

	// n := len(mp) - 1

	fmt.Println(start, mp)
	fmt.Println(intervals)
	fmt.Println(offsets)
	// fmt.Println(1 << 62)

	startFromBack := limit
	startFromBack -= (limit % max)

	// for i := startFromBack; i > -1; i -= max {
	// 	if wha2t(i, &intervals, &offsets, &n) {
	// 		fmt.Println(i - offSetOfMax)
	// 		// break
	// 	}
	// 	if i%j == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	fmt.Println(crt(ah, nh))

	os.Exit(0)
}


func solve3(inps []string) {
	var temp int


	var nh, ah []*big.Int


	for i, inp := range inps {
		if inp != "x" {
			temp, _ = strconv.Atoi(inp)

			nh = append(nh, big.NewInt(int64(temp)))
			ah = append(ah, big.NewInt(int64(temp-i)))
		}
	}

	fmt.Println(nh, ah)

	fmt.Println(crt(ah, nh))
}


const limit int = 1 << 62
const j int = 100000000000000000

var k int

func wha2t(target int, intervals, offsets *[]int, n *int) bool {
	for k = 0; k < *n; k++ {
		if (target+(*offsets)[k])%(*intervals)[k] != 0 {
			return false
		}
	}

	return true
}

// 3162341
// 1068781


var one = big.NewInt(1)
 
func crt(a, n []*big.Int) (*big.Int, error) {
    p := new(big.Int).Set(n[0])
    for _, n1 := range n[1:] {
        p.Mul(p, n1)
    }
    var x, q, s, z big.Int
    for i, n1 := range n {
        q.Div(p, n1)
        z.GCD(nil, &s, n1, &q)
        if z.Cmp(one) != 0 {
            return nil, fmt.Errorf("%d not coprime", n1)
        }
        x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
    }
    return x.Mod(&x, p), nil
}
