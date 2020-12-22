package main

import (
	"testing"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := aocutils.ReadInputAsRows("input_test.txt")
	want := []int{1, 10, 27, 78, 438, 1836}
	for i, row := range input {
		got := part1(row)
		assert.Equal(t, want[i], got)
	}
}
func TestPart2(t *testing.T) {
	input := aocutils.ReadInputAsRows("input_test.txt")
	want := []int{2578, 3544142, 261214, 6895259, 18, 362}
	for i, row := range input {
		got := part2(row)
		assert.Equal(t, want[i], got)
	}
}

var result1 int

func BenchmarkPart1(b *testing.B) {
	var r int
	for x := 1; x <= b.N; x++ {
		input := "7,14,0,17,11,1,2"
		part1(input)
	}
	result1 = r
}

var result2 int

func BenchmarkPart2(b *testing.B) {
	var r int
	for x := 1; x <= b.N; x++ {
		input := "7,14,0,17,11,1,2"
		r = part2(input)
	}
	result2 = r
}
