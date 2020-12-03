package advent2020

// TODO - i would like to figure out how to test this :)

import (
	"bufio"
	"log"
	"os"
)

// ReadFile reads the input file and returns an array of string that represent each line of the input
func ReadFile(filePath string) (lines []string, errorCode error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // this will not execute until this ReadFile function returns

	// a warning with this method is that this scanner doesn't handle large lines
	scanner := bufio.NewScanner(file)

	// make a slice of lines - documentation says it is like an array of unfixed length
	lines = make([]string, 0, 10) // makes a slice of string, with a length of 0 and a cap of 10
	var count int = 0
	for scanner.Scan() {

		var value = scanner.Text()
		// append the value to the slice lines
		lines = append(lines, value)

		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, err
}
