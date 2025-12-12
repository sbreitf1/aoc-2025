package main

// https://adventofcode.com/2025/day/10

import (
	"aoc/helper"
	"aoc/helper/dijkstra"
	"fmt"
	"regexp"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	machines := ParseMachines(lines)

	solution1 := SumMinButtonPresses1(machines)
	fmt.Println("-> part 1:", solution1)

	solution2 := SumMinButtonPresses2(machines)
	fmt.Println("-> part 2:", solution2)
}

type Machine struct {
	DstLights  Lights
	DstJoltage Joltage
	Buttons    [][]int
}

func (m *Machine) StartLights() Lights {
	return Lights{states: 0, count: m.DstLights.count}
}

func (m *Machine) StartJoltage() Joltage {
	return Joltage{count: m.DstJoltage.count}
}

type Lights struct {
	states uint64
	count  int
}

func ParseLights(str string) Lights {
	var states uint64
	for i, r := range str {
		if r == '#' {
			states += helper.Pow(2, uint64(i))
		}
	}
	return Lights{
		states: states,
		count:  len(str),
	}
}

func (l Lights) String() string {
	var str string
	for i := 0; i < l.count; i++ {
		v := helper.Pow(2, uint64(i))
		if l.states&v == v {
			str += "#"
		} else {
			str += "."
		}
	}
	return str
}

func (l Lights) Toggle(indices []int) Lights {
	newStates := l.states
	for _, idx := range indices {
		v := helper.Pow(2, uint64(idx))
		if newStates&v == v {
			newStates -= v
		} else {
			newStates += v
		}
	}
	return Lights{
		states: newStates,
		count:  l.count,
	}
}

func ParseMachines(lines []string) []Machine {
	patternMachine := regexp.MustCompile(`\[([.#]+)\]\s*(\([^{]*\))\s*{([0-9,]*)}`)
	patternButton := regexp.MustCompile(`\(([0-9,]+)\)`)
	machines := make([]Machine, 0, len(lines))
	for _, l := range lines {
		m := patternMachine.FindStringSubmatch(l)
		if len(m) == 4 {
			dstLights := ParseLights(m[1])
			mb := patternButton.FindAllStringSubmatch(m[2], -1)
			buttons := make([][]int, 0, len(mb))
			for _, m := range mb {
				buttons = append(buttons, helper.ExtractInts[int](m[1]))
			}
			dstJoltage := ParseJoltage(m[3])
			machines = append(machines, Machine{
				DstLights:  dstLights,
				Buttons:    buttons,
				DstJoltage: dstJoltage,
			})
		}
	}
	return machines
}

func (m *Machine) FindMinButtonPresses1() int {
	_, pressCount := dijkstra.MustFindPath(m.StartLights(), m.DstLights, dijkstra.Params[int, Lights]{
		SuccessorGenerator: func(current Lights, currentDist int) []dijkstra.Successor[int, Lights] {
			successors := make([]dijkstra.Successor[int, Lights], 0, len(m.Buttons))
			for _, b := range m.Buttons {
				sl := current.Toggle(b)
				successors = append(successors, dijkstra.Successor[int, Lights]{
					Obj:  sl,
					Dist: currentDist + 1,
				})
			}
			return successors
		},
	})
	return pressCount
}

func SumMinButtonPresses1(machines []Machine) int {
	var sum int
	for _, m := range machines {
		sum += m.FindMinButtonPresses1()
	}
	return sum
}

type Joltage struct {
	values [64]int
	count  int
}

func ParseJoltage(str string) Joltage {
	ints := helper.ExtractInts[int](str)
	var values [64]int
	copy(values[:len(ints)], ints)
	return Joltage{
		values: values,
		count:  len(ints),
	}
}

func (j Joltage) String() string {
	str := "{"
	for i := 0; i < j.count; i++ {
		if i > 0 {
			str += ","
		}
		str += fmt.Sprintf("%d", j.values[i])
	}
	return str + "}"
}

func (j Joltage) Slice() []int {
	values := make([]int, j.count)
	copy(values, j.values[:j.count])
	return values
}

func (j Joltage) Increase(indices []int, delta int) Joltage {
	var newValues [64]int
	copy(newValues[:], j.values[:])
	for _, idx := range indices {
		newValues[idx] += delta
	}
	return Joltage{
		values: newValues,
		count:  j.count,
	}
}

func (j Joltage) Exceeds(other Joltage) bool {
	if len(j.values) != len(other.values) {
		panic("incompatible joltages")
	}

	for i := range j.values {
		if j.values[i] > other.values[i] {
			return true
		}
	}
	return false
}

func (m *Machine) findMinButtonPresses2(currentJoltage Joltage, btnIndex int) (int, bool) {
	if currentJoltage == m.DstJoltage {
		return 0, true
	}
	if btnIndex >= len(m.Buttons) {
		return 0, false
	}

	var btnPresses int

	bestResult := 100000000
	hasResult := false

	for {
		followingPresses, ok := m.findMinButtonPresses2(currentJoltage, btnIndex+1)
		if ok && (btnPresses+followingPresses) < bestResult {
			bestResult = btnPresses + followingPresses
			hasResult = true
		}

		currentJoltage = currentJoltage.Increase(m.Buttons[btnIndex], 1)
		if currentJoltage.Exceeds(m.DstJoltage) {
			break
		}
		btnPresses++
	}

	return bestResult, hasResult
}

func (m *Machine) FindMinButtonPresses2() int {
	result, ok := m.findMinButtonPresses2(m.StartJoltage(), 0)
	if !ok {
		helper.ExitWithMessage("no solution for part 2 found")
	}
	return result
}

func SumMinButtonPresses2(machines []Machine) int {
	var sum int
	for _, m := range machines {
		sum += m.FindMinButtonPresses2()
		fmt.Println(sum)
	}
	return sum
}
