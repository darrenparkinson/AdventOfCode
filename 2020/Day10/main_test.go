package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := ReadInputAsInts("input2_test.txt")
	want := 220
	got := part1(input)
	assert.Equal(t, want, got)
}
func TestPart2(t *testing.T) {
	input := ReadInputAsInts("input_test.txt")
	want := 8
	got := part2(input)
	assert.Equal(t, want, got)
}

// func TestPart2b(t *testing.T) {
// 	input := ReadInputAsInts("input2_test.txt")
// 	want := 19208
// 	got := part2(input)
// 	assert.Equal(t, want, got)
// }
