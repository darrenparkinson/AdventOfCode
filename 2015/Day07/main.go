package main

import (
	"fmt"
	"log"
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

type Wire struct {
	left, right string
	op          string
}

func part1(input []string) uint16 {
	re := regexp.MustCompile("^(.+) -> (.+)$")
	gatere := regexp.MustCompile("^(.+) (AND|OR|LSHIFT|RSHIFT) (.+)$")
	notgatere := regexp.MustCompile("^NOT (.+)$")

	wires := make(map[string]Wire)
	memo := make(map[string]uint16)

	rowsDone := 0

	// initialise destinations
	// for _, row := range input {
	// 	commandMatch := re.FindStringSubmatch(row)
	// 	dest := commandMatch[2]
	// 	if _, ok := wires[dest]; !ok {
	// 		wires[dest] = 0
	// 	}
	// }
	for _, row := range input[0:5] {
		commandMatch := re.FindStringSubmatch(row)
		cmd := commandMatch[1]
		dest := commandMatch[2]

		c, err := strconv.ParseInt(cmd, 10, 16)
		if err == nil {
			// we got a number so just assign it
			// fmt.Printf("%016b\n", c)
			// wires[dest] = uint16(c)
			wires[dest] = Wire{}
			rowsDone++
			continue
		}
		// otherwise we now have to work out what to do
		// fmt.Println(cmd, dest)
		gateCmd := gatere.FindStringSubmatch(cmd)
		if gateCmd != nil {
			// we found an AND/OR/LSHIFT/RSHIFT
			l, op, r := gateCmd[1], gateCmd[2], gateCmd[3]
			fmt.Println("MATCHED", l, r, op)
			switch op {
			case "AND":
				wires[dest] = wires[l] & wires[r]
				rowsDone++
			case "OR":
				wires[dest] = wires[l] | wires[r]
				rowsDone++
			case "LSHIFT":
				ri, _ := strconv.ParseInt(r, 10, 16)
				wires[dest] = wires[l] << ri
				rowsDone++
			case "RSHIFT":
				ri, _ := strconv.ParseInt(r, 10, 16)
				wires[dest] = wires[l] >> ri
				rowsDone++
			}
			fmt.Println("wires", dest, "is", wires[dest])
			continue
		}

		notgateCmd := notgatere.FindStringSubmatch(cmd)
		if notgateCmd != nil {
			// we got a NOT
			r := notgateCmd[1]
			fmt.Println("MATCHED NOT", r)
			wires[dest] = ^wires[r]
			rowsDone++
			continue
		}
		fmt.Println("Skipping:", cmd, dest)
		// didn't match anything, so probably just wire to wire
		wires[dest] = wires[cmd]
		fmt.Println(wires[dest])

	}
	fmt.Println(rowsDone)
	return wires["a"]
}

func part2(input []string) int {
	return len(input)
}
