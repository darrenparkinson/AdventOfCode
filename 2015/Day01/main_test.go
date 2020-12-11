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
		{"(())", 0},
		{"()()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}

	for _, test := range tests {
		got := identifyFloor(test.input)
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
		{")", 1},
		{"()())", 5},
	}

	for _, test := range tests {
		got := identifyBasement(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}
