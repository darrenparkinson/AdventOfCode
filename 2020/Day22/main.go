package main

import (
	"strings"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
)

func main() {

	player1, player2 := getPlayers("input.txt")

	// start := time.Now()
	// result1 := part1(input)
	// log.Println("Result1:", result1)
	// log.Println("completed in ", time.Since(start))

	// start = time.Now()
	// result2 := part2(input)
	// log.Println("Result2:", result2)
	// log.Println("completed in ", time.Since(start))

}

func getPlayers(file string) ([]int, []int) {
	input := aocutils.ReadInputRaw(file)
	players := strings.Split(input, "\n\n")
	player1 := strings.Split(players[0], "\n")[1:]
	player2 := strings.Split(players[1], "\n")[1:]
	return aocutils.StringSliceToIntSlice(player1), StringSliceToIntSlice(player2)
}

func part1(input []string) int {
	return len(input)
}

func part2(input []string) int {
	return len(input)
}
