package main

import (
	"fmt"
	"log"

	specimens "stev.dev/advent2020"
)

func main() {
	var lines, err = specimens.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: Number of yes:", SumOfYesPart1(lines))
	fmt.Println("Part 2: Number of yes:", SumOfYesPart2(lines))
}

func SumOfYesPart1(lines []string) int {
	var sum = 0

	var group = make([]string, 0, 10)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || i == len(lines)-1 {
			if i == len(lines)-1 {
				group = append(group, lines[i])
			}
			sum += GetSumOfGroupPart1(group)
			group = make([]string, 0, 10)
		} else {
			group = append(group, lines[i])
		}
	}

	return sum
}

func GetSumOfGroupPart1(group []string) int {
	var answers map[byte]bool
	answers = make(map[byte]bool)
	var sum = 0

	for i := 0; i < len(group); i++ {
		for j := 0; j < len(group[i]); j++ {
			answers[group[i][j]] = true
		}
	}

	for _, element := range answers {
		if element {
			sum++
		}
	}

	return sum
}

func SumOfYesPart2(lines []string) int { // TODO this is duped from part 1 besides the get sum of group
	var sum = 0

	var group = make([]string, 0, 10)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || i == len(lines)-1 {
			if i == len(lines)-1 {
				group = append(group, lines[i])
			}

			sum += GetSumOfGroupPart2(group)
			group = make([]string, 0, 10)
		} else {
			group = append(group, lines[i])
		}
	}

	return sum
}

func GetSumOfGroupPart2(group []string) int {
	var answers map[byte]int
	answers = make(map[byte]int)
	var sum = 0

	for i := 0; i < len(group); i++ {
		for j := 0; j < len(group[i]); j++ {
			answers[group[i][j]] = answers[group[i][j]] + 1
		}
	}

	for _, element := range answers {
		if element == len(group) {
			sum++
		}
	}

	return sum
}
