package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// StringToIntSlice returns a slice of Ints given a string of ints.
func StringToIntSlice(input string) []int {
	result := make([]int, len(input))
	for i, r := range input {
		s, err := strconv.Atoi(string(r))
		if err != nil {
			log.Fatal(err)
		}
		result[i] = s
	}
	return result
}

// MinMaxIntSlice returns the min and max of a slice of ints
func MinMaxIntSlice(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

// ReadInputAsRows reads the input file and returns a slice of strings containing each line
func ReadInputAsRows(filename string) []string {
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

// ReadInputRaw returns a string of the entire input file
func ReadInputRaw(path string) string {

	byteContent, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	content := string(byteContent)
	return content
}

// ReadInputAsRows reads the input file and returns a slice of ints containing each line
func ReadInputAsInts(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []int
	for scanner.Scan() {
		s := scanner.Text()
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, i)
	}
	return lines
}
