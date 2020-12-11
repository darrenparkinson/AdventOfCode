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

	homes := part1(input[0])
	log.Println("Homes Visited:", homes)

	homes2 := part2(input[0])
	log.Println("Homes Visited (Part 2):", homes2)

	log.Println("completed in ", time.Since(start))
}

type Vector struct {
	x int
	y int
}

func part1(input string) int {

	homes := make(map[Vector]int)
	currentHome := Vector{0, 0}
	homes[currentHome] = 1

	for _, rune := range input {
		switch string(rune) {
		case "^":
			currentHome = Vector{currentHome.x, currentHome.y + 1}
		case "<":
			currentHome = Vector{currentHome.x - 1, currentHome.y}
		case ">":
			currentHome = Vector{currentHome.x + 1, currentHome.y}
		case "v":
			currentHome = Vector{currentHome.x, currentHome.y - 1}
		}

		if _, found := homes[currentHome]; found {
			homes[currentHome]++
			continue
		}
		homes[currentHome] = 1
	}
	return len(homes)
}

func part2(input string) int {
	homes := make(map[Vector]int)
	currentHome := Vector{0, 0}
	currentHomeRobot := Vector{0, 0}
	homes[currentHome] = 1
	homes[currentHomeRobot] = 1

	for i, rune := range input {
		switch string(rune) {
		case "^":
			if i%2 == 0 {
				currentHome = Vector{currentHome.x, currentHome.y + 1}
			} else {
				currentHomeRobot = Vector{currentHomeRobot.x, currentHomeRobot.y + 1}
			}
		case "<":
			if i%2 == 0 {
				currentHome = Vector{currentHome.x - 1, currentHome.y}
			} else {
				currentHomeRobot = Vector{currentHomeRobot.x - 1, currentHomeRobot.y}
			}
		case ">":
			if i%2 == 0 {
				currentHome = Vector{currentHome.x + 1, currentHome.y}
			} else {
				currentHomeRobot = Vector{currentHomeRobot.x + 1, currentHomeRobot.y}
			}
		case "v":
			if i%2 == 0 {
				currentHome = Vector{currentHome.x, currentHome.y - 1}
			} else {
				currentHomeRobot = Vector{currentHomeRobot.x, currentHomeRobot.y - 1}
			}
		}

		if i%2 == 0 {

			if _, found := homes[currentHome]; found {
				homes[currentHome]++
				continue
			}
			homes[currentHome] = 1
		}
		if i%2 == 1 {

			if _, found := homes[currentHomeRobot]; found {
				homes[currentHomeRobot]++
				continue
			}
			homes[currentHomeRobot] = 1
		}
	}
	return len(homes)
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
