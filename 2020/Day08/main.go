package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	start := time.Now()
	input := readInput("input.txt")

	result1, _ := part1(input)
	log.Println("Result1:", result1)
	result2 := part2(input)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string) (int, bool) {
	inc := 0
	rownum := 0
	finished := false
	visited := make(map[int]bool)
	for {
		if _, ok := visited[rownum]; ok {
			// fmt.Println("visited")
			break
		}
		if rownum >= len(input) {
			finished = true
			break
		}
		visited[rownum] = true
		op := strings.Split(input[rownum], " ")
		num, _ := strconv.Atoi(op[1])
		switch op[0] {
		case "acc":
			inc += num
			rownum++
		case "nop":
			rownum++
		case "jmp":
			rownum += num

		}
	}
	return inc, finished
}

func part2(input []string) int {
	var jmpLines []int
	for i, line := range input {
		lineParts := strings.Fields(line)
		if lineParts[0] == "jmp" {
			jmpLines = append(jmpLines, i)
		}
	}

	for _, jmpLineIndex := range jmpLines {
		trialLines := make([]string, len(input))
		copy(trialLines, input)
		trialLines[jmpLineIndex] = "nop 0"
		inc, finished := part1(trialLines)
		if finished {
			return inc
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
