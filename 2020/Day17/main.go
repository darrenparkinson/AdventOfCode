package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
)

// Point represents a 3 dimensional point in space.
type Point struct {
	x, y, z int
}

func (p Point) String() string {
	return fmt.Sprintf("%d, %d, %d", p.x, p.y, p.z)
}

func main() {

	input := aocutils.ReadInputRaw("input.txt")
	grid := map[Point]bool{} // active = true, inactive = false
	for row, points := range strings.Fields(input) {
		for pos, point := range points {
			var val bool = point == '#'
			grid[Point{x: pos, y: row}] = val // sets x,y since z is already zero
		}
	}
	// fmt.Println(grid)
	for i, p := range grid {
		fmt.Println(i, p)
	}

	start := time.Now()
	result1 := part1(input)
	log.Println("Result1:", result1)
	log.Println("completed in ", time.Since(start))

	start = time.Now()
	result2 := part2(input)
	log.Println("Result2:", result2)
	log.Println("completed in ", time.Since(start))
}

func part1(input string) int {

	return len(input)
}

func part2(input string) int {
	return len(input)
}
