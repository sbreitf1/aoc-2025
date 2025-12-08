package main

// https://adventofcode.com/2025/day/6

import (
	"aoc/helper"
	"fmt"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	grid := ReadGrid(lines)

	solution1 := SumResults(ParseProblems1(grid))
	fmt.Println("-> part 1:", solution1)

	solution2 := SumResults(ParseProblems2(grid))
	fmt.Println("-> part 2:", solution2)
}

func ReadGrid(lines []string) [][]string {
	lineRunes := helper.LinesToRunes(lines)

	rows := make([][]string, len(lineRunes))
	colStart := 0
Outter:
	for i := 0; i <= len(lineRunes[0]); i++ {
		if i < len(lineRunes[0]) {
			for j := 0; j < len(lineRunes); j++ {
				if lineRunes[j][i] != ' ' {
					continue Outter
				}
			}
		}

		for j := 0; j < len(lineRunes); j++ {
			rows[j] = append(rows[j], string(lineRunes[j][colStart:i]))
		}
		colStart = i + 1
	}
	return rows
}

func ParseProblems1(grid [][]string) []Problem {
	problems := make([]Problem, 0, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		nums := make([]int64, len(grid)-1)
		for j := 0; j < len(nums); j++ {
			nums[j] = helper.ParseInt[int64](strings.TrimSpace(grid[j][i]))
		}
		op := strings.TrimSpace(grid[len(grid)-1][i])
		problems = append(problems, Problem{
			Numbers:  nums,
			Operator: op,
		})
	}
	return problems
}

func ParseProblems2(grid [][]string) []Problem {
	problems := make([]Problem, 0, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		numStrs := make([]string, len(grid[0][i]))
		for j := 0; j < len(numStrs); j++ {
			for k := 0; k < len(grid)-1; k++ {
				r := []rune(grid[k][i])[j]
				if r >= '0' && r <= '9' {
					numStrs[j] = numStrs[j] + string(r)
				}
			}
		}

		nums := make([]int64, len(numStrs))
		for j := 0; j < len(numStrs); j++ {
			nums[j] = helper.ParseInt[int64](strings.TrimSpace(numStrs[j]))
		}

		op := strings.TrimSpace(grid[len(grid)-1][i])
		problems = append(problems, Problem{
			Numbers:  nums,
			Operator: op,
		})
	}
	return problems
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

func SumResults(problems []Problem) int64 {
	var sum int64
	for _, p := range problems {
		sum += p.Calc()
	}
	return sum
}
