package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("input.txt")
	entries := strings.Split(input[0], ",")
	for idx, val := range entries {
		entry := idx + 1
		ival, _ := strconv.Atoi(val)
		// fmt.Println(entry, entry%3)
		if (entry%3 == 0) && (ival == 255) {
			fmt.Println(val)
		}
	}
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
