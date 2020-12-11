package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// RowsToIntSlice returns a slice of int given a slice of strings
func RowsToIntSlice(input []string) []int {
	entries := make([]int, len(input))
	var err error
	for i, row := range input {
		entries[i], err = strconv.Atoi(row)
	}
	if err != nil {
		log.Fatal(err)
	}
	return entries
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
