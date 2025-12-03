package main

// https://adventofcode.com/2025/day/1

import (
	"aoc/helper"
	"fmt"
	"regexp"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	rotations := ParseRotations(lines)

	solution1, solution2 := CountZeros(rotations)
	fmt.Println("-> part 1:", solution1)
	fmt.Println("-> part 2:", solution2)
}

func ParseRotations(lines []string) []int {
	patternRotation := regexp.MustCompile(`^([LR])(\d+)$`)

	rotations := make([]int, 0, len(lines))
	for _, l := range lines {
		m := patternRotation.FindStringSubmatch(l)
		if len(m) > 0 {
			sign := helper.MustMapValue(m[1], map[string]int{"L": -1, "R": 1})
			distance := helper.ParseInt[int](m[2])
			rotations = append(rotations, sign*distance)
		}
	}
	return rotations
}

func CountZeros(rotations []int) (int, int) {
	var count1, count2 int
	pos := 50
	for _, r := range rotations {
		sign := helper.Sign(r)
		for i := 0; i < helper.Abs(r); i++ {
			pos = (pos + 100 + sign) % 100
			if pos == 0 {
				count2++
			}
		}
		if pos == 0 {
			count1++
		}
	}
	return count1, count2
}
