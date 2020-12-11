package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	start := time.Now()
	input := readInput("input.txt")

	result := calculateProduct(input)
	log.Println(result)

	log.Println("completed in ", time.Since(start))
}

func calculateProduct(input []string) int {

	for _, i := range input {
		for _, j := range input {
			for _, k := range input {

				a, _ := strconv.Atoi(i)
				b, _ := strconv.Atoi(j)
				c, _ := strconv.Atoi(k)
				if a+b+c == 2020 {
					return a * b * c
				}
			}
		}
	}
	return 0
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
