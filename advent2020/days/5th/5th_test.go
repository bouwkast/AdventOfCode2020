package main

import (
	"testing"
)

func TestGetHighestSeatId(t *testing.T) {
	var input = []string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"}
	var output = GetHighestSeatId(input)

	if output != 820 {
		t.Errorf("GetHighestSeatId(%s) = %d; want 820", input, output)
	}
}

func TestGetSeatId(t *testing.T) {
	var input = "FBFBBFFRLR"
	var output = GetSeatId(input)

	if output != 357 {
		t.Errorf("GetSeatId(%s) = %d; want 357", input, output)
	}
}

func TestGetRow(t *testing.T) {
	var input = "FBFBBFF"
	var output = GetRow(input)

	if output != 44 {
		t.Errorf("GetRow(%s) = %d; want 44", input, output)
	}
}

func TestGetCol(t *testing.T) {
	var input = "RLR"
	var output = GetCol(input)

	if output != 5 {
		t.Errorf("GetCol(%s) = %d; want 8", input, output)
	}
}
