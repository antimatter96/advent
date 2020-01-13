package main

import(
	"os"
	"fmt"
	"bufio"
	"strings"
)

var mp map[string]int

func main() {
	mp = make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	var commands [][]string
	for scanner.Scan() {
		inp := scanner.Text()
		strs := strings.Split(inp, ")")
		_, present := mp[strs[0]]
		if !present {
			mp[strs[0]] = total
			total++
		}

		_, present = mp[strs[1]]
                if !present {
                        mp[strs[1]] = total
                        total++
                }
		commands = append(commands, []string{ strs[0], strs[1] })
		//total += totalFuelForModule(i)
	}

	if scanner.Err() != nil {
		// handle error.
	}
	fmt.Println(total)
	grid := make([][]bool, total)
	for i := 0; i < total; i++ {
		grid[i] = make([]bool, total)
	}
	grid2 := make([]map[int]bool, total)
	for i := 0; i < total; i++ {
                grid2[i] = make(map[int]bool)
        }
	var posPl int
	var posSu int
	for _, planets := range commands {
		posPl = mp[planets[1]]
		posSu = mp[planets[0]]
		if grid[posPl][posSu] {
			fmt.Println("collide")
		} else {
			grid[posPl][posSu] = true
		}
		grid2[posPl][posSu] = true
		for i, _  := range grid2[posSu] {
			grid2[posPl][i] = true
		}
		for _, mapps  := range grid2 {
                        _, present := mapps[posPl]
			if present {
				for i, _ := range grid2[posPl] {
					mapps[i] = true
				}
			}
                }
	}
	tots := 0
	for _, mapss := range grid2 {
                tots += len(mapss)
        }
	fmt.Println(tots)
}
