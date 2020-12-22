package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
)

func main() {

	input := aocutils.ReadInputAsRows("input.txt")

	start := time.Now()
	result1 := part1(input)
	log.Println("Result1:", result1)
	log.Println("completed in ", time.Since(start))

	start = time.Now()
	result2 := part2(input)
	log.Println("Result2:", result2)
	log.Println("completed in ", time.Since(start))

}

func part1(input []string) int {

	_, ingredientCount, allergensWithIngredients := extractDetails(input)

	// allergensWithIngredients now contains all the possible ingredients for a given allergen
	// we still can't be certain, but we know these may well be allergens
	// so we'll make a set that contains each of the ingredients
	setOfPotentialIngredientsContainingAllergens := make(map[string]bool)
	for _, v := range allergensWithIngredients {
		for ing := range v {
			setOfPotentialIngredientsContainingAllergens[ing] = true
		}
	}
	// now we can work how those ingredients that don't contain those allergens
	total := 0
	for k, v := range ingredientCount {
		if _, ok := setOfPotentialIngredientsContainingAllergens[k]; !ok {
			total += v
		}
	}

	return total
}

func part2(input []string) string {
	_, _, allergensWithIngredients := extractDetails(input)
	// fmt.Println(foods)
	fmt.Println(allergensWithIngredients)
	cont := true
	for cont == true {
		cont = false
		for a, i := range allergensWithIngredients {
			if len(i) > 1 {
				cont = true
			}
			//TODO: add check to make sure we're not trying to remove them all again
			if len(i) == 1 {
				for a2, i2 := range allergensWithIngredients {
					if a == a2 {
						continue
					}
					for ing := range i2 {
						// fmt.Println(ing)
						if _, ok := i[ing]; ok {
							delete(i2, ing)
						}
					}
				}
			}
		}
	}
	fmt.Println(allergensWithIngredients)
	var sortkeys []string
	var vals []string
	for a := range allergensWithIngredients {
		sortkeys = append(sortkeys, a)
	}
	sort.Strings(sortkeys)

	for _, k := range sortkeys {
		a := allergensWithIngredients[k]
		for v := range a {
			// fmt.Println(k, v)
			vals = append(vals, v)
		}
	}
	return strings.Join(vals, ",")
}

type Food struct {
	ingredients []string
	allergens   []string
}

func parseInput(input []string) []Food {
	var foods []Food
	for _, row := range input {
		res := regexp.MustCompile("^(.+)(\\(contains )(.+)\\)$").FindStringSubmatch(row)
		food := Food{strings.Fields(res[1]), strings.Split(res[3], ", ")}
		foods = append(foods, food)
	}
	return foods
}

func extractDetails(input []string) ([]Food, map[string]int, map[string]map[string]bool) {
	foods := parseInput(input)
	ingredientCount := make(map[string]int)
	allergensWithIngredients := make(map[string]map[string]bool)
	for _, food := range foods {
		// add to the count which we might need later
		for _, ingredient := range food.ingredients {
			if _, ok := ingredientCount[ingredient]; !ok {
				ingredientCount[ingredient] = 0
			}
			ingredientCount[ingredient]++
		}
		// now work out the allergens
		for _, allergen := range food.allergens {
			if _, ok := allergensWithIngredients[allergen]; !ok {
				allergensWithIngredients[allergen] = make(map[string]bool)
				// add the first lot of ingredients we find since it didn't exist before
				for _, food := range food.ingredients {
					allergensWithIngredients[allergen][food] = true
				}
			}
			// now intersect the ingredients of this allergen
			n := make(map[string]bool)
			for _, food := range food.ingredients {
				if _, ok := allergensWithIngredients[allergen][food]; ok {
					// allergensWithIngredients[allergen][food] = true
					n[food] = true
				}
			}
			allergensWithIngredients[allergen] = n
		}
	}
	return foods, ingredientCount, allergensWithIngredients
}
