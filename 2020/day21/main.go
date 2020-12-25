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

	for len(known) != uniqueAllegrens {
		for k, v := range found {
			if v == 1 {
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

	for _, v := range allegrens {
		ingredients = append(ingredients, knownReverse[v])
	}

	fmt.Println(strings.Join(ingredients, ","))

}

func getThisAllegren(allegren string, foods []food) {
	var current = allIngredients

	for _, food := range foods {
		hasThis := false

		for _, v := range food.allegrens {
			if v == allegren {
				hasThis = true
				break
			}
		}

		if hasThis {
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

// Can be made faster by going over smaller array first

func intersect(arr1, arr2 []string) []string {
	var intersection []string

	if len(arr1) == 0 || len(arr2) == 0 {
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

func parse(unparsed string) food {
	split := strings.Split(unparsed, "(")
	ingredients := strings.Split(strings.TrimSpace(split[0]), " ")

	ssss := strings.Replace(split[1], ")", "", -1)
	ssss = strings.Replace(ssss, ",", "", -1)
	ssss = strings.TrimSpace(ssss)
	sssss := strings.Split(ssss, " ")[1:]

	return food{ingredients: ingredients, allegrens: sssss}
}
