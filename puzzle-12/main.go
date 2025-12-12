package main

// https://adventofcode.com/2025/day/12

import (
	"aoc/helper"
	"fmt"
	"regexp"
	"sync"
	"sync/atomic"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	present, regions := Parse(lines)

	solution1 := CountFittingRegions(present, regions) // 448
	fmt.Println("-> part 1:", solution1)

	solution2 := 0
	fmt.Println("-> part 2:", solution2)
}

type Present struct {
	Shapes     []Shape
	SolidCount int
}

type Shape [][]int

type Region struct {
	Size          helper.Vec2D[int]
	PresentCounts []int
}

func Parse(lines []string) ([]Present, []Region) {
	patternPresentIndex := regexp.MustCompile(`^\s*(\d+)\s*:\s*$`)
	patternPresentBlocks := regexp.MustCompile(`^\s*([.#]+)\s*$`)
	patternTreeRegion := regexp.MustCompile(`^\s*(\d+)x(\d+)\s*:([0-9 ]+)`)

	presents := make([]Present, 0)
	regions := make([]Region, 0)

	for _, l := range lines {
		if m := patternPresentIndex.FindStringSubmatch(l); len(m) == 2 {
			presents = append(presents, Present{Shapes: []Shape{Shape{}}})
		} else if m := patternPresentBlocks.FindStringSubmatch(l); len(m) == 2 {
			row := helper.MapValuesIgnoreUnknown([]rune(m[1]), map[rune]int{'.': 0, '#': 1})
			presents[len(presents)-1].Shapes[0] = append(presents[len(presents)-1].Shapes[0], row)
		} else if m := patternTreeRegion.FindStringSubmatch(l); len(m) == 4 {
			regions = append(regions, Region{
				Size:          helper.Vec2D[int]{X: helper.ParseInt[int](m[1]), Y: helper.ParseInt[int](m[2])},
				PresentCounts: helper.ExtractInts[int](m[3]),
			})
		}
	}

	for i := range presents {
		var solidCount int
		for y := 0; y < len(presents[i].Shapes[0]); y++ {
			for x := 0; x < len(presents[i].Shapes[0][y]); x++ {
				solidCount += presents[i].Shapes[0][y][x]
			}
		}
		presents[i].SolidCount = solidCount

		presents[i].Shapes = PreparePresentVariations(presents[i].Shapes[0])
	}

	return presents, regions
}

func PreparePresentVariations(baseShape Shape) []Shape {
	variations := make([]Shape, 0)

	addVariation := func(shape Shape) {
		for _, s := range variations {
			if s.Equals(shape) {
				return
			}
		}
		variations = append(variations, shape)
	}

	for i := 0; i < 4; i++ {
		addVariation(baseShape)
		addVariation(baseShape.FlipX())
		baseShape = baseShape.RotCW()
	}
	return variations
}

func (shape Shape) FlipX() Shape {
	return Shape{
		[]int{shape[0][2], shape[0][1], shape[0][0]},
		[]int{shape[1][2], shape[1][1], shape[1][0]},
		[]int{shape[2][2], shape[2][1], shape[2][0]},
	}
}

func (shape Shape) RotCW() Shape {
	return Shape{
		[]int{shape[2][0], shape[1][0], shape[0][0]},
		[]int{shape[2][1], shape[1][1], shape[0][1]},
		[]int{shape[2][2], shape[1][2], shape[0][2]},
	}
}

func (shape Shape) Equals(other Shape) bool {
	for y := 0; y < len(shape); y++ {
		for x := 0; x < len(shape[y]); x++ {
			if shape[y][x] != other[y][x] {
				return false
			}
		}
	}
	return true
}

func CountFittingRegions(presents []Present, regions []Region) int {
	var count int32
	remaining := int32(len(regions))
	var wg sync.WaitGroup
	for _, r := range regions {
		wg.Go(func() {
			if Fits(presents, r) {
				atomic.AddInt32(&count, 1)
			}

			atomic.AddInt32(&remaining, -1)
			fmt.Println(remaining, "remaining ->", count)
		})
	}
	wg.Wait()
	return int(count)
}

func Fits(presents []Present, region Region) bool {
	// early out
	var totalSolid int
	for i := range region.PresentCounts {
		totalSolid += presents[i].SolidCount * region.PresentCounts[i]
	}
	if totalSolid > region.Size.X*region.Size.Y {
		// presents take more space than this region has available
		return false
	}

	return fits(presents, helper.NewRuneField(region.Size.X, region.Size.Y, '.'), helper.Vec2D[int]{X: 0, Y: 0}, helper.Clone(region.PresentCounts))
}

func fits(presents []Present, field helper.Field[rune], p helper.Vec2D[int], remainingPresents []int) bool {
	for y := helper.Max(0, p.Y-3); y < helper.Min(field.Height()-2, p.Y+3); y++ {
		for x := max(0, p.X-3); x < helper.Min(field.Width()-2, p.X+3); x++ {
			allZero := true
			for pi := range remainingPresents {
				if remainingPresents[pi] > 0 {
					allZero = false

					for vi := range presents[pi].Shapes {
						if canPlace(presents[pi].Shapes[vi], field, x, y) {
							remainingPresents[pi]--
							place(presents[pi].Shapes[vi], field, x, y, rune('A'+pi))

							if fits(presents, field, helper.Vec2D[int]{X: x, Y: y}, remainingPresents) {
								return true
							}

							remainingPresents[pi]++
							unplace(presents[pi].Shapes[vi], field, x, y)
						}
					}
				}
			}

			if allZero {
				// we have placed all presents
				return true
			}
		}
	}
	return false
}

func canPlace(shape Shape, field helper.Field[rune], x, y int) bool {
	for dy := 0; dy < len(shape); dy++ {
		for dx := 0; dx < len(shape[dy]); dx++ {
			if shape[dy][dx] > 0 {
				if field.AtXY(x+dx, y+dy) != '.' {
					return false
				}
			}
		}
	}
	return true
}

func place(shape Shape, field helper.Field[rune], x, y int, val rune) {
	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 3; dx++ {
			if shape[dy][dx] > 0 {
				field.SetXY(x+dx, y+dy, val)
			}
		}
	}
}

func unplace(shape Shape, field helper.Field[rune], x, y int) {
	for dy := 0; dy < 3; dy++ {
		for dx := 0; dx < 3; dx++ {
			if shape[dy][dx] > 0 {
				field.SetXY(x+dx, y+dy, '.')
			}
		}
	}
}
