package main

import (
	"log"
	"time"
)

func main() {

	start := time.Now()
	input := ReadInputAsRows("input.txt")

	result1 := part1(input, 25)
	log.Println("Result1:", result1)
	result2 := part2(input, 25)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string, preamble int) int {
	entries := RowsToIntSlice(input)
	subset := entries[:preamble]
	entry := preamble
	for _, e := range entries[preamble:] {
		valid := isValid(subset, e)
		if !valid {
			return e
		}
		entry++
		subset = entries[entry-preamble : entry]
	}
	return len(input)
}

func isValid(nums []int, num int) bool {
	for _, a := range nums {
		for _, b := range nums {
			if num == a+b {
				return true
			}
		}
	}
	return false
}

func part2(input []string, preamble int) int {
	entries := RowsToIntSlice(input)
	subset := entries[:preamble]
	entry := preamble
	interestingNumber := 0
	for _, e := range entries[preamble:] {
		valid := isValid(subset, e)
		if !valid {
			interestingNumber = e
			break
		}
		entry++
		subset = entries[entry-preamble : entry]
	}
	result := calcSum(entries, interestingNumber)
	return result
}

func calcSum(input []int, interestingNumber int) int {
	for i := range input {
		nums := input[i:]
		total := 0
		for j, e := range nums {
			total += e
			if total == interestingNumber {
				min, max := MinMaxIntSlice(nums[0 : j+1])
				return min + max
			}
		}
	}
	return len(input)
}
