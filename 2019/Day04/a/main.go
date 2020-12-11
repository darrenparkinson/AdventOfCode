package main

import (
	"fmt"
	"strconv"
)

func main() {
	const RangeMin int = 156218
	const RangeMax int = 652527

	var matchCount int
	for i := RangeMin; i < RangeMax; i++ {
		if checkCombination(i) {
			matchCount++
		}
	}
	fmt.Println("Matches: ", matchCount)
}

func checkCombination(combo int) bool {
	// OK, so, let's go through the rules...
	combostr := strconv.Itoa(combo)

	// 1. Must be six digits
	if len(combostr) != 6 {
		return false
	}

	ones, twos, threes, fours, fives, sixs := combostr[0], combostr[1], combostr[2], combostr[3], combostr[4], combostr[5]
	one, two, three, four, five, six := int(ones), int(twos), int(threes), int(fours), int(fives), int(sixs)

	// 2. Two adjacent digits are same
	if one != two && two != three && three != four && four != five && five != six {
		return false
	}

	// 3. Left to right, the digits never decrease
	if one > two || two > three || three > four || four > five || five > six {
		return false
	}

	return true
}
