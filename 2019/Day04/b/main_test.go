package main

import (
	"testing"
)

func TestCheckCombo(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
		{122345, true},
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for _, test := range tests {
		got := checkCombination(test.input)
		if got != test.want {
			t.Errorf("%v: got %v, want %v", test.input, got, test.want)
		}
	}
}
