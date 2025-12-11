package main

// https://adventofcode.com/2025/day/9

import (
	"aoc/helper"
	"fmt"
	"math"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	redTiles := ParseRedTiles(lines)

	solution1 := FindLargestRect(redTiles)
	fmt.Println("-> part 1:", solution1)

	solution2 := FindLargestRect2(redTiles)
	fmt.Println("-> part 2:", solution2)
}

func ParseRedTiles(lines []string) []helper.Vec2D[int] {
	redTiles := make([]helper.Vec2D[int], 0, len(lines))
	for _, l := range lines {
		ints := helper.ExtractInts[int](l)
		redTiles = append(redTiles, helper.Vec2D[int]{X: ints[0], Y: ints[1]})
	}
	return redTiles
}

func FindLargestRect(redTiles []helper.Vec2D[int]) int64 {
	var largestArea int64
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			area := int64(helper.Abs(redTiles[i].X-redTiles[j].X)+1) * int64(helper.Abs(redTiles[i].Y-redTiles[j].Y)+1)
			if area > largestArea {
				largestArea = area
			}
		}
	}
	return largestArea
}

func FindLargestRect2(redTiles []helper.Vec2D[int]) int64 {
	grid := prepareGrid(redTiles)
	emptyLineRanges := prepareEmptyLineRanges(redTiles, grid)

	var largestArea int64
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			size := helper.Vec2D[int]{X: helper.Abs(redTiles[i].X-redTiles[j].X) + 1, Y: helper.Abs(redTiles[i].Y-redTiles[j].Y) + 1}
			area := int64(size.X) * int64(size.Y)
			if area > largestArea {
				ul := helper.Vec2D[int]{X: helper.Min(redTiles[i].X, redTiles[j].X), Y: helper.Min(redTiles[i].Y, redTiles[j].Y)}
				if !checkRectTouchesEmpty(emptyLineRanges, ul, size) {
					largestArea = area
				}
			}
		}
	}
	return largestArea
}

func checkRectTouchesEmpty(emptyLineRanges map[int][]Range, ul, size helper.Vec2D[int]) bool {
	for y := ul.Y; y < ul.Y+size.Y; y++ {
		for _, r := range emptyLineRanges[y] {
			if (r.EndX >= ul.X) && (r.StartX <= (ul.X + size.X - 1)) {
				return true
			}
		}
	}
	return false
}

func prepareGrid(redTiles []helper.Vec2D[int]) map[helper.Vec2D[int]]rune {
	grid := make(map[helper.Vec2D[int]]rune)
	for i := 0; i < len(redTiles); i++ {
		delta := redTiles[(i+1)%len(redTiles)].Sub(redTiles[i])
		dir := delta.Normalized()
		for j := 0; j < int(delta.Len()); j++ {
			p := redTiles[i].Add(dir.Mul(j))
			if j == 0 {
				grid[p] = '#'
			} else {
				grid[p] = 'X'
			}
		}
	}
	return grid
}

type Range struct {
	StartX, EndX int
}

func prepareEmptyLineRanges(redTiles []helper.Vec2D[int], grid map[helper.Vec2D[int]]rune) map[int][]Range {
	min, max := findMinMax(redTiles)
	emptyLineRanges := make(map[int][]Range)

	addRange := func(y, startX, endX int) {
		if startX < endX {
			emptyLineRanges[y] = append(emptyLineRanges[y], Range{StartX: startX, EndX: endX})
		}
	}

	for y := min.Y; y <= max.Y; y++ {
		const stateNone = 0
		const stateOutside = 1
		const stateOnLine = 2
		const stateInside = 3

		state := stateNone
		startX := -1

		for x := min.X; x <= max.X; x++ {
			val, ok := grid[helper.Vec2D[int]{X: x, Y: y}]

			switch state {
			case stateNone:
				if !ok {
					state = stateOutside
					startX = x
				} else if ok && val == 'X' {
					state = stateInside
				} else if ok && val == '#' {
					state = stateOnLine
				}

			case stateOutside:
				if ok && val == 'X' {
					addRange(y, startX, x-1)
					state = stateInside
				} else if ok && val == '#' {
					addRange(y, startX, x-1)
					state = stateOnLine
				}

			case stateInside:
				if ok && val == 'X' {
					state = stateOutside
					startX = x + 1
				} else if ok && val == '#' {
					state = stateOnLine
				}

			case stateOnLine:
				if ok && val == '#' {
					fmt.Println("check")
					if isInside(redTiles, helper.Vec2D[int]{X: x + 1, Y: y}) {
						state = stateInside
					} else {
						state = stateOutside
						startX = x + 1
					}
				} else if !ok {
					helper.ExitWithMessage("invalid state transition")
				}
			}
		}

		if state == stateOutside {
			addRange(y, startX, max.X)
		}
	}
	return emptyLineRanges
}

func isInside(redTiles []helper.Vec2D[int], p helper.Vec2D[int]) bool {
	var windingNumber float64
	for i := 0; i < len(redTiles); i++ {
		windingNumber += computeWindingNumberOfLine(p, redTiles[i], redTiles[(i+1)%len(redTiles)])
	}
	return windingNumber < -0.5 || windingNumber > 0.5
}

func computeWindingNumberOfLine(p, l1, l2 helper.Vec2D[int]) float64 {
	if p.X == l1.X && p.X == l2.X {
		return 0
	}
	if p.Y == l1.Y && p.Y == l2.Y {
		return 0
	}
	a1 := math.Atan2(float64(p.Y)-float64(l1.Y), float64(p.X)-float64(l1.X))
	a2 := math.Atan2(float64(p.Y)-float64(l2.Y), float64(p.X)-float64(l2.X))
	diff := a2 - a1
	if diff >= math.Pi {
		diff -= 2 * math.Pi
	}
	if diff <= -math.Pi {
		diff += 2 * math.Pi
	}
	return diff
}

func findMinMax(redTiles []helper.Vec2D[int]) (helper.Vec2D[int], helper.Vec2D[int]) {
	min := helper.Vec2D[int]{X: 100000000, Y: 100000000}
	max := helper.Vec2D[int]{X: -100000000, Y: -100000000}
	for _, t := range redTiles {
		min.X = helper.Min(min.X, t.X)
		min.Y = helper.Min(min.Y, t.Y)
		max.X = helper.Max(max.X, t.X)
		max.Y = helper.Max(max.Y, t.Y)
	}
	return min, max
}
