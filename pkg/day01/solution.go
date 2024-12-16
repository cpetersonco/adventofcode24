package day01

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Solve() {
	// Read input file
	input, err := ioutil.ReadFile("pkg/day01/inputs/day01.txt")
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	// Part 1
	part1Result := solvePart1(string(input))
	fmt.Printf("Part 1: %d\n", part1Result)

	// Part 2
	part2Result := solvePart2(string(input))
	fmt.Printf("Part 2: %d\n", part2Result)
}

func solvePart1(input string) int {
	// Implement part 1 solution
	return 0
}

func solvePart2(input string) int {
	// Implement part 2 solution
	return 0
}
