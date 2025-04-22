package day2

import (
	"fmt"
	"math"
	"pietempesti/advent/internal/runner"
	"strconv"
	"strings"
)

func init() {
	runner.Register(2015, 2, Solution{})
}

type Solution struct{}

func (s Solution) Part1(input string) (string, error) {
	// Parse the string per line
	lines := strings.Split(input, "\n")
	var tot int
	for _, line := range lines {
		// Obtain dimensions
		l, w, h, err := parseLine(line)
		if err != nil {
			return "", err
		}

		// Compute the minimum of the three dimensions
		min := int(math.Min(float64(w*h), math.Min(float64(l*w), float64(h*l))))
		// Compute the surface of the present
		surface := (2*l*w + 2*w*h + 2*l*h) + min
		tot += surface
	}

	return fmt.Sprintf("%d", tot), nil
}

func (s Solution) Part2(input string) (string, error) {
	// Parse the string per line
	lines := strings.Split(input, "\n")
	var tot int
	for _, line := range lines {
		// Obtain dimensions
		l, w, h, err := parseLine(line)
		if err != nil {
			return "", err
		}

		// Compute the minimum perimeter
		min := int(math.Min(float64(2*w+2*h), math.Min(float64(2*l+2*w), float64(2*h+2*l))))
		// Compute the measure of the ribbon
		surface := (w * l * h) + min
		tot += surface
	}

	return fmt.Sprintf("%d", tot), nil
}

func parseLine(line string) (l, w, h int, err error) {
	dimensions := strings.Split(line, "x")
	l, err = strconv.Atoi(dimensions[0])
	if err != nil {
		return 0, 0, 0, err
	}
	w, err = strconv.Atoi(dimensions[1])
	if err != nil {
		return 0, 0, 0, err
	}
	h, err = strconv.Atoi(dimensions[2])
	if err != nil {
		return 0, 0, 0, err
	}
	return l, w, h, nil
}
