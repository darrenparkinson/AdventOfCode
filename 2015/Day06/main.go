package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
)

// Point represents a single location
type Point struct {
	x, y int
}

// Rect represents the top left and bottom right of a rectangle
type Rect struct {
	topleft     Point
	bottomright Point
}

func main() {

	start := time.Now()
	input := aocutils.ReadInputAsRows("input.txt")

	result1 := part1(input)
	log.Println("Result1:", result1)
	result2 := part2(input)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string) int {

	lights := initialiseLightGrid(1000)
	fmt.Println(len(lights))

	re := regexp.MustCompile("^(.+) (.+,.+) through (.+,.+)$")
	// commandMatch := re.FindStringSubmatch(input[7])
	// fmt.Printf("%+v\n", commandMatch[1])
	// fmt.Printf("%+v\n", commandMatch[2])
	// fmt.Printf("%+v\n", commandMatch[3])

	for _, row := range input {
		commandMatch := re.FindStringSubmatch(row)
		command := commandMatch[1]
		rect := Rect{pointFromString(commandMatch[2]), pointFromString(commandMatch[3])}
		// to := strings.Split(commandMatch[3], ",")
		// rect := Rect{strinPoint{command[2]}}
		fmt.Println(command, rect)
	}

	return len(input)
}

func performCommand(lights map[Point]bool, command string, rect Rect)

func pointFromString(s string) Point {
	a := strings.Split(s, ",")
	x, err := strconv.Atoi(a[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(a[1])
	if err != nil {
		log.Fatal(err)
	}
	return Point{x, y}
}

func part2(input []string) int {
	return len(input)
}

func initialiseLightGrid(num int) map[Point]bool {
	lights := make(map[Point]bool)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			lights[Point{i, j}] = true
		}
	}
	return lights
}
