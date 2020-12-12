package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"time"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
)

func main() {

	start := time.Now()
	input := aocutils.ReadInputAsRows("input.txt")

	result1 := part1(input)
	log.Println("Result1:", result1)
	result2 := part2(input)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

func part1(input []string) float64 {

	re := regexp.MustCompile("^([NSEWLRF])(\\d+)$")

	currentDirection := 90.0 // East
	position := Point{0, 0}

	for _, row := range input {
		instructions := re.FindStringSubmatch(row)
		direction := instructions[1]
		distance, _ := strconv.ParseFloat(instructions[2], 64)

		if direction == "F" {
			switch currentDirection {
			case 0.0:
				direction = "N"
			case 90.0:
				direction = "E"
			case 180.0:
				direction = "S"
			case 270.0:
				direction = "W"
			}
		}

		switch direction {
		case "N":
			position.y -= distance
		case "S":
			position.y += distance
		case "E":
			position.x += distance
		case "W":
			position.x -= distance
		case "L", "R":
			currentDirection = rotate(currentDirection, direction, distance)
		}
		// fmt.Println(direction, distance, position, currentDirection)
	}

	return md(Point{0, 0}, position)
}

func part2(input []string) float64 {
	re := regexp.MustCompile("^([NSEWLRF])(\\d+)$")

	// currentDirection := 90.0 // East
	shipPosition := Point{0, 0}
	wayPointPosition := Point{10, -1}
	wayPointDiff := Point{10, -1}

	for _, row := range input {
		instructions := re.FindStringSubmatch(row)
		direction := instructions[1]
		distance, _ := strconv.ParseFloat(instructions[2], 64)

		switch direction {
		case "N":
			wayPointDiff.y -= distance
			wayPointPosition.y -= distance
		case "S":
			wayPointDiff.y += distance
			wayPointPosition.y += distance
		case "E":
			wayPointDiff.x += distance
			wayPointPosition.x += distance
		case "W":
			wayPointDiff.x -= distance
			wayPointPosition.x -= distance
		case "F":
			shipPosition.x += (wayPointPosition.x - shipPosition.x) * distance
			shipPosition.y += (wayPointPosition.y - shipPosition.y) * distance
			wayPointPosition.x = shipPosition.x + wayPointDiff.x
			wayPointPosition.y = shipPosition.y + wayPointDiff.y
		case "L", "R":
			wayPointDiff = rotatePart2(direction, distance, wayPointDiff)
			wayPointPosition.x = shipPosition.x + wayPointDiff.x
			wayPointPosition.y = shipPosition.y + wayPointDiff.y
		}
		// fmt.Println(direction, distance, shipPosition, wayPointPosition, wayPointDiff)
	}

	return md(Point{0, 0}, shipPosition)
}

func rotatePart2(direction string, amount float64, diff Point) Point {
	// R = N -> E -> S -> W -> N
	// L = N -> W -> S -> E -> N
	times := amount / 90 // number of times to rotate 90, assuming the value will be 90,180,270
	for i := times; i > 0; i-- {
		switch direction {
		case "R":
			y := diff.x
			diff.x = -diff.y
			diff.y = y
		case "L":
			x := diff.y
			diff.y = -diff.x
			diff.x = x
		}
	}
	return diff
}

func rotate(currentDirection float64, rotationDirection string, amount float64) float64 {

	switch rotationDirection {
	case "L":
		currentDirection -= amount
	case "R":
		currentDirection += amount
	}
	if currentDirection < 0 {
		currentDirection += 360
	}
	if currentDirection >= 360 {
		currentDirection -= 360
	}
	return currentDirection
}

// Point represents a position
type Point struct {
	x float64
	y float64
}

func (p Point) distance() float64 {
	return math.Abs(p.x) + math.Abs(p.y)
}

// md calculates the manhattan distance between two points
// in this case though, start point is 0, so could just use X and Y from end
func md(v1, v2 Point) float64 {
	return math.Abs(v1.x-v2.x) + math.Abs(v1.y-v2.y)
}
