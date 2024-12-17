package day05

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	// Read input file
	input, err := ioutil.ReadFile("inputs/day05.txt")
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

type Key struct {
	X, Y int
}

func solvePart1(input string) int {
	sections := strings.Split(input, "\n\n")

	// 1. Compose rules
	r, _ := regexp.Compile(`(\d{2})\|(\d{2})`)
	rulesPairs := r.FindAllStringSubmatch(sections[0], -1)

	rules := make(map[Key]bool)
	for _, rulePair := range rulesPairs {
		x, err := strconv.Atoi(rulePair[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(rulePair[2])
		if err != nil {
			panic(err)
		}
		rules[Key{x, y}] = true
	}

	// 2. Compose updates
	updateLines := strings.Split(sections[1], "\n")
	updateLines = updateLines[:len(updateLines)-1]
	updates := make([][]int, len(updateLines))
	for x, updateLine := range updateLines {
		// Split the string on commas and store as int
		strs := strings.Split(updateLine, ",")

		// Convert each string to int
		ints := make([]int, len(strs))
		for i, str := range strs {
			ints[i], _ = strconv.Atoi(str)
		}
		updates[x] = ints
	}

	// 3. Verify updates
	sum := 0
	for _, update := range updates {
		updateValid := isUpdateValid(update, rules)
		if updateValid {
			middle := len(update) / 2
			sum += update[middle]
		}
	}

	return sum
}

func isUpdateValid(update []int, rules map[Key]bool) bool {
	updateValid := true
	for i := 0; i < len(update)-1; i++ {
		for j := i; j < len(update); j++ {
			if i == j {
				continue
			}
			exists := rules[Key{update[i], update[j]}]

			if !exists {
				updateValid = false
			}
		}
	}
	return updateValid
}

func solvePart2(input string) int {
	sections := strings.Split(input, "\n\n")

	// 1. Compose rules
	r, _ := regexp.Compile(`(\d{2})\|(\d{2})`)
	rulesPairs := r.FindAllStringSubmatch(sections[0], -1)

	rules := make(map[Key]bool)
	for _, rulePair := range rulesPairs {
		x, err := strconv.Atoi(rulePair[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(rulePair[2])
		if err != nil {
			panic(err)
		}
		rules[Key{x, y}] = true
	}

	// 2. Compose updates
	updateLines := strings.Split(sections[1], "\n")
	updateLines = updateLines[:len(updateLines)-1]
	updates := make([][]int, len(updateLines))
	for x, updateLine := range updateLines {
		strs := strings.Split(updateLine, ",")
		ints := make([]int, len(strs))
		for i, str := range strs {
			ints[i], _ = strconv.Atoi(str)
		}
		updates[x] = ints
	}

	// 3. Verify updates
	sum := 0
	for _, update := range updates {
		updateValid := isUpdateValid(update, rules)
		if !updateValid {
			// Sort the update using custom comparator
			sort.Slice(update, func(i, j int) bool {
				return rules[Key{update[i], update[j]}]
			})

			middle := len(update) / 2
			sum += update[middle]
		}
	}

	return sum
}
