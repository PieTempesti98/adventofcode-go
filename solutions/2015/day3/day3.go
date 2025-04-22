package day3

import (
	"fmt"
	"pietempesti/advent/internal/runner"
)

func init() {
	runner.Register(2015, 3, Solution{})
}

type Solution struct{}

func (s Solution) Part1(input string) (string, error) {

	x, y := 1, 1
	positions := make(map[string]struct{})
	strPos := fmt.Sprintf("%d-%d", x, y)
	positions[strPos] = struct{}{}
	for _, c := range input {
		switch c {
		case '>':
			x = x + 1
		case '<':
			x = x - 1
		case '^':
			y = y + 1
		case 'v':
			y = y - 1
		}
		strPos := fmt.Sprintf("%d-%d", x, y)
		positions[strPos] = struct{}{}
	}
	return fmt.Sprintf("%d", len(positions)), nil
}

func (s Solution) Part2(input string) (string, error) {
	x, y := 1, 1
	roboX, roboY := 1, 1
	isRobo := false
	positions := make(map[string]struct{})
	strPos := fmt.Sprintf("%d-%d", x, y)
	positions[strPos] = struct{}{}
	for _, c := range input {
		switch c {
		case '>':
			if !isRobo {
				x += 1
				strPos = fmt.Sprintf("%d-%d", x, y)
			} else {
				roboX += 1
				strPos = fmt.Sprintf("%d-%d", roboX, roboY)
			}
		case '<':
			if !isRobo {
				x -= 1
				strPos = fmt.Sprintf("%d-%d", x, y)
			} else {
				roboX -= 1
				strPos = fmt.Sprintf("%d-%d", roboX, roboY)
			}
		case '^':
			if !isRobo {
				y += 1
				strPos = fmt.Sprintf("%d-%d", x, y)
			} else {
				roboY += 1
				strPos = fmt.Sprintf("%d-%d", roboX, roboY)
			}
		case 'v':
			if !isRobo {
				y -= 1
				strPos = fmt.Sprintf("%d-%d", x, y)
			} else {
				roboY -= 1
				strPos = fmt.Sprintf("%d-%d", roboX, roboY)
			}
		}
		positions[strPos] = struct{}{}
		isRobo = !isRobo
	}
	return fmt.Sprintf("%d", len(positions)), nil
}
