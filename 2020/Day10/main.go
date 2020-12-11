package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func main() {

	start := time.Now()
	input := ReadInputAsInts("input.txt")

	result1 := part1(input)
	log.Println("Result1:", result1)
	result2 := part2(input)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []int) int {
	min, max := MinMaxIntSlice(input)
	rated := max + 3
	outlet := 0
	inputDict := make(map[int]int)
	for _, in := range input {
		inputDict[in] = in
	}
	diff1v := 0
	diff3v := 0
	fmt.Println(outlet, min, max, rated, inputDict, diff1v, diff3v)
	current := 0 // haha current, did you see what I did there
	for j := 0; current != max; j++ {

		if val, ok := inputDict[current+1]; ok {
			diff1v++
			current = val
			continue
		}
		if val, ok := inputDict[current+3]; ok {
			diff3v++
			current = val
			continue
		}
	}
	diff3v++ // for rated
	return diff1v * diff3v
}

func part2(input []int) int {
	sort.Ints(input)
	acc := map[int]int{0: 1}
	for _, i := range input {
		acc[i] = acc[i-1] + acc[i-2] + acc[i-3]
	}
	lastItem := input[len(input)-1]
	return acc[lastItem]
}
