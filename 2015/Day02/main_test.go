package main

import (
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	want := []Present{
		{L: 2, W: 3, H: 4, SurfaceArea: 52, PaperRequired: 58, Ribbon: 10, Bow: 24, TotalRibbon: 34},
		{L: 1, W: 1, H: 10, SurfaceArea: 42, PaperRequired: 43, Ribbon: 4, Bow: 10, TotalRibbon: 14}}
	input := readInput("input_test.txt")
	got := parsePresents(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	result1 := part1(got)
	if result1 != 101 {
		t.Errorf("got %v want %v", result1, 94)
	}

}

func TestPart2(t *testing.T) {
	input := readInput("input_test.txt")
	presents := parsePresents(input)
	result2 := part2(presents)
	if result2 != 48 {
		t.Errorf("got %v want %v", result2, 48)
	}
}
