package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := ReadInputAsRows("input_test.txt")
	want := 609043
	got := part1(input)
	assert.Equal(t, want, got)
}
func TestPart2(t *testing.T) {
	input := ReadInputAsRows("input_test.txt")
	want := 6742839
	got := part2(input)
	assert.Equal(t, want, got)
}
