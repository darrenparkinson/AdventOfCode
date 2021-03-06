package main

import (
	"log"
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
	return len(input)
}

func part2(input []string) int {
	return len(input)
}
