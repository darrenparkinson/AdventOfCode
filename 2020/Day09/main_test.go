package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := ReadInputAsRows("input_test.txt")
	want := 127
	got := part1(input, 5)
	assert.Equal(t, want, got)
}
func TestPart2(t *testing.T) {
	input := ReadInputAsRows("input_test.txt")
	want := 62
	got := part2(input, 5)
	assert.Equal(t, want, got)
}
