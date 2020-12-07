package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	specimens "stev.dev/advent2020"
)

// NOTE: I am assuming that the input text is correct!

func main() {
	var lines, err = specimens.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var validPassports = FindNumberOfValidPassports(lines)

	fmt.Println("Number of valid passports:", validPassports)
}

// FindNumberOfValidPassports goes through all entries, parses them into passports, and validates whether a passport is correct
func FindNumberOfValidPassports(entries []string) (validPassports int) {
	var passport = make([]string, 0, 10)
	var totalPassports = 0
	for i := 0; i < len(entries); i++ {
		if i == len(entries)-1 {
			// hack to force the last line to be appeneded - otherwise we skip it as we need two blank lines
			passport = append(passport, entries[i])
		}
		if strings.TrimSpace(entries[i]) == "" || i == len(entries)-1 {
			// we  have a passport - evaluate it
			if PassportIsValid(strings.Join(passport, " ")) {
				validPassports++
			}
			totalPassports++
			passport = make([]string, 0, 10) // reset the  passport
		} else {
			//lines = append(lines, value)
			passport = append(passport, entries[i])
		}
	}
	fmt.Println("Total passports:", totalPassports)
	return validPassports
}

// PassportIsValid returns true if the passport is valid; otherwise, false
func PassportIsValid(passport string) bool {
	if !IsBirthYearValid(passport) {
		return false
	}

	if !IsIssueYearValid(passport) {
		return false
	}

	if !IsExpirationYearValid(passport) {
		return false
	}

	if !IsHeightValid(passport) {
		return false
	}
	if !IsHairColorValid(passport) {
		return false
	}
	if !IsEyeColorValid(passport) {
		return false
	}
	if !IsPassportIdValid(passport) {
		return false
	}

	// part 1 says this is optional
	// if !IsCountryIdValid(passport) {
	// return false
	// }
	return true
}

func IsBirthYearValid(passport string) bool {
	matched, err := regexp.MatchString(`[^\s]?byr:(?:19[2-9]\d|200[0-2])[$\s]?`, passport)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return false
	}
	return true
}

func IsIssueYearValid(passport string) bool {
	matched, err := regexp.MatchString(`[^\s]?iyr:(:?201\d|2020)[$\s]?`, passport)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return false
	}
	return true
}

func IsExpirationYearValid(passport string) bool {
	matched, err := regexp.MatchString(`[^\s]?eyr:(?:202\d|2030)[$\s]?`, passport)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return false
	}
	return true
}

func IsHeightValid(passport string) bool {
	matched, err := regexp.MatchString(`[^\s]?hgt:(?:(?:1[5-8][0-9]cm|19[0-3]cm)|(?:59in|6\din|7[0-6]in))[$\s]?`, passport)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return false
	}
	return true
}

func IsHairColorValid(passport string) bool {
	matched, err := regexp.MatchString(`[^\s]?hcl:#(?:[a-f\d]{6})[$\s]?`, passport)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return false
	}
	return true
}

func IsEyeColorValid(passport string) bool {
	matched, err := regexp.MatchString(`[^\s]?ecl:(?:amb|blu|brn|gry|grn|hzl|oth){1}[$\s]?`, passport)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return false
	}
	return true
}

func IsPassportIdValid(passport string) bool {
	matched, err := regexp.MatchString(`[^\s]?pid:(?:[0-9]{9}(?:[$\s]+|\n|$))`, passport)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return false
	}
	return true
}

func IsCountryIdValid(passport string) bool {
	matched, err := regexp.MatchString(`([^\s]?cid:[$\s]?)`, passport)
	if err != nil {
		log.Fatal(err)
	}
	if !matched {
		return false
	}
	return true
}

// regex stuff

// ([^\s]?byr:[0-9]{4}[$\s]?)
//		match any whitespace, match exactly "byr:", match exactly 4 numbers within 0 to 9
// ([^\s]?iyr:[0-9]{4}[$\s]?)
// ([^\s]?eyr:[0-9]{4}[$\s]?)
// ([^\s]?hgt:[0-9]{2,3}[$\s]?)
// ([^\s]?eyr:[0-9]{4}[$\s]?)
// ([^\s]?hcl:#.{6}[$\s]?)
// ([^\s]?ecl:[a-zA-Z]{3}[$\s]?)
// ([^\s]?pid:[$\s]?)
// optional: ([^\s]?cid:[0-9]{3}[$\s]?)
//
//
