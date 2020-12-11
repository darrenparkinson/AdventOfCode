package main

import (
	"crypto/md5"
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
	return getResult(input, "00000")
}

func part2(input []string) int {
	return getResult(input, "000000")
}

func getResult(input []string, prefix string) int {
	code := input[0]
	i := 0
	for {
		s := fmt.Sprintf("%s%d", code, i)
		sum := md5.Sum([]byte(s))
		result := fmt.Sprintf("%x", sum)
		if strings.HasPrefix(result, prefix) {
			return i
		}
		i++
	}

	return i
}
