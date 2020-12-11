package main

import "testing"

func TestGetFuelRecursively(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, test := range tests {
		got := getFuelRecursively(test.input)
		if got != test.want {
			t.Errorf("Fuel for %0.f was incorrect, got: %f, want: %f", test.input, got, test.want)
		}
	}
}
