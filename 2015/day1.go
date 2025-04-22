package main

import (
	"fmt"
	"os"
	"strings"
)

const path = "2015/inputs/1.txt"

func main() {

	input := parseInput(path)

	fmt.Println(part1(input))

	fmt.Println(part2(input))

}

func parseInput(path string) string {

	data, err := os.ReadFile(path)

	if err != nil {
		panic("error parsing input file " + path)
	}

	return string(data)
}

func part1(input string) int {
	// Count # of occurrences of '(' and ')', the difference is the floor
	plus := strings.Count(input, "(")
	minus := strings.Count(input, ")")

	return plus - minus
}

func part2(input string) int {
	// For loop with an accumulator, when the accumulator is < 1 we should break the cycle and return the index + 1

	floor := 0
	for i, char := range input {
		if char == '(' {
			floor++
		}
		if char == ')' {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return len(input)
}
