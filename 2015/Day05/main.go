package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	input := ReadInputAsRows("input.txt")

	result1 := part1(input)
	log.Println("Result1:", result1)
	result2 := part2(input)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string) int {
	niceCount := 0
	for _, row := range input {
		if isValid(row) {
			niceCount++
		}
	}
	return niceCount
}

func part2(input []string) int {
	niceCount := 0
	for _, row := range input {
		if isValid2(row) {
			niceCount++
		}
	}
	return niceCount
}

func isValid2(input string) bool {
	input = strings.ToLower(input)
	if !containsTwoPairs(input) {
		return false
	}
	if !containsRepeatingLetterWithLetterBetween(input) {
		return false
	}
	return true
}

func containsRepeatingLetterWithLetterBetween(input string) bool {
	i := 0
	for {
		if i+2 >= len(input) {
			break
		}
		if input[i] == input[i+2] {
			return true
		}
		i++
	}
	return false
}

func containsTwoPairs(input string) bool {
	i := 0
	for {
		if i+2 > len(input) {
			break
		}
		matchPair := input[i : i+2]
		for j, s := range input[i+2:] {
			if i+3+j >= len(input) {
				break
			}
			if matchPair == fmt.Sprintf("%s%s", string(s), string(input[i+3+j])) {
				return true
			}
		}
		i++
	}
	return false
}

func isValid(input string) bool {
	input = strings.ToLower(input)
	invalidCombos := []string{"ab", "cd", "pq", "xy"}
	for _, combo := range invalidCombos {
		if strings.Contains(input, combo) {
			return false
		}
	}
	if !containsThreeVowels(input) {
		return false
	}

	// match any double character
	return containsDoubleChar(input)
}

func containsDoubleChar(input string) bool {
	previous := '\n'
	for _, r := range input {
		if r == previous {
			return true
		}
		previous = r
	}
	return false
}

func containsThreeVowels(input string) bool {
	input = strings.ToLower(input)
	vowelCount := 0
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, vowel := range vowels {
		vowelCount += strings.Count(input, vowel)
	}
	return vowelCount >= 3
}
