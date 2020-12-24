package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Not goroutine safe
type graph struct {
	mapping map[string]int
	reverse []string

	internal       [][]int
	reverseInteral [][]int

	counter int

	combinedWeight []int
}

func (g *graph) New(n int) {
	g.internal = make([][]int, n)
	g.reverseInteral = make([][]int, n)
	g.combinedWeight = make([]int, n)

	for i := 0; i < n; i++ {
		g.internal[i] = make([]int, n)
		g.reverseInteral[i] = make([]int, n)

		g.combinedWeight[i] = -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			g.internal[i][j] = -1
			g.reverseInteral[i][j] = -1
		}
	}

	g.reverse = make([]string, n)
}

func (g *graph) Add(u, v string, weight int) {
	if g.mapping == nil {
		g.mapping = make(map[string]int)
	}

	var indexU = -1
	var indexV = -1

	if k, present := g.mapping[u]; present {
		indexU = k
	} else {
		indexU = g.counter
		g.counter++
		g.mapping[u] = indexU
		g.reverse[indexU] = u

		//fmt.Println(u, indexU)
	}

	if k, present := g.mapping[v]; present {
		indexV = k
	} else {
		indexV = g.counter
		g.counter++
		g.mapping[v] = indexV
		g.reverse[indexV] = v

		//fmt.Println(v, indexV)
	}

	//fmt.Println(u, v, weight)

	g.internal[indexU][indexV] = weight
	g.reverseInteral[indexV][indexU] = weight
}

func (g *graph) totalWeight(index int) int {
	if g.combinedWeight[index] != -1 {
		//fmt.Println("finding total weight", index, "found")
		return g.combinedWeight[index]
	}

	//fmt.Println("finding total weight", index)

	changed := false
	sum := 0
	for i := 0; i < len(g.reverse); i++ {
		if i != index {
			// fmt.Println("finding total weight", index, i)
			if g.reverseInteral[index][i] != -1 {
				//fmt.Println("finding total weight", index, i, g.reverseInteral[index][i], "*", g.totalWeight(i))
				sum += (g.reverseInteral[index][i] * g.totalWeight(i))
				changed = true
			} else {
				//fmt.Println("NOT finding total weight", index, i, g.reverseInteral[index][i])
			}
		}
	}

	sum++

	if !changed {
		//fmt.Println("no change I guess for", index)
	}

	g.combinedWeight[index] = sum

	//fmt.Println(g.combinedWeight)

	return g.combinedWeight[index]
}

func (g *graph) Print() {
	fmt.Printf("|%s", "          ")
	for i := 0; i < len(g.reverse); i++ {
		fmt.Printf("|%s", g.reverse[i][0:8])
	}
	fmt.Printf("|\n")
	for i := 0; i < len(g.reverse); i++ {
		fmt.Printf("| %s ", g.reverse[i][0:8])
		for j := 0; j < len(g.reverse); j++ {
			fmt.Printf("|   %2d   ", g.reverseInteral[i][j])
		}
		fmt.Printf("|\n")
	}
}

func (g *graph) whatsMyIndex(u string) int {
	return g.mapping[u]
}

func (g *graph) doDFS(u string, mp map[int]bool) {
	//fmt.Println("Checking who can reach", u, g.internal[g.mapping[u]])
	mp[g.mapping[u]] = true
	for i, v := range g.internal[g.mapping[u]] {
		if v == -1 {
			continue
		}
		if mp[i] {
			//fmt.Println("Checking who can reach", u, "already seen", g.reverse[v])
			continue
		}
		//fmt.Println("Checking who can reach", u, "CAN", g.reverse[v])
		mp[i] = true
		for j, w := range g.internal[i] {
			if w == -1 {
				continue
			}
			g.doDFS(g.reverse[j], mp)
		}
	}
}

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

var re = regexp.MustCompile(`(\d+)? ([a-z]+ [a-z]+) bags?`)

type edge struct {
	u, v   string
	weight int
}

func day1(unparsed []string) {

	var edges []edge

	vertex := make(map[string]bool)

	for _, u := range unparsed {
		split := strings.Split(u, " bags contain ")

		ss := re.FindAllStringSubmatch(split[1], -1)

		if len(ss) > 0 {
			for _, s := range ss {
				weight, _ := strconv.Atoi(s[1])
				edges = append(edges, edge{s[2], split[0], weight})
				//fmt.Println(split[0], s[1], s[2])
				vertex[s[2]] = true
			}
		}

		vertex[split[0]] = true
	}

	p := &graph{}
	p.New(len(vertex))

	for _, edg := range edges {
		//fmt.Println("Adding", edg)
		p.Add(edg.u, edg.v, edg.weight)
	}

	result := make(map[int]bool)

	p.doDFS("shiny gold", result)

	fmt.Println(len(result) - 1)

	//p.Print()

	fmt.Println(p.totalWeight(p.whatsMyIndex("shiny gold")) - 1)

	//fmt.Println(edges)
}
