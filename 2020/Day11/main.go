package main

import (
	"log"
	"time"
)

func main() {

	start := time.Now()
	input := ReadInputAsRows("input.txt")

	result1 := part1(input)
	log.Println("Result1:", result1)
	result2 := part2(input)
	log.Println("Result2:", result2)

	log.Println("completed in ", time.Since(start))
}

type Seat struct {
	row, col int
}

func part2(input []string) int {
	seats := parseSeats(input)
	plan := make(map[Seat]bool)

	maxRows, maxCols := len(input), len(input[0])

	// set all seats to empty (empty = true)
	for _, seat := range seats {
		plan[seat] = true
	}
	seatsWithInterestingSeats := make(map[Seat][]Seat)
	for _, seat := range seats {
		seatsWithInterestingSeats[seat] = addInterestingSeats(seat, maxRows, maxCols, plan)
	}

	parseCount := 1
	for {
		if checkPlan2(seatsWithInterestingSeats, plan) == true {
			parseCount++
			continue
		}
		break
	}
	occupiedSeatCount := 0
	for _, seat := range seats {
		if plan[seat] == false {
			occupiedSeatCount++
		}
	}
	return occupiedSeatCount

}

func checkPlan2(seats map[Seat][]Seat, plan map[Seat]bool) bool {
	changeCount := 0
	tempPlan := make(map[Seat]bool)
	for k, v := range plan {
		tempPlan[k] = v
	}
	for k, v := range seats {
		adjacentCountUnoccupied := 0
		adjacentCountOccupied := 0
		seatCount := len(v)
		for _, seat := range v {
			if tempPlan[seat] {
				adjacentCountUnoccupied++
			}
			if !tempPlan[seat] {
				adjacentCountOccupied++
			}
		}
		if tempPlan[k] == true && seatCount == adjacentCountUnoccupied {
			plan[k] = false
			changeCount++
		}
		if tempPlan[k] != true && seatCount >= 5 && adjacentCountOccupied >= 5 {
			plan[k] = true
			changeCount++
		}
	}
	if changeCount > 0 {
		return true
	}
	return false
}

func addInterestingSeats(s Seat, maxRows, maxCols int, plan map[Seat]bool) []Seat {
	// interesting seats are only the "first seat they can see in each direction"
	interestingSeats := make(map[Seat]bool)

	// This is horrible, there must be a better way...
	// Seat negative x
	for x := s.col; x >= 0; x-- {
		if x == s.col {
			continue
		}
		if v, ok := plan[Seat{s.row, x}]; ok {
			interestingSeats[Seat{s.row, x}] = v
			break
		}
	}
	for x := s.col; x < maxCols; x++ {
		if x == s.col {
			continue
		}
		if v, ok := plan[Seat{s.row, x}]; ok {
			interestingSeats[Seat{s.row, x}] = v
			break
		}
	}
	// y
	for y := s.row; y >= 0; y-- {
		if y == s.row {
			continue
		}
		if v, ok := plan[Seat{y, s.col}]; ok {
			interestingSeats[Seat{y, s.col}] = v
			break
		}
	}
	for y := s.row; y < maxRows; y++ {
		if y == s.row {
			continue
		}
		if v, ok := plan[Seat{y, s.col}]; ok {
			interestingSeats[Seat{y, s.col}] = v
			break
		}
	}
	// ↘
	x, y := s.col, s.row
	for {
		if x > maxCols-1 || y > maxRows-1 {
			break
		}
		x++
		y++
		if v, ok := plan[Seat{y, x}]; ok {
			interestingSeats[Seat{y, x}] = v
			break
		}
	}
	// ↖
	x, y = s.col, s.row
	for {
		if x < 0 || y < 0 {
			break
		}
		x--
		y--
		if v, ok := plan[Seat{y, x}]; ok {
			interestingSeats[Seat{y, x}] = v
			break
		}
	}
	// ↗
	x, y = s.col, s.row
	for {
		if x > maxCols-1 || y < 0 {
			break
		}
		x++
		y--
		if v, ok := plan[Seat{y, x}]; ok {
			interestingSeats[Seat{y, x}] = v
			break
		}
	}
	// ↙
	x, y = s.col, s.row
	for {
		if x < 0 || y > maxRows-1 {
			break
		}
		x--
		y++
		if v, ok := plan[Seat{y, x}]; ok {
			interestingSeats[Seat{y, x}] = v
			break
		}
	}
	var returnSeats []Seat
	for seat := range interestingSeats {
		returnSeats = append(returnSeats, seat)
	}
	return returnSeats
}

func part1(input []string) int {

	seats := parseSeats(input)
	plan := make(map[Seat]bool)

	// set all seats to empty (empty = true)
	for _, seat := range seats {
		plan[seat] = true
	}
	parseCount := 1
	for {
		if checkPlan(seats, plan) == true {
			parseCount++
			continue
		}
		break
	}
	occupiedSeatCount := 0
	for _, seat := range seats {
		if plan[seat] == false {
			occupiedSeatCount++
		}
	}
	return occupiedSeatCount
}

func checkPlan(seats []Seat, plan map[Seat]bool) bool {
	changeCount := 0
	tempPlan := make(map[Seat]bool)
	for k, v := range plan {
		tempPlan[k] = v
	}
	for _, seat := range seats {
		// e.g. not the floor
		if current, found := plan[seat]; found {
			changed := checkSeat(seat, seats, tempPlan)
			if changed != current {
				changeCount++
			}
			plan[seat] = changed
		}
	}
	if changeCount > 0 {
		return true
	}
	return false
}
func checkSeat(seat Seat, seats []Seat, plan map[Seat]bool) bool {

	interestingSeats := []Seat{
		{seat.row - 1, seat.col - 1},
		{seat.row - 1, seat.col},
		{seat.row - 1, seat.col + 1},
		{seat.row, seat.col - 1},
		{seat.row, seat.col + 1},
		{seat.row + 1, seat.col - 1},
		{seat.row + 1, seat.col},
		{seat.row + 1, seat.col + 1},
	}
	adjacentCountUnoccupied := 0
	adjacentCountOccupied := 0
	seatCount := 0
	for _, iSeat := range interestingSeats {
		if empty, found := plan[iSeat]; found {
			seatCount++
			if empty {
				adjacentCountUnoccupied++
			}
			if !empty {
				adjacentCountOccupied++
			}
		}
	}
	// if the seat is empty and all adjacent seats are NOT occupied
	if plan[seat] == true && seatCount == adjacentCountUnoccupied {
		return false
	}
	// if seat is occupied and four mor more seats adjacent are occupied
	if plan[seat] != true && seatCount >= 4 && adjacentCountOccupied >= 4 {
		return true
	}
	// otherwise return unchanged
	return plan[seat]
}

func parseSeats(input []string) []Seat {
	var seats []Seat
	for i, row := range input {
		for j, col := range row {
			if string(col) == "L" {
				seats = append(seats, Seat{i, j})
			}
		}
	}
	return seats
}
