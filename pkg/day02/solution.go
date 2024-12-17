package day02

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Solve() {
	// Read input file
	input, err := ioutil.ReadFile("inputs/day02.txt")
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

func isSafe(levels []int) (bool, int) {
	increasing := levels[0] < levels[1]
	safe := true
	for i := range levels {
		if i == 0 {
			continue
		}

		prev := levels[i-1]
		curr := levels[i]

		if prev == curr || absoluteDifferenceInt(prev, curr) > 3 || increasing && prev >= curr || !increasing && prev <= curr {
			return false, i - 1
		}

	}
	return safe, -1
}

func removeElement(n []int, i int) []int {
	nCopy := append(n[:0:0], n...)
	return append(nCopy[:i], nCopy[i+1:]...)
}

func solvePart1(input string) int {
	reports := strings.Split(input, "\n")

	safeReports := 0
	for _, report := range reports[:len(reports)-1] {
		var levels []int
		for _, f := range strings.Fields(report) {
			i, err := strconv.Atoi(f)
			if err == nil {
				levels = append(levels, i)
			}
		}

		safe, _ := isSafe(levels)

		if safe {
			safeReports++
		}

	}

	return safeReports
}

func solvePart2(input string) int {
	reports := strings.Split(input, "\n")

	safeReports := 0
	for _, report := range reports[:len(reports)-1] {
		var levels []int
		for _, f := range strings.Fields(report) {
			i, err := strconv.Atoi(f)
			if err == nil {
				levels = append(levels, i)
			}
		}

		safe, i := isSafe(levels)

		// Check if removing before, at, and after the unsafe value makes the report safe
		if !safe {
			for x := -1; x <= 1; x++ {
				if i+x < 0 || i+x > len(levels)-1 {
					continue
				}

				withoutX, _ := isSafe(removeElement(levels, i+x))

				if withoutX {
					safeReports++
					break
				}
			}
		}

		if safe {
			safeReports++
		}

	}

	return safeReports
}
