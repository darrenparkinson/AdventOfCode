package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	input := readInput("input.txt")

	start := time.Now()
	boardingPasses := part1(input)
	log.Println("part 1 completed in ", time.Since(start))

	start = time.Now()
	part2(boardingPasses)
	log.Println("part 2 completed in ", time.Since(start))

}

func part2(bps []BoardingPass) {
	sort.Slice(bps, func(i, j int) bool {
		return bps[i].SeatID < bps[j].SeatID
	})
	previous := bps[0].SeatID - 1
	log.Println(previous)
	for _, bp := range bps {
		if (bp.SeatID-1 != previous) && bp.SeatID+1 != (previous+2) {
			log.Println("SeatID:", bp.SeatID-1)
		}
		previous = bp.SeatID
		// log.Println(bp.SeatID)
	}
}

func part1(input []string) []BoardingPass {
	var bps []BoardingPass
	for _, row := range input {
		b := NewBoardingPass(row)
		bps = append(bps, *b)
	}
	// fmt.Println(bps)
	var highest int64 = 0
	for _, bp := range bps {
		if bp.SeatID > highest {
			highest = bp.SeatID
		}
	}
	log.Println("Highest SeatID:", highest)
	return bps
}

type BoardingPass struct {
	Code   string
	Bin    string
	Row    int64
	Col    int64
	SeatID int64
}

// NewBoardingPass calculates the SeatID, Row and Column for a given Boarding pass given the code
func NewBoardingPass(code string) *BoardingPass {
	bp := &BoardingPass{Code: code}
	r := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	bp.Bin = r.Replace(code)
	bp.SeatID, _ = strconv.ParseInt(bp.Bin, 2, 64)
	bp.Row, _ = strconv.ParseInt(bp.Bin[:7], 2, 64)
	bp.Col, _ = strconv.ParseInt(bp.Bin[7:], 2, 64)
	return bp
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
