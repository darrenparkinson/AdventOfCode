package aocutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInputAsRows(t *testing.T) {
	expected := []string{"a", "b", "c", "d"}
	result := ReadInputAsRows("input_test.txt")
	assert.Equal(t, expected, result)
}

func TestReadInputRaw(t *testing.T) {
	expected := "a\nb\nc\nd"
	result := ReadInputRaw("input_test.txt")
	assert.Equal(t, expected, result)
}

func TestReadInputAsInts(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	result := ReadInputAsInts("input_int_test.txt")
	assert.Equal(t, expected, result)
}

func TestStringToIntSlice(t *testing.T) {
	input := "1321131112"
	want := []int{1, 3, 2, 1, 1, 3, 1, 1, 1, 2}
	got := StringToIntSlice(input)
	assert.Equal(t, want, got)
}

func TestMinMaxIntSlice(t *testing.T) {
	input := []int{100, 135, 120, 110, 91, 31, 111, 211, 121, 12}
	wantmin, wantmax := 12, 211
	gotmin, gotmax := MinMaxIntSlice(input)
	assert.Equal(t, wantmin, gotmin)
	assert.Equal(t, wantmax, gotmax)
}
