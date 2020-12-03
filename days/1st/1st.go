package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var entries, err = ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	// find two entries that sum to 2020
	var value1, value2 = FindSumValue(entries, 2020)

	// need to return those numbers multiplied by eachother
	fmt.Println(value1 * value2)
}

func FindSumValue(entries []int, expectedValue int) (value1 int, value2 int) {
	for i := 0; i < len(entries)-1; i++ { // -1 is cause we don't want the last number here
		for j := i + 1; j < len(entries); j++ { // start j at i+1 cause we need  it to be different numbers
			if entries[i]+entries[j] == expectedValue {
				return entries[i], entries[j]
			}
		}
	}
	return value1, value2
}

func ReadFile(filePath string) (entries []int, errorCode error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // this will not execute until this ReadFile function returns

	// a warning with this method is that this scanner doesn't handle large lines
	scanner := bufio.NewScanner(file)

	// make a slice of entries - documentation says it is like an array of unfixed length
	entries = make([]int, 0, 10) // makes a slice of ints, with a length of 0 and a cap of 10
	var count int = 0
	for scanner.Scan() {
		var value, err = strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		// append the value to the slice entries
		entries = append(entries, value)

		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return entries, err
}
