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

	solution1 := SumInvalidIDsPart1(ranges)
	fmt.Println("-> part 1:", solution1)

	solution2 := SumInvalidIDsPart2(ranges)
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

func SumInvalidIDsPart1(ranges []Range) int64 {
	var sum int64
	for _, r := range ranges {
		invalidIDs := GetInvalidIDs(r, 2)
		for _, id := range invalidIDs {
			sum += int64(id)
		}
	}
	return sum
}

func SumInvalidIDsPart2(ranges []Range) int64 {
	var sum int64
	for _, r := range ranges {
		invalidIDs := make([]int, 0)
		for i := 2; i < 16; i++ {
			invalidIDs = append(invalidIDs, GetInvalidIDs(r, i)...)
		}
		invalidIDs = helper.GetUniqueValues(invalidIDs)
		for _, id := range invalidIDs {
			sum += int64(id)
		}
	}
	return sum
}

func GetInvalidIDs(idRange Range, numParts int) []int {
	minLen := helper.Digits(idRange.Min)
	maxLen := helper.Digits(idRange.Max)

	invalidIDs := make([]int, 0)
	for l := minLen; l <= maxLen; l++ {
		if l%numParts == 0 {
			partLen := l / numParts
			pow := int(math.Pow10(partLen))
			minID := int(math.Pow10(partLen - 1))
			maxID := 10*minID - 1
			for idPart := minID; idPart <= maxID; idPart++ {
				var id int
				for i := 0; i < numParts; i++ {
					id += int(math.Pow(float64(pow), float64(i))) * idPart
				}

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
