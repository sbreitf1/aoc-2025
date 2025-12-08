package main

// https://adventofcode.com/2025/day/6

import (
	"aoc/helper"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	problems := ParseProblems(lines)

	solution1 := SumResults(problems)
	fmt.Println("-> part 1:", solution1)

	solution2 := 0
	fmt.Println("-> part 2:", solution2)
}

type Problem struct {
	Numbers  []int64
	Operator string
}

func (p Problem) Calc() int64 {
	switch p.Operator {
	case "+":
		result := int64(0)
		for _, num := range p.Numbers {
			result += num
		}
		return result

	case "*":
		result := int64(1)
		for _, num := range p.Numbers {
			result *= num
		}
		return result

	default:
		helper.ExitWithMessage("operator %q unknown", p.Operator)
		return 0
	}
}

func ParseProblems(lines []string) []Problem {
	splitPattern := regexp.MustCompile(`\s+`)
	lineParts := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		lineParts[i] = splitPattern.Split(strings.TrimSpace(lines[i]), -1)
		if i > 0 && len(lineParts[i]) != len(lineParts[0]) {
			helper.ExitWithMessage("lines have different number of operands")
		}
	}

	numCount := len(lineParts) - 1
	problemCount := len(lineParts[0])

	problems := make([]Problem, problemCount)
	for i := 0; i < problemCount; i++ {
		nums := make([]int64, numCount)
		for j := 0; j < numCount; j++ {
			nums[j] = helper.ParseInt[int64](lineParts[j][i])
		}
		problems[i] = Problem{
			Numbers:  nums,
			Operator: lineParts[numCount][i],
		}
	}
	return problems
}

func SumResults(problems []Problem) int64 {
	var sum int64
	for _, p := range problems {
		sum += p.Calc()
	}
	return sum
}
