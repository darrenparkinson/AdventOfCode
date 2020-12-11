package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// Point represents a position on the wire grid
type Point struct {
	x float64
	y float64
}

func main() {

	// Definitely not the most efficient when it comes to Big O,
	// but we it works and we can refactor and optimise if we need
	// to, but it's a puzzle! ;)

	start := time.Now()
	input := readInput("input.txt")
	log.Println("read input in", time.Since(start))

	wire1, wire2 := input[0], input[1]

	start = time.Now()
	wire1Layout := getLayout(wire1)
	log.Println("read layout 1 in", time.Since(start))

	start = time.Now()
	wire2Layout := getLayout(wire2)
	log.Println("read layout 2 in", time.Since(start))

	start = time.Now()
	intersections := calculateIntersections(wire1Layout, wire2Layout)
	log.Println("retrieved intersections in", time.Since(start))

	// log.Println("steps:", steps)

	start = time.Now()
	log.Println("nearest point:", nearestPoint(intersections))
	log.Println("retrieved nearest point in", time.Since(start))

}

// getLayout returns a slice of points given an input string
func getLayout(wire string) []Point {
	points := []Point{}
	directions := strings.Split(wire, ",")
	var i, x, y float64

	for _, move := range directions {
		direction := strings.ToUpper(string(move[0]))
		steps, _ := strconv.ParseFloat(move[1:], 64)
		for i = 0; i < steps; i++ {
			switch direction {
			case "R":
				x++
			case "U":
				y++
			case "D":
				y--
			case "L":
				x--
			}
			points = append(points, Point{x, y})
		}
	}
	return points
}

// Yikes, terrible Big O... takes about 25s to complete...
// func calculateIntersections(line1, line2 []Point) []Point {
// 	var intersections []Point
// 	for _, p := range line1 {
// 		for _, q := range line2 {
// 			if p.x == q.x && p.y == q.y {
// 				intersections = append(intersections, p)
// 			}
// 		}
// 	}
// 	return intersections
// }

// take 2 to see if we can perform better -- this one takes about 41ms
func calculateIntersections(line1, line2 []Point) []Point {
	var intersections []Point
	// Map of points and their distance
	points := make(map[Point]float64)
	for _, point := range line1 {
		points[point] = 1 // point.distance()
	}
	for _, point := range line2 {
		if points[point] != 0 {
			intersections = append(intersections, point)
		}
	}

	return intersections
}

// nearestPoint will identify the nearest point to origin given a slice of points
func nearestPoint(line []Point) float64 {
	nearest := line[0].distance()
	for _, point := range line[1:] {
		distance := point.distance()
		if distance < nearest {
			nearest = distance
		}
	}
	return nearest
}

// Manhattan Distance
// Example: Coordinates x = (a,b) y = (c,d)
// Manhattan Distance = | a - c | + | b - d |
// The | means the absolute value of a - c

// md calculates the manhattan distance between two points
func md(v1, v2 Point) float64 {
	return math.Abs(v1.x-v2.x) + math.Abs(v1.y-v2.y)
}

// Because we're starting with x,y at 0,0, the distance is actually always just
// | c | + | d |
// So we could just have that on the point itself

func (p Point) distance() float64 {
	return math.Abs(p.x) + math.Abs(p.y)
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
