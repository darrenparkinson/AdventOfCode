package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	input := readInput("input.txt")

	presents := parsePresents(input)

	result1 := part1(presents)
	log.Println("Part1:", result1)

	result2 := part2(presents)
	log.Println("Part2:", result2)

	log.Println("completed in ", time.Since(start))
}

// Present represents the dimensions of a given present
type Present struct {
	L             int
	W             int
	H             int
	SurfaceArea   int
	PaperRequired int
	Ribbon        int
	Bow           int
	TotalRibbon   int
}

func part2(presents []Present) int {
	total := 0
	for _, present := range presents {
		total += present.TotalRibbon
	}
	return total
}

func part1(presents []Present) int {
	total := 0
	for _, present := range presents {
		total += present.PaperRequired
	}
	return total
}

func parsePresents(input []string) []Present {
	var presents []Present
	for _, row := range input {
		dims := strings.Split(row, "x")
		l, _ := strconv.Atoi(dims[0])
		w, _ := strconv.Atoi(dims[1])
		h, _ := strconv.Atoi(dims[2])
		a := 2 * l * w
		b := 2 * w * h
		c := 2 * h * l

		min := a
		if b < min {
			min = b
		}
		if c < min {
			min = c
		}

		lengths := []int{l, w, h}
		sort.Ints(lengths)

		ribbon := lengths[0]*2 + lengths[1]*2
		bow := l * w * h

		present := Present{l, w, h, a + b + c, a + b + c + min/2, ribbon, bow, ribbon + bow}

		presents = append(presents, present)
	}
	return presents
}

// readInput reads the input file and returns a slice of strings containing each line
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

// readRawInput returns a string of the entire input file
func readRawInput(path string) string {

	byteContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	content := string(byteContent)
	return content
}
