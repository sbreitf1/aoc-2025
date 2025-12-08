package main

// https://adventofcode.com/2025/day/7

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	field := helper.NewRuneFieldFromLines(lines)

	solution1, solution2 := CountBeamSplits(&field, field.MustFindOnce('S').X)
	fmt.Println("-> part 1:", solution1)
	fmt.Println("-> part 2:", solution2)
}

func CountBeamSplits(field *helper.Field[rune], startX int) (int, int64) {
	var totalSplitCount int
	beams := map[int]int64{startX: 1}
	for y := 0; y < field.Height(); y++ {
		newBeams := make(map[int]int64)
		for x, c := range beams {
			if field.AtXY(x, y) == '^' {
				totalSplitCount++
				newBeams[x-1] = newBeams[x-1] + c
				newBeams[x+1] = newBeams[x+1] + c
			} else {
				newBeams[x] = newBeams[x] + c
			}
		}
		beams = helper.Clone(newBeams)
	}
	var timelineCount int64
	for _, c := range beams {
		timelineCount += c
	}
	return totalSplitCount, timelineCount
}
