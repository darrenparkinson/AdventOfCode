package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := ReadInputAsRows("input_test.txt")
	want := 2
	got := part1(input)
	assert.Equal(t, want, got)
}
func TestPart2(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, test := range tests {
		got := isValid2(test.input)
		assert.Equal(t, test.want, got)
	}

}
