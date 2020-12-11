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
	}

	for _, test := range tests {
		got := checkCombination(test.input)
		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
