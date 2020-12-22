package main

import (
	"testing"

	"github.com/darrenparkinson/AdventOfCode/pkg/aocutils"
	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := aocutils.ReadInputAsRows("input_test.txt")
	want := 5
	got := part1(input)
	assert.Equal(t, want, got)
}

func TestPart2(t *testing.T) {
	input := aocutils.ReadInputAsRows("input_test.txt")
	want := "mxmxvkd,sqjhc,fvjkl"
	got := part2(input)
	assert.Equal(t, want, got)
}
