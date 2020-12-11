package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestPart1(t *testing.T) {
// 	input := ReadInputAsRows("input_test.txt")
// 	want := 37
// 	got := part1(input)
// 	assert.Equal(t, want, got)
// }

func TestPart2(t *testing.T) {
	input := ReadInputAsRows("input_test.txt")
	want := 26
	got := part2(input)
	assert.Equal(t, want, got)
}
