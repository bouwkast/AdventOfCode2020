package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	specimens "stev.dev/advent2020"
)

// NOTE: I am assuming that the input text is correct!

func main() {
	var lines, err = specimens.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var validPasswordsPart1, validPasswordsPart2 = FindNumberOfValidPasswords(lines)

	fmt.Println("PART1 The number of valid passwords:", validPasswordsPart1)
	fmt.Println("PART2 The number of valid passwords:", validPasswordsPart2)
}

// FindNumberOfValidPasswords searches the passwordEntries for valid passwords and returns the total count of valid passwords
func FindNumberOfValidPasswords(passwordEntries []string) (validPasswordsPart1 int, validPasswordsPart2 int) {
	// part 1 calculation
	for i := 0; i < len(passwordEntries); i++ {
		if IsPasswordValidPart1(passwordEntries[i]) {
			validPasswordsPart1++
		}
	}
	// part 2 calculation
	for i := 0; i < len(passwordEntries); i++ {
		if IsPasswordValidPart2(passwordEntries[i]) {
			validPasswordsPart2++
		}
	}
	return validPasswordsPart1, validPasswordsPart2
}

// IsPasswordValidPart1 checks whether the given policy and password is valid, returns true if it is valid, returns false otherwise
func IsPasswordValidPart1(passwordEntry string) bool {
	// example: "1-3 a: abcde"
	//		means "abcde" is the password
	// 		policy is that there must be at least one "a", but not more than three "a"s in the password

	var min, max, combination, password = ParseEntry(passwordEntry)

	// in password, find the number of uses of combination
	// if uses is less than min - false
	// if uses is greater than max - false

	var uses = strings.Count(password, combination)

	if uses < min {
		return false
	}

	if uses > max {
		return false
	}

	return true
}

// IsPasswordValidPart2 checks whether the given policy and password is valid, returns true if it is valid, returns false otherwise
func IsPasswordValidPart2(passwordEntry string) bool {
	// example: "1-3 a: abcde"
	//		means "abcde" is the password
	// 		policy is that there must be exactly one "a" - either in position 1 (index 0) or position 3 (index 2)

	var min, max, combination, password = ParseEntry(passwordEntry)
	// min is the first index, max is the second index - only one index may contain the combination
	// NOTE indices have an assumed +1 to them, so we have to do a -1 to get them to line up correctly with our password

	// X xor Y => (X || Y) && !(X && Y)
	var firstPositionValid = password[min-1] == combination[0]
	var secondPositionValid = password[max-1] == combination[0]

	return (firstPositionValid || secondPositionValid) && !(firstPositionValid && secondPositionValid)
}

// ParseEntry parses the password entry into the min/max policy ints, the required combination, and the password itself
func ParseEntry(passwordEntry string) (min int, max int, requiredCombination string, password string) {

	var split = strings.Split(passwordEntry, " ") // split on the space

	// split is now: {"1-3", "a:", "abcde"}
	var minMax = strings.Split(split[0], "-")

	min, err := strconv.Atoi(minMax[0])
	if err != nil {
		log.Fatal(err)
	}

	max, err = strconv.Atoi(minMax[1])
	if err != nil {
		log.Fatal(err)
	}

	requiredCombination = strings.ReplaceAll(split[1], ":", "")

	password = split[2]
	return min, max, requiredCombination, password
}
