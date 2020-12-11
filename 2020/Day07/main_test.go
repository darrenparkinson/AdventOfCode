package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := readInput("input_test.txt")
	got := part1(input)
	want := 4
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := readInput("input_test.txt")
	got := part2(input)
	want := 32
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
