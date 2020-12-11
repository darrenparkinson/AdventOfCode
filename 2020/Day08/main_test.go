package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := readInput("input_test.txt")
	want := 5
	got, _ := part1(input)
	assert.Equal(t, want, got)
}
func TestPart2(t *testing.T) {
	input := readInput("input_test.txt")
	want := 8
	got := part2(input)
	assert.Equal(t, want, got)
}
