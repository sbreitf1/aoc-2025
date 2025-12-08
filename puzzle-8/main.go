package main

// https://adventofcode.com/2025/day/8

import (
	"aoc/helper"
	"fmt"
	"sort"
	"strings"
)

const (
	printSteps = false
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	junctionBoxes := ParseJunctionBoxes(lines)
	var solution1, solution2 int64
	if len(junctionBoxes) > 100 {
		solution1, solution2 = GetSolutions(junctionBoxes, 1000)
	} else {
		solution1, solution2 = GetSolutions(junctionBoxes, 10)
	}

	fmt.Println("-> part 1:", solution1)
	fmt.Println("-> part 2:", solution2)
}

type JunctionBox helper.Vec3D[int]

func (jb JunctionBox) DistanceTo(other JunctionBox) float64 {
	return helper.Vec3D[int](jb).Sub(helper.Vec3D[int](other)).Len()
}

func ParseJunctionBoxes(lines []string) []JunctionBox {
	lights := make([]JunctionBox, 0, len(lines))
	for _, l := range lines {
		parts := strings.Split(l, ",")
		lights = append(lights, JunctionBox{X: helper.ParseInt[int](parts[0]), Y: helper.ParseInt[int](parts[1]), Z: helper.ParseInt[int](parts[2])})
	}
	return lights
}

type Circuit []JunctionBox

func CountNonEmptyCircuits(circuits []Circuit) int {
	var count int
	for _, c := range circuits {
		if len(c) > 0 {
			count++
		}
	}
	return count
}

func CountJunctionBoxes(circuits []Circuit) int {
	var count int
	for _, c := range circuits {
		count += len(c)
	}
	return count
}

func GetSolutions(junctionBoxes []JunctionBox, part1At int) (int64, int64) {
	pointPairs := PreparePointPairs(junctionBoxes)

	var part1Solution int64

	circuits := make([]Circuit, 0)
	circuitMap := make(map[JunctionBox]int)
	for i, conn := range pointPairs {
		jb1 := junctionBoxes[conn.P1]
		ci1, ok1 := circuitMap[jb1]

		jb2 := junctionBoxes[conn.P2]
		ci2, ok2 := circuitMap[jb2]

		if !ok1 && !ok2 {
			if printSteps {
				fmt.Println(junctionBoxes[conn.P1], junctionBoxes[conn.P2], "-> new")
			}
			// new circuit for these lights
			circuits = append(circuits, Circuit{jb1, jb2})
			circuitMap[jb1] = len(circuits) - 1
			circuitMap[jb2] = len(circuits) - 1
		} else if ok1 && ok2 {
			if ci1 == ci2 {
				if printSteps {
					fmt.Println(junctionBoxes[conn.P1], junctionBoxes[conn.P2], "-> nop")
				}
				// already in same circuit, nothing happens
				continue
			}
			if printSteps {
				fmt.Println(junctionBoxes[conn.P1], junctionBoxes[conn.P2], "-> combine")
			}

			// combine circuits
			circuits[ci1] = append(circuits[ci1], circuits[ci2]...)
			for _, l := range circuits[ci2] {
				circuitMap[l] = ci1
			}
			circuits[ci2] = Circuit{} // cannot remove, this would scramble indices

		} else if ok1 && !ok2 {
			if printSteps {
				fmt.Println(junctionBoxes[conn.P1], junctionBoxes[conn.P2], "-> join 1")
			}
			// join light 2 into circuit of light 1
			circuits[ci1] = append(circuits[ci1], jb2)
			circuitMap[jb2] = ci1
		} else if !ok1 && ok2 {
			if printSteps {
				fmt.Println(junctionBoxes[conn.P1], junctionBoxes[conn.P2], "-> join 2")
			}
			// join light 1 into circuit of light 2
			circuits[ci2] = append(circuits[ci2], jb1)
			circuitMap[jb1] = ci2
		} else {
			// wtf?!?!
			helper.ExitWithMessage("you forgot something!")
		}

		if (i + 1) == part1At {
			part1Solution = GetSolution1(circuits)
		}

		if CountNonEmptyCircuits(circuits) == 1 && CountJunctionBoxes(circuits) == len(junctionBoxes) {
			return part1Solution, int64(jb1.X) * int64(jb2.X)
		}
	}
	helper.ExitWithMessage("no connection found!")
	return 0, 0
}

type JBDist struct {
	P1, P2   int
	Distance float64
}

func PreparePointPairs(junctionBoxes []JunctionBox) []JBDist {
	pointPairs := make([]JBDist, 0)
	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			pointPairs = append(pointPairs, JBDist{
				P1:       i,
				P2:       j,
				Distance: junctionBoxes[i].DistanceTo(junctionBoxes[j]),
			})
		}
	}
	sort.Slice(pointPairs, func(i, j int) bool {
		return pointPairs[i].Distance < pointPairs[j].Distance
	})
	return pointPairs
}

func GetSolution1(circuits []Circuit) int64 {
	lens := make([]int, 0)
	for _, c := range circuits {
		// skip deleted circuits
		if len(c) > 0 {
			lens = append(lens, len(c))
		}
	}
	sort.Ints(lens)

	val := int64(lens[len(lens)-1]) * int64(lens[len(lens)-2]) * int64(lens[len(lens)-3])
	return val
}
