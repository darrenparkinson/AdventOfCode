package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {

	start := time.Now()
	input := readInput("input.txt")

	// part1(input)
	part2(input)
	// part2alternate(input)

	log.Println("completed in ", time.Since(start))
}

func part2original(input []string) {
	answer := 1
	result := calculateTrees(input, 1, 1)
	log.Println(result)
	answer = answer * result
	result = calculateTrees(input, 3, 1)
	log.Println(result)
	answer = answer * result
	result = calculateTrees(input, 5, 1)
	log.Println(result)
	answer = answer * result
	result = calculateTrees(input, 7, 1)
	log.Println(result)
	answer = answer * result
	result = calculateTrees(input, 1, 2)
	log.Println(result)
	answer = answer * result
	log.Println(answer)
}

func part2(input []string) {
	answer := 1
	runs := [][]int{[]int{1, 1}, []int{3, 1}, []int{5, 1}, []int{7, 1}, []int{1, 2}}
	for _, run := range runs {
		result := calculateTrees(input, run[0], run[1])
		answer = answer * result
	}
	log.Println(answer)
}

func part1(input []string) {
	result := calculateTrees(input, 3, 1)
	log.Println(result)
}

func calculateTrees(input []string, right, down int) int {
	treeCount := 0
	xpos := 0
	ypos := 0
	for index, row := range input {
		if index != ypos {
			continue
		}
		if xpos >= len(row) {
			xpos = xpos - len(row)
		}
		// log.Println("row:", index, "column:", xpos, "ypos:", ypos, "mod:", index%down, "char:", string(row[xpos]))
		if string(row[xpos]) == "#" {
			treeCount++
		}
		xpos += right
		ypos += down
	}
	return treeCount
}

// readInput will read the two wires routes and return as two strings
func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
