package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

var known = make(map[string]string)
var knownReverse = make(map[string]string)
var found = make(map[string]int)

var _temp = make(map[string]bool)
var allIngredients []string

type food struct {
	ingredients []string
	allegrens   []string
}

var uniqueAllegrens int

func day1(unparsed []string) {
	var foods []food
	for _, v := range unparsed {
		foods = append(foods, parse(v))
	}

	// fmt.Println(len(foods))

	// for i := 0; i < len(foods); i++ {
	// 	fmt.Println(len(foods[i].allegrens))
	// }

	uniqueAllegrens = 0

	for i := 0; i < len(foods); i++ {
		for _, v := range foods[i].allegrens {
			if found[v] == 0 {
				found[v]++
			}
		}

		for _, v := range foods[i].ingredients {
			if !_temp[v] {
				_temp[v] = true
				allIngredients = append(allIngredients, v)
			}
		}
	}

	uniqueAllegrens = len(found)
	// fmt.Println(len(allIngredients))

	for len(known) != uniqueAllegrens {
		for k, v := range found {
			// fmt.Println(k, v)
			if v == 1 {
				// fmt.Println(k, v)
				getThisAllegren(k, foods)
			}
		}

		reduce(foods)
	}

	reduce(foods)

	tots := 0
	for _, v := range foods {
		tots += len(v.ingredients)
	}

	fmt.Println(tots)

	var allegrens []string

	for k := range found {
		allegrens = append(allegrens, k)
	}

	sort.Strings(allegrens)
	var ingredients []string

	fmt.Println(allegrens)

	for _, v := range allegrens {
		ingredients = append(ingredients, knownReverse[v])
	}

	fmt.Println(strings.Join(ingredients, ","))

}

func getThisAllegren(allegren string, foods []food) {
	var current = allIngredients

	for _, food := range foods {
		do := false

		for _, v := range food.allegrens {
			if v == allegren {
				do = true
				break
			}
		}

		if do {
			current = intersect(food.ingredients, current)

			if len(current) == 1 {
				known[current[0]] = allegren
				knownReverse[allegren] = current[0]

				found[allegren] = 2

				fmt.Println(allegren, current[0], found[allegren])

				break
			}
		}
	}
}

// var skipList = make(map[string]bool)

// func shouldISkip(arr []int) bool {
// 	for i := 0; i < len(arr); i++ {
// 		for j := i + 1; j < len(arr); j++ {
// 			if skipList[fmt.Sprintf("%d,%d", i, j)] {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func updateSkipList(i, j int) {
// 	skipList[fmt.Sprintf("%d,%d", i, j)] = true
// }

// func process(foods []food) {
// 	skipList = make(map[string]bool)

// 	subProcess1(foods)

// 	subProcess2(foods)

// 	reduce(foods)

// 	subProcess3(foods)
// }

// func subProcess1(foods []food) {
// 	for i := 0; i < len(foods); i++ {
// 		for j := i + 1; j < len(foods); j++ {
// 			res := determineMany([]food{foods[i], foods[j]})
// 			if res == -1 {
// 				updateSkipList(i, j)
// 			}
// 		}
// 	}
// }

// func subProcess3(foods []food) {
// 	for i := 0; i < len(foods); i++ {
// 		determine1(foods[i])
// 	}
// }

// func subProcess2(foods []food) {
// 	fmt.Println(len(skipList), len(foods)*(len(foods)-1))
// 	limit := int64(1 << len(foods))

// 	s := fmt.Sprintf("%%0%db\n", len(foods))

// 	r := rand.New(rand.NewSource(time.Now().UnixNano()))

// 	mp := make(map[int64]bool)

// 	for i := int64(1); i < limit; i++ {
// 		w := r.Int63n(limit)
// 		//fmt.Println(w)
// 		if mp[w] {
// 			i--
// 			continue
// 			// i--
// 		}

// 		mp[w] = true

// 		if bits.OnesCount(uint(w)) < 3 {
// 			continue
// 		}

// 		var temp []food
// 		var tempInt []int

// 		jj := fmt.Sprintf(s, i)

// 		for k, v := range jj {
// 			if v == '1' {
// 				temp = append(temp, foods[k])
// 				tempInt = append(tempInt, k)
// 			}
// 		}

// 		if shouldISkip(tempInt) {
// 			// fmt.Println("SKIP")
// 			continue
// 		}

// 		if w%2 == 0 {
// 			fmt.Println(w, i, limit/i)
// 		}

// 		determineMany(temp)

// 		// fmt.Println(taken)

// 		if len(known) > 1 {
// 			fmt.Println("HOHOHO", len(known))
// 			// break
// 		}

// 		if len(known) == uniqueAllegrens {
// 			break
// 		}
// 	}
// }

func reduce(foods []food) {
	for i := 0; i < len(foods); i++ {
		var ingredients []string
		var allegrens []string

		for _, v := range foods[i].allegrens {
			allegrens = append(allegrens, v)
		}

		for _, v := range foods[i].ingredients {
			if known[v] != "" {
				allegrens = remove(allegrens, known[v])
			} else {
				ingredients = append(ingredients, v)
			}
		}

		foods[i].ingredients = ingredients
		foods[i].allegrens = allegrens
	}
}

func remove(arr []string, this string) []string {
	var temp []string
	for _, v := range arr {
		if v != this {
			temp = append(temp, v)
		}
	}

	return temp
}

// func determine1(food1 food) {
// 	if len(food1.ingredients) == len(food1.allegrens) && len(food1.ingredients) == 1 {
// 		known[food1.ingredients[0]] = food1.allegrens[0]
// 		knownReverse[food1.allegrens[0]] = food1.ingredients[0]
// 	}
// }

// func filterKnowns(arr1, arr2 []string) ([]string, []string) {
// 	var filtered1, filtered2 []string

// 	for _, v := range arr1 {
// 		if _, present := known[v]; !present {
// 			filtered1 = append(filtered1, v)
// 		}
// 	}
// 	for _, v := range arr2 {
// 		if _, present := knownReverse[v]; !present {
// 			filtered2 = append(filtered2, v)
// 		}
// 	}

// 	return filtered1, filtered2
// }

func intersect(arr1, arr2 []string) []string {
	var intersection []string

	if len(arr1) == 0 || len(arr2) == 0 {
		// fmt.Println("breaking")
		return intersection
	}

	mp := make(map[string]bool)

	for _, v := range arr1 {
		mp[v] = true
	}
	for _, v := range arr2 {
		if mp[v] {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

// func intersectMany(arr [][]string) []string {
// 	var intersection = arr[0]

// 	for i := 1; i < len(arr); i++ {
// 		intersection = intersect(arr[i], intersection)
// 		if len(intersection) == 0 {
// 			// fmt.Println("breaking")
// 			break
// 		}
// 	}

// 	return intersection
// }

// func determineMany(foods []food) int {
// 	if len(foods) == 1 {
// 		return 0
// 	}

// 	ingredients := make([][]string, len(foods))
// 	allegrens := make([][]string, len(foods))

// 	for i := 0; i < len(foods); i++ {
// 		ingredients[i] = foods[i].ingredients
// 		allegrens[i] = foods[i].allegrens

// 		if len(foods[i].allegrens) == 0 || len(foods[i].ingredients) == 0 {
// 			fmt.Println("empty array")
// 			return 0
// 		}
// 	}

// 	ingredientsFinal := intersectMany(ingredients)
// 	allegrensFinal := intersectMany(allegrens)

// 	if len(ingredientsFinal) == 0 || len(allegrensFinal) == 0 {
// 		// fmt.Println("empty intersection")
// 		return -1
// 	}

// 	//fmt.Println(len(ingredients), len(ingredientsFinal), len(allegrens), len(allegrensFinal))

// 	ingredientsFinal, allegrensFinal = filterKnowns(ingredientsFinal, allegrensFinal)

// 	// fmt.Println(len(foods), len(ingredientsFinal), len(allegrensFinal))

// 	if len(ingredientsFinal) == len(allegrensFinal) && len(ingredientsFinal) == 1 {
// 		known[ingredientsFinal[0]] = allegrensFinal[0]
// 		knownReverse[allegrensFinal[0]] = ingredientsFinal[0]
// 		return 1
// 	}

// 	return 0
// }

func parse(unparsed string) food {
	ss := strings.Split(unparsed, "(")
	sss := strings.Split(strings.TrimSpace(ss[0]), " ")

	ssss := strings.Replace(ss[1], ")", "", -1)
	ssss = strings.Replace(ssss, ",", "", -1)
	ssss = strings.TrimSpace(ssss)
	sssss := strings.Split(ssss, " ")[1:]

	return food{ingredients: sss, allegrens: sssss}
}
