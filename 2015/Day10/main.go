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
	result2 := part2Ints(input)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string) int {

	result := input[0]
	for i := 0; i < 40; i++ {
		result = phonetic(result)
	}
	return len(result)
}

func part2(input []string) int {
	result := input[0]
	for i := 0; i < 50; i++ {
		result = phonetic(result)
	}
	return len(result)
}

func part2Ints(input []string) int {
	result := []int{1, 3, 2, 1, 1, 3, 1, 1, 1, 2}
	for i := 0; i < 50; i++ {
		result = phoneticIntBuilder(result)
	}
	return len(result)
}

func phoneticIntBuilder(s []int) (result []int) {
	for i := 0; i < len(s); {
		c := s[i]
		num := 1
		for i+num < len(s) && s[i+num] == c {
			num++
		}
		result = append(result, num, c)
		i += num
	}
	return result
}

func phoneticBuilder(input string) string {
	var sb strings.Builder
	toaddCount := 1
	currentChar := input[0]
	for _, s := range input[1:] {
		if string(s) == string(currentChar) {
			toaddCount++
			continue
		}
		sb.WriteString(fmt.Sprintf("%d%s", toaddCount, string(currentChar)))
		toaddCount = 1
		currentChar = byte(s)
	}
	sb.WriteString(fmt.Sprintf("%d%s", toaddCount, string(currentChar)))
	return sb.String()
}

func phonetic(input string) string {
	result := ""
	toaddCount := 1
	currentChar := input[0]
	for _, s := range input[1:] {
		if string(s) == string(currentChar) {
			toaddCount++
			continue
		}
		result += fmt.Sprintf("%d%s", toaddCount, string(currentChar))
		toaddCount = 1
		currentChar = byte(s)
	}
	result += fmt.Sprintf("%d%s", toaddCount, string(currentChar))
	return result
}
