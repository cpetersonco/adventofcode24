package day01

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	// Read input file
	input, err := ioutil.ReadFile("inputs/day01.txt")
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

func absoluteDifferenceInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func solvePart1(input string) int {
	lines := strings.Fields(input)

	// 1. Separate the left and right distances
	var left []int
	var right []int
	for i, value := range lines {
		distance, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			left = append(left, distance)
		} else {
			right = append(right, distance)
		}
	}

	// 2. Sort the slices
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	// 3. Sum each distance
	totalDistance := 0
	for i := range left {
		totalDistance += absoluteDifferenceInt(left[i], right[i])
	}

	return totalDistance
}

func solvePart2(input string) int {
	lines := strings.Fields(input)

	// 1. Separate the left and right distances
	var left []int
	var right []int
	for i, value := range lines {
		distance, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		if i%2 == 0 {
			left = append(left, distance)
		} else {
			right = append(right, distance)
		}
	}

	// 2. Build frequency map of right slice
	frequency := make(map[int]int)
	for _, val := range right {
		frequency[val] += 1
	}

	// 3. Compute similarity score
	similarityScore := 0
	for _, num := range left {
		similarityScore += num * frequency[num]
	}

	return similarityScore
}
