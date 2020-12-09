package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	specimens "stev.dev/advent2020"
)

func main() {
	var lines, err = specimens.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var highestSeatId = GetHighestSeatId(lines)
	fmt.Println("Highest seat ID:", highestSeatId)

	var mySeatId = GetMySeatId(lines)
	fmt.Println("My seat ID:", mySeatId)
}

func GetMySeatId(lines []string) int {
	var seatMap map[int]int
	seatMap = make(map[int]int)
	for i := 0; i < len(lines); i++ {
		var seatID = GetSeatId(lines[i])
		seatMap[seatID] = seatID
	}

	// i don't really understand how to do this part w/o brute forcing it
	for row := 1; row < 127; row++ {
		for col := 0; col < 8; col++ {
			var seatID = CalculateSeatId(row, col)

			// check to see if this seat ID does not exist in our map
			_, found := seatMap[seatID]
			if !found {
				_, previous := seatMap[seatID-1]
				_, next := seatMap[seatID+1]
				if previous && next {
					return seatID
				}
			}

		}
	}

	return -1 // error - hate this func, really bad implementation in my opinion but it works
}

func GetHighestSeatId(lines []string) int {
	var highestSeatId = 0

	for i := 0; i < len(lines); i++ {
		var seatID = GetSeatId(lines[i])
		if seatID > highestSeatId {
			highestSeatId = seatID
		}
	}
	return highestSeatId
}

// F = 0
// B = 1

// L = 0
// R = 1

func GetSeatId(input string) int {
	var row = GetRow(input[0:7])
	var col = GetCol(input[7:len(input)])

	return CalculateSeatId(row, col)
}

func CalculateSeatId(row int, col int) int {
	return row*8 + col
}

func GetRow(input string) int {
	// Example): FBFBBFF
	// the first seven characters
	input = strings.ReplaceAll(input, "F", "0")
	input = strings.ReplaceAll(input, "B", "1")

	row, err := strconv.ParseInt(input, 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	return int(row)
}

func GetCol(input string) int {
	// the last three characters
	input = strings.ReplaceAll(input, "L", "0")
	input = strings.ReplaceAll(input, "R", "1")

	col, err := strconv.ParseInt(input, 2, 64)

	if err != nil {
		log.Fatal(err)
	}

	return int(col)
}
