package main

import (
	"testing"
)

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		input Point
		want  float64
	}{
		{Point{3, 3}, 6},
		{Point{6, -6}, 12},
		{Point{-6, 6}, 12},
		{Point{-6, -6}, 12},
		{Point{0, 0}, 0},
	}

	for _, test := range tests {
		got1 := md(Point{0, 0}, test.input)
		got2 := test.input.distance()
		if got1 != test.want || got2 != test.want {
			t.Errorf("got1 %v, got2 %v want %v", got1, got2, test.want)
		}
	}
}

func TestGetLayout(t *testing.T) {
	// TODO 😂
	return
}

func TestNearestPoint(t *testing.T) {
	tests := []struct {
		input1 string
		input2 string
		want   float64
	}{
		{"R8,U5,L5,D3", "U7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	}
	for _, test := range tests {
		// I know I shouldn't mix the functions here, but hey, it's just for fun!!
		layout1 := getLayout(test.input1)
		layout2 := getLayout(test.input2)
		intersections := calculateIntersections(layout1, layout2)
		got := nearestPoint(intersections)
		if got != test.want {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}

// func TestBestSteps(t *testing.T) {
// 	tests := []struct {
// 		input1 string
// 		input2 string
// 		want   float64
// 	}{
// 		{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83", 610},
// 		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 410},
// 	}
// 	for _, test := range tests {
// 		// I know I shouldn't mix the functions here, but hey, it's just for fun!!
// 		layout1 := getLayout(test.input1)
// 		layout2 := getLayout(test.input2)
// 		_, got := calculateIntersections(layout1, layout2)

// 		if got != test.want {
// 			t.Errorf("got %v want %v", got, test.want)
// 		}
// 	}
// }
