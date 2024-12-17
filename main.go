package main

import (
	"fmt"
	"os"

	"github.com/cpetersonco/adventofcode24/pkg/day01"
	"github.com/cpetersonco/adventofcode24/pkg/day02"
	"github.com/cpetersonco/adventofcode24/pkg/day03"
	"github.com/cpetersonco/adventofcode24/pkg/day04"
	"github.com/cpetersonco/adventofcode24/pkg/day05"
	// Import other days as you solve them
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify a day (e.g., day01)")
		os.Exit(1)
	}

	day := os.Args[1]
	switch day {
	case "day01":
		day01.Solve()
	case "day02":
		day02.Solve()
	case "day03":
		day03.Solve()
	case "day04":
		day04.Solve()
	case "day05":
		day05.Solve()
	// Add more cases for other days
	default:
		fmt.Printf("No solution found for %s\n", day)
		os.Exit(1)
	}
}
