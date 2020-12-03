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

	fmt.Println(lines[0])

}
