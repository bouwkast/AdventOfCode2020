package main

import (
	"strconv"
	"testing"
)

func TestIsBirthYearValid(t *testing.T) {
	var min = 1920
	var max = 2002
	for i := min; i <= max; i++ {
		var input = "byr:" + strconv.Itoa(i)
		valid := IsBirthYearValid(input)
		if !valid {
			t.Errorf("IsBirthYearValid(%s) = %t; want true", input, valid)
		}
	}
}

func TestIsIssueYearValid(t *testing.T) {
	var min = 2000
	var max = 2009
	for i := min; i <= max; i++ {
		var input = "iyr:" + strconv.Itoa(i)
		valid := IsIssueYearValid(input)
		if valid {
			t.Errorf("IsIssueYearValid(%s) = %t; want false", input, valid)
		}
	}
}

func TestIsHeightValid(t *testing.T) {
	var min = 59
	var max = 76
	for i := min; i <= max; i++ {
		var input = "hgt:" + strconv.Itoa(i) + "in"
		valid := IsHeightValid(input)
		if !valid {
			t.Errorf("IsHeightValid(%s) = %t; want true", input, valid)
		}
	}
}

func TestIsPassportIdValid(t *testing.T) {
	valid := IsPassportIdValid("pid:6533951177")
	if valid {
		t.Errorf("IsPassportIdValid(%s) = %t; want false", "pid:6533951177", valid)
	}
}
