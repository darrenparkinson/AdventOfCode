package main

import (
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
	re := regexp.MustCompile("^(.+) (.+,.+) through (.+,.+)$")
	for _, row := range input {
		commandMatch := re.FindStringSubmatch(row)
		command := commandMatch[1]
		rect := Rect{pointFromString(commandMatch[2]), pointFromString(commandMatch[3])}
		performCommand(lights, command, rect)
	}
	lightCount := 0
	for _, lightStatus := range lights {
		if lightStatus {
			lightCount++
		}
	}
	return lightCount
}

func part2(input []string) int {
	lights := initialiseLightGrid2(1000)
	re := regexp.MustCompile("^(.+) (.+,.+) through (.+,.+)$")
	for _, row := range input {
		commandMatch := re.FindStringSubmatch(row)
		command := commandMatch[1]
		rect := Rect{pointFromString(commandMatch[2]), pointFromString(commandMatch[3])}
		performCommand2(lights, command, rect)
	}
	totalBrightness := 0
	for _, brightness := range lights {
		totalBrightness += brightness
	}
	return totalBrightness
}

func performCommand2(lights map[Point]int, command string, rect Rect) {
	for i := rect.topleft.x; i < rect.bottomright.x+1; i++ {
		for j := rect.topleft.y; j < rect.bottomright.y+1; j++ {
			switch command {
			case "turn on":
				lights[Point{i, j}]++
			case "turn off":
				lights[Point{i, j}]--
				if lights[Point{i, j}] < 0 {
					lights[Point{i, j}] = 0
				}
			case "toggle":
				lights[Point{i, j}] += 2
			}
		}
	}
}

func performCommand(lights map[Point]bool, command string, rect Rect) {

	for i := rect.topleft.x; i < rect.bottomright.x+1; i++ {
		for j := rect.topleft.y; j < rect.bottomright.y+1; j++ {
			switch command {
			case "turn on":
				lights[Point{i, j}] = true
			case "turn off":
				lights[Point{i, j}] = false
			case "toggle":
				lights[Point{i, j}] = !lights[Point{i, j}]
			}
		}
	}
}

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

func initialiseLightGrid(num int) map[Point]bool {
	lights := make(map[Point]bool)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			lights[Point{i, j}] = false
		}
	}
	return lights
}

func initialiseLightGrid2(num int) map[Point]int {
	lights := make(map[Point]int)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			lights[Point{i, j}] = 0
		}
	}
	return lights
}
