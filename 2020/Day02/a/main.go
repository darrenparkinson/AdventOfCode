package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	input := readInput("input.txt")

	result := calculateValidPasswords(input)
	log.Println(result)

	log.Println("completed in ", time.Since(start))
}

func calculateValidPasswords(input []string) int {
	counter := 0
	for _, line := range input {
		if calculateValidPassword(line) {
			counter++
		}
	}
	return counter
}

func calculateValidPassword(line string) bool {
	splitLine := strings.Split(line, " ")
	numberRange := strings.Split(splitLine[0], "-")
	low, _ := strconv.Atoi(numberRange[0])
	high, _ := strconv.Atoi(numberRange[1])
	letter, password := strings.TrimSuffix(splitLine[1], ":"), splitLine[2]

	counter := 0
	for _, c := range password {
		if string(c) == letter {
			counter++
		}
	}
	return counter <= high && counter >= low
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
