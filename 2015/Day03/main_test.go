package main

import (
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {

	tests := []struct {
		input string
		want  int
	}{
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}

	for _, test := range tests {
		got := part1(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}
func TestPart2(t *testing.T) {

	tests := []struct {
		input string
		want  int
	}{
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}

	for _, test := range tests {
		got := part2(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}
