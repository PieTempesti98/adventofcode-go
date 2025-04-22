package day1

import (
	"fmt"
	"pietempesti/advent/internal/runner"
	"strings"
)

func init() {
	runner.Register(2015, 1, Solution{})
}

type Solution struct{}

func (s Solution) Part1(input string) (string, error) {
	// Count # of occurrences of '(' and ')', the difference is the floor
	plus := strings.Count(input, "(")
	minus := strings.Count(input, ")")

	floor := plus - minus
	return fmt.Sprintf("%d", floor), nil
}

func (s Solution) Part2(input string) (string, error) {
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
			index := i + 1
			return fmt.Sprintf("%d", index), nil
		}
	}
	return fmt.Sprintf("%d", len(input)), nil
}
