package main

// https://adventofcode.com/2025/day/2

import (
	"aoc/helper"
	"fmt"
	"math"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	ranges := ParseRanges(lines)

	solution1 := SumInvalidIDs(ranges)
	fmt.Println("-> part 1:", solution1)

	solution2 := 0
	fmt.Println("-> part 1:", solution2)
}

type Range struct {
	Min, Max int
}

func ParseRanges(lines []string) []Range {
	ranges := make([]Range, 0)
	for _, l := range lines {
		parts := strings.Split(l, ",")
		for _, p := range parts {
			parts := strings.Split(p, "-")
			if len(parts) == 2 {
				min := helper.ParseInt[int](parts[0])
				max := helper.ParseInt[int](parts[1])
				ranges = append(ranges, Range{Min: min, Max: max})
			}
		}
	}
	return ranges
}

func SumInvalidIDs(ranges []Range) int64 {
	var sum int64
	for _, r := range ranges {
		invalidIDs := GetInvalidIDs(r)
		for _, id := range invalidIDs {
			sum += int64(id)
		}
	}
	return sum
}

func GetInvalidIDs(idRange Range) []int {
	minLen := helper.Digits(idRange.Min)
	maxLen := helper.Digits(idRange.Max)

	invalidIDs := make([]int, 0)
	for l := minLen; l <= maxLen; l++ {
		if l%2 == 0 {
			pow := int(math.Pow10(l / 2))
			minID := int(math.Pow10(l/2 - 1))
			maxID := 10*minID - 1
			for idPart := minID; idPart <= maxID; idPart++ {
				id := pow*idPart + idPart
				if id < idRange.Min {
					continue
				}
				if id > idRange.Max {
					break
				}
				invalidIDs = append(invalidIDs, id)
			}
		}
	}

	return invalidIDs
}
