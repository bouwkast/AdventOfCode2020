package main

import (
	"log"
	"testing"

	specimens "stev.dev/advent2020"
)

func TestSumOfPart1(t *testing.T) {
	var output = SumOfYesPart1(GetTestInput())

	if output != 11 {
		t.Errorf("SumOfYesPart1 = %d; want 11", output)
	}
}

func TestSumOfPart2(t *testing.T) {
	var output = SumOfYesPart2(GetTestInput())

	if output != 6 {
		t.Errorf("SumOfYesPart2 = %d; want 6", output)
	}
}

func TestGetSumOfGroupPart2(t *testing.T) {
	var input = []string{"abc"}
	var output = GetSumOfGroupPart2(input)
	if output != 3 {
		t.Errorf("GetSumOfGroupPart2() = %d; want 3", output)
	}
}

func TestGetSumOfGroupPart2_b(t *testing.T) {
	var input = []string{"a", "b", "c"}
	var output = GetSumOfGroupPart2(input)
	if output != 0 {
		t.Errorf("GetSumOfGroupPart2() = %d; want 0", output)
	}
}

func TestGetSumOfGroupPart2_c(t *testing.T) {
	var input = []string{"ab", "ac"}
	var output = GetSumOfGroupPart2(input)
	if output != 1 {
		t.Errorf("GetSumOfGroupPart2() = %d; want 1", output)
	}
}

func TestGetSumOfGroupPart2_d(t *testing.T) {
	var input = []string{"a", "a", "a", "a"}
	var output = GetSumOfGroupPart2(input)
	if output != 1 {
		t.Errorf("GetSumOfGroupPart2() = %d; want 1", output)
	}
}

func GetTestInput() []string { // this isn't the best to use our file reading but makes life easy
	var lines, err = specimens.ReadFile("sample_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	return lines
}
