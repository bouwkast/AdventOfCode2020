package main

import (
	"fmt"
	"log"

	specimens "stev.dev/advent2020"
)

// NOTE: I am assuming that the input text is correct!

func main() {
	var lines, err = specimens.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var slope1 = SolveInstance(lines, 1, 1)
	var slope2 = SolveInstance(lines, 3, 1)
	var slope3 = SolveInstance(lines, 5, 1)
	var slope4 = SolveInstance(lines, 7, 1)
	var slope5 = SolveInstance(lines, 1, 2)

	var overallAnswer = slope1 * slope2 * slope3 * slope4 * slope5
	fmt.Println("Part 2 answer:", overallAnswer)
}

// SolveInstance will solve the given instance and output the answer
func SolveInstance(lines []string, right int, down int) (result int) {
	result = CalculateNumberOfTreesWeHit(lines, right, down)
	fmt.Println("Number of trees hit:", result, "\tright:", right, "\tdown:", down)
	return result
}

// CalculateNumberOfTreesWeHit will iterate over each line and compute whether we hit a tree with our toboggan on that line, supply the slope with right and down
func CalculateNumberOfTreesWeHit(lines []string, right int, down int) (treesHit int) {
	var index = 0
	for line := down; line < len(lines); line += down { // we skip line 0 as we can't start on a tree
		if lines[line][(index+right)%len(lines[line])] == []byte("#")[0] {
			treesHit++
		}
		index = (index + right) % len(lines[line])
	}

	return treesHit
}
