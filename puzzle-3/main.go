package main

// https://adventofcode.com/2025/day/3

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	banks := ParseBatteryBanks(lines)

	solution1 := SumMaxJoltages(banks, 2)
	fmt.Println("-> part 1:", solution1)

	solution2 := SumMaxJoltages(banks, 12)
	fmt.Println("-> part 2:", solution2)
}

type BatteryBank []int

func ParseBatteryBanks(lines []string) []BatteryBank {
	banks := make([]BatteryBank, 0)
	for _, l := range lines {
		bank := make(BatteryBank, 0, len(l))
		for _, r := range l {
			bank = append(bank, int(r-'0'))
		}
		banks = append(banks, bank)
	}
	return banks
}

func (b BatteryBank) GetMaxJoltage(count int) int64 {
	var joltage int64

	lastPos := -1
	for i := (count - 1); i >= 0; i-- {
		maxVal, maxPos := b.FindMaxValInRange(lastPos+1, len(b)-i-1)
		lastPos = maxPos
		joltage += helper.Pow(10, int64(i)) * int64(maxVal)
	}

	return joltage
}

func (b BatteryBank) FindMaxValInRange(start, end int) (int, int) {
	var maxVal, maxPos int
	for i := start; i <= end; i++ {
		if b[i] > maxVal {
			maxVal = b[i]
			maxPos = i
		}
	}
	return maxVal, maxPos
}

func SumMaxJoltages(banks []BatteryBank, count int) int64 {
	var sum int64
	for _, b := range banks {
		sum += b.GetMaxJoltage(count)
	}
	return sum
}
