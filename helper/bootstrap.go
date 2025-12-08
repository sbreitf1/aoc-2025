package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func init() {
	if dir, puzzleNum, ok := getActivePuzzleDir(); ok {
		lines := ReadLines(filepath.Join(dir, "main.go"))

		var hasLinkHeader, hasInputFile bool
		for _, l := range lines {
			if strings.HasPrefix(l, "// https://adventofcode.com/2025/day") {
				hasLinkHeader = true
				if l != fmt.Sprintf("// https://adventofcode.com/2025/day/%d", puzzleNum) {
					fmt.Println("puzzle", puzzleNum, "has invalid link header", l)
				}
			}

			if strings.Contains(l, "input.txt") {
				hasInputFile = true
			}
		}

		if !hasLinkHeader {
			fmt.Println("puzzle", puzzleNum, "is missing link header", fmt.Sprintf("// https://adventofcode.com/2025/day/%d", puzzleNum))
		}
		if !hasInputFile {
			fmt.Println("WARN: puzzle", puzzleNum, "is probably processing example data")
		}
	}
}

func getActivePuzzleDir() (string, int, bool) {
	rootDir, ok := getRootDir()
	if !ok {
		return "", 0, false
	}

	if len(os.Args) == 0 {
		return "", 0, false
	}

	pattern := regexp.MustCompile(`puzzle-(\d+)`)
	m := pattern.FindStringSubmatch(os.Args[0])
	if len(m) != 2 {
		return "", 0, false
	}

	puzzleNum := ParseInt[int](m[1])
	return filepath.Join(rootDir, fmt.Sprintf("puzzle-%d", puzzleNum)), puzzleNum, true
}

func getRootDir() (string, bool) {
	dir, err := os.Getwd()
	if err != nil {
		return "", false
	}

	for i := 0; i < 3; i++ {
		_, err := os.Stat(filepath.Join(dir, "go.mod"))
		if err == nil {
			return dir, true
		}

		dir = filepath.Dir(dir)
	}
	return "", false
}
