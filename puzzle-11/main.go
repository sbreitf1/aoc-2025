package main

// https://adventofcode.com/2025/day/11

import (
	"aoc/helper"
	"fmt"
	"strings"
)

func main() {
	lines := helper.ReadNonEmptyLines("input.txt")

	graph := ParseGraph(lines)

	solution1 := CountPathsTo1(graph, "you", "out")
	fmt.Println("-> part 1:", solution1)

	solution2 := CountPathsTo2(graph, "svr", "out", false, false)
	fmt.Println("-> part 2:", solution2)
}

type Graph struct {
	Nodes           map[string][]string
	pathCountCache2 map[crumble]int
}

func ParseGraph(lines []string) Graph {
	nodes := make(map[string][]string)
	for _, l := range lines {
		parts := strings.Split(l, ":")
		nodeName := strings.TrimSpace(parts[0])
		for _, next := range strings.Split(parts[1], " ") {
			next = strings.TrimSpace(next)
			if len(next) > 0 {
				nodes[nodeName] = append(nodes[nodeName], strings.TrimSpace(next))
			}
		}
	}
	return Graph{
		Nodes:           nodes,
		pathCountCache2: make(map[crumble]int),
	}
}

func CountPathsTo1(graph Graph, current, dst string) int {
	if current == dst {
		return 1
	}

	var pathCount int
	for _, next := range graph.Nodes[current] {
		pathCount += CountPathsTo1(graph, next, dst)
	}
	return pathCount
}

type crumble struct {
	name       string
	dacVisited bool
	fftVisited bool
}

func CountPathsTo2(graph Graph, current, dst string, dacVisited, fftVisited bool) int {
	if dacVisited && fftVisited && current == dst {
		return 1
	}
	if pathCount, ok := graph.pathCountCache2[crumble{current, dacVisited, fftVisited}]; ok {
		return pathCount
	}

	if current == "dac" {
		dacVisited = true
	}
	if current == "fft" {
		fftVisited = true
	}

	var pathCount int
	for _, next := range graph.Nodes[current] {
		pathCount += CountPathsTo2(graph, next, dst, dacVisited, fftVisited)
	}
	graph.pathCountCache2[crumble{current, dacVisited, fftVisited}] = pathCount
	return pathCount
}
