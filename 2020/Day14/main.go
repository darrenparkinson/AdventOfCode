package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
)

func main() {

	start := time.Now()
	input := aocutils.ReadInputAsRows("input.txt")

	result1 := part1(input)
	log.Println("Result1:", result1)
	result2 := part2(input)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string) int {
	mask := ""
	mem := map[int]int{}
	for _, row := range input {
		if strings.HasPrefix(row, "mask = ") {
			mask = strings.TrimPrefix(row, "mask = ")
			continue
		}
		var loc, val int
		fmt.Sscanf(row, "mem[%d] = %d", &loc, &val)
		result := applyMaskToValue(mask, val)
		mem[loc] = result
	}
	sum := 0
	for _, v := range mem {
		sum += v
	}
	return sum
}

func part2(input []string) int {
	return len(input)
}

func applyMaskToValue(mask string, value int) int {
	andMask := strings.ReplaceAll(mask, "X", "1")
	andVal, _ := strconv.ParseUint(andMask, 2, 0)
	anded := value & int(andVal)

	orMask := strings.ReplaceAll(mask, "X", "0")
	orVal, _ := strconv.ParseUint(orMask, 2, 0)
	ord := anded | int(orVal)

	return ord
}

// func calcMask(mask string) int64 {
// 	m := reverse(mask)
// 	var result int64
// 	r := regexp.MustCompile(`1`).FindAllStringIndex(m, -1)
// 	for _, val := range r {
// 		result += int64(math.Pow(2, float64(val[1])))
// 	}
// 	return result
// }

// func reverse(s string) string {
// 	chars := []rune(s)
// 	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
// 		chars[i], chars[j] = chars[j], chars[i]
// 	}
// 	return string(chars)
// }
