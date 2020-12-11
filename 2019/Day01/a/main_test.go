package main

import "testing"

func TestGetFuel(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, test := range tests {
		got := getFuel(test.input)
		if got != test.want {
			t.Errorf("Fuel for %0.f was incorrect, got: %f, want: %f", test.input, got, test.want)
		}
	}
}
