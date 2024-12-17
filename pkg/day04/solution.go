package day04

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Solve() {
	// Read input file
	input, err := ioutil.ReadFile("inputs/day04.txt")
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

func countXMAS(x, y int, puzzle [][]rune) int {
	count := 0
	xmas := []rune("XMAS")
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			inWord := true
			for k := 0; k < 4; k++ {
				a := x + (i * k)
				b := y + (j * k)
				if a < 0 || b < 0 || a > len(puzzle)-1 || b > len(puzzle[a])-1 {
					inWord = false
					break
				}

				if puzzle[a][b] != xmas[k] {
					inWord = false
					break
				}
			}
			if inWord {
				count += 1
			}
		}
	}
	return count
}

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")
	// TODO: Bug fixed, extra line found when splitting testcase
	lines = lines[:len(lines)-1]

	puzzle := make([][]rune, len(lines))
	for i, line := range lines {
		puzzle[i] = []rune(line)
	}

	count := 0
	for i, row := range puzzle {
		for j := range row {
			count += countXMAS(i, j, puzzle)
		}
	}

	return count
}

func isMAS(x, y int, puzzle [][]rune) bool {
	if puzzle[x][y] == 'A' {
		if x-1 < 0 || y-1 < 0 || x+1 > len(puzzle)-1 || y+1 > len(puzzle[x])-1 {
			return false
		}

		if ((puzzle[x-1][y-1] == 'M' && puzzle[x+1][y+1] == 'S') ||
			(puzzle[x-1][y-1] == 'S' && puzzle[x+1][y+1] == 'M')) &&
			((puzzle[x-1][y+1] == 'M' && puzzle[x+1][y-1] == 'S') ||
				(puzzle[x-1][y+1] == 'S' && puzzle[x+1][y-1] == 'M')) {

			return true
		}

	}
	return false
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")
	// TODO: Bug fixed, extra line found when splitting testcase
	lines = lines[:len(lines)-1]

	puzzle := make([][]rune, len(lines))

	for i, line := range lines {
		puzzle[i] = []rune(line)
	}

	count := 0
	for i, row := range puzzle {
		for j := range row {
			if isMAS(i, j, puzzle) {
				count++
			}
		}
	}

	return count
}
