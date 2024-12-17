package day03

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func Solve() {
	// Read input file
	input, err := ioutil.ReadFile("inputs/day03.txt")
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
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	fmt.Println(matches)

	sum := 0
	for _, match := range matches {
		X, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		Y, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		sum += X * Y
	}

	return sum
}

func solvePart2(input string) int {
	// Add in do() and don't() to substring matches, as to be handled later
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don\'t\(\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	fmt.Println(matches)

	sum := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
			continue
		}

		if match[0] == "don't()" {
			enabled = false
			continue
		}

		if !enabled {
			continue
		}

		X, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		Y, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		sum += X * Y
	}

	return sum
}
