package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
)

func main() {

	input := aocutils.ReadInputAsRows("input.txt")

	start := time.Now()
	result1 := part1(input[0])
	log.Println("Result1:", result1)
	log.Println("completed in ", time.Since(start))

	start = time.Now()
	result2 := part2(input[0])
	log.Println("Result2:", result2)
	log.Println("completed in ", time.Since(start))
}

func part1(input string) int {
	turn := 1
	lastnum := 0
	lastseen := make(map[int]int)
	for _, s := range strings.Split(input, ",") {
		lastnum, _ := strconv.Atoi(s)
		lastseen[lastnum] = turn // e.g. last time seen
		turn++
	}

	for turn < 2020 {
		if val, ok := lastseen[lastnum]; ok {
			// we've seen the number before
			diff := turn - val
			lastseen[lastnum] = turn
			lastnum = diff
			turn++
			continue
		}
		// we haven't seen it before
		lastseen[lastnum] = turn
		lastnum = 0
		turn++
	}

	return lastnum
}

func part2(input string) int {

	turn := 1
	lastnum := 0
	lastseen := make(map[int]int)
	for _, s := range strings.Split(input, ",") {
		lastnum, _ := strconv.Atoi(s)
		lastseen[lastnum] = turn // e.g. last time seen
		turn++
	}

	for turn < 30000000 {
		if val, ok := lastseen[lastnum]; ok {
			// we've seen the number before
			diff := turn - val
			lastseen[lastnum] = turn
			lastnum = diff
			turn++
			continue
		}
		// we haven't seen it before
		lastseen[lastnum] = turn
		lastnum = 0
		turn++
	}

	return lastnum

}
