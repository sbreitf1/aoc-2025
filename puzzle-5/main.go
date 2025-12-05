package main

// https://adventofcode.com/2025/day/5

import (
	"aoc/helper"
	"fmt"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	ranges, ids := Parse(lines)

	solution1 := CountFreshIDs(ranges, ids)
	fmt.Println("-> part 1:", solution1)

	mergedRanges := MergeRanges(ranges)
	solution2 := SumRanges(mergedRanges)
	fmt.Println("-> part 2:", solution2)
}

type Range struct {
	Min, Max int64
}

func Parse(lines []string) ([]Range, []int64) {
	ranges := make([]Range, 0)
	ids := make([]int64, 0)
	for _, l := range lines {
		parts := strings.Split(l, "-")
		if len(parts) == 2 {
			ranges = append(ranges, Range{Min: helper.ParseInt[int64](parts[0]), Max: helper.ParseInt[int64](parts[1])})
		} else {
			ids = append(ids, helper.ParseInt[int64](parts[0]))
		}
	}
	return ranges, ids
}

func CountFreshIDs(ranges []Range, ids []int64) int {
	var count int
	for _, id := range ids {
		for _, r := range ranges {
			if r.InRange(id) {
				count++
				break
			}
		}
	}
	return count
}

func (r Range) InRange(id int64) bool {
	return id >= r.Min && id <= r.Max
}

func MergeRanges(ranges []Range) []Range {
	for {
		beforeCount := len(ranges)
		ranges = MergeRangesOnce(ranges)
		if beforeCount == len(ranges) {
			break
		}
	}
	return ranges
}

func MergeRangesOnce(ranges []Range) []Range {
	for i := 0; i < len(ranges); i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[i].Overlap(ranges[j]) {
				mergedRanges := make([]Range, 0, len(ranges)-1)
				mergedRanges = append(mergedRanges, ranges[i].Merge(ranges[j]))
				for k := 0; k < len(ranges); k++ {
					if k != i && k != j {
						mergedRanges = append(mergedRanges, ranges[k])
					}
				}
				return mergedRanges
			}
		}
	}
	return ranges
}

func (r Range) Overlap(other Range) bool {
	if other.Max < r.Min {
		return false
	}
	if other.Min > r.Max {
		return false
	}
	return true
}

func (r Range) Merge(other Range) Range {
	return Range{Min: helper.Min(r.Min, other.Min), Max: helper.Max(r.Max, other.Max)}
}

func SumRanges(ranges []Range) int64 {
	var sum int64
	for _, r := range ranges {
		sum += r.Max - r.Min + 1
	}
	return sum
}
