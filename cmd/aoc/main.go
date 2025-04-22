package main

import (
	"flag"
	"fmt"
	"pietempesti/advent/internal/common"
	"pietempesti/advent/internal/runner"
	"time"

	_ "pietempesti/advent/solutions/2015/day1"
	_ "pietempesti/advent/solutions/2015/day2"
)

func main() {
	year := flag.Int("year", time.Now().Year(), "Soutions year")
	day := flag.Int("day", 0, "Solution day (0 for each day)")
	flag.Parse()

	//If day is specified, run the execution only for that day
	if *day > 0 {
		executeDay(*year, *day)
		return
	}

	// Otherwise run the execution for each day of the year
	for d := 1; d <= 25; d++ {
		if _, exists := runner.Solutions[*year][d]; exists {
			executeDay(*year, d)
		}
	}
}

func executeDay(year, day int) {
	solution, exists := runner.Solutions[year][day]
	if !exists {
		fmt.Printf("No solutions for year %d day %d\n", year, day)
		return
	}

	inputPath := fmt.Sprintf("inputs/%d/day%d.txt", year, day)
	input, err := common.ReadInput(inputPath)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return
	}

	fmt.Printf("=== Year %d, Day %d ===\n", year, day)

	start := time.Now()
	result1, err := solution.Part1(input)
	duration1 := time.Since(start)
	if err != nil {
		fmt.Printf("Error in part 1: %v\n", err)
	} else {
		fmt.Printf("Part 1: %s [%v]\n", result1, duration1)
	}

	start = time.Now()
	result2, err := solution.Part2(input)
	duration2 := time.Since(start)
	if err != nil {
		fmt.Printf("Errorin part 2: %v\n", err)
	} else {
		fmt.Printf("Part 2: %s [%v]\n", result2, duration2)
	}

	fmt.Println()
}
