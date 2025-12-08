package main

// https://adventofcode.com/2025/day/7

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	field := helper.NewRuneFieldFromLines(lines)

	solution1 := CountBeamSplits(&field, field.MustFindOnce('S').X)
	fmt.Println("-> part 1:", solution1)

	solution2 := 0
	fmt.Println("-> part 2:", solution2)
}

func CountBeamSplits(field *helper.Field[rune], startX int) int {
	var totalSplitCount int
	beams := []int{startX}
	for y := 0; y < field.Height(); y++ {
		newBeams := make(map[int]int)
		for _, b := range beams {
			if field.AtXY(b, y) == '^' {
				totalSplitCount++
				newBeams[b-1] = newBeams[b-1] + 1
				newBeams[b+1] = newBeams[b+1] + 1
			} else {
				newBeams[b] = newBeams[b] + 1
			}
		}
		beams = helper.GetKeySlice(newBeams)
	}
	return totalSplitCount
}
