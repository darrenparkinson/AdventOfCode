package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {

	start := time.Now()
	input := readInput("input.txt")

	log.Println("Part1:", identifyFloor(input[0]))
	log.Println("Part2:", identifyBasement(input[0]))

	log.Println("completed in ", time.Since(start))
}

func identifyFloor(instructions string) int {
	floor := 0
	for _, rune := range instructions {
		switch string(rune) {
		case "(":
			floor++
		case ")":
			floor--
		}
	}
	return floor
}

func identifyBasement(instructions string) int {
	floor := 0
	for index, rune := range instructions {
		switch string(rune) {
		case "(":
			floor++
		case ")":
			floor--
		}
		if floor == -1 {
			return index + 1
		}
	}
	return 0
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
