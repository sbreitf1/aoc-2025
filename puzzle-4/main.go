package main

// https://adventofcode.com/2025/day/4

import (
	"aoc/helper"
	"fmt"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	field := helper.LinesToRunes(lines)

	solution1 := len(GetAccessibleRolls(field))
	fmt.Println("-> part 1:", solution1)

	solution2 := CountTotalAccessibleRolls(field)
	fmt.Println("-> part 2:", solution2)
}

func GetAccessibleRolls(field [][]rune) []helper.Vec2D[int] {
	rolls := make([]helper.Vec2D[int], 0)
	for y := 0; y < len(field); y++ {
		for x := 0; x < len(field[y]); x++ {
			p := helper.Vec2D[int]{X: x, Y: y}
			if field[p.Y][p.X] == '@' && CountSurroundingRolls(field, p) < 4 {
				rolls = append(rolls, p)
			}
		}
	}
	return rolls
}

func CountSurroundingRolls(field [][]rune, p helper.Vec2D[int]) int {
	var count int
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if p.X+dx < 0 || p.X+dx >= len(field[0]) {
				continue
			}
			if p.Y+dy < 0 || p.Y+dy >= len(field) {
				continue
			}

			if field[p.Y+dy][p.X+dx] == '@' {
				count++
			}
		}
	}
	return count
}

func CountTotalAccessibleRolls(field [][]rune) int {
	var count int
	for {
		rolls := GetAccessibleRolls(field)
		if len(rolls) == 0 {
			break
		}

		count += len(rolls)
		for _, p := range rolls {
			field[p.Y][p.X] = '.'
		}
	}
	return count
}
