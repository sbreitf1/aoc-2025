package helper

import (
	"fmt"
	"strings"
)

func Dirs4() []Vec2D[int] {
	return []Vec2D[int]{{X: 0, Y: -1}, {X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}}
}

func Dirs8() []Vec2D[int] {
	return []Vec2D[int]{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1}, {X: -1, Y: 0}, {X: 1, Y: 0}, {X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1}}
}

type Field[T comparable] struct {
	Field [][]T
}

func NewRuneFieldFromLines(lines []string) Field[rune] {
	return Field[rune]{
		Field: LinesToRunes(lines),
	}
}

func NewRuneField(width, height int, empty rune) Field[rune] {
	lines := make([]string, height)
	for y := range lines {
		lines[y] = strings.Repeat(string(empty), width)
	}
	return NewRuneFieldFromLines(lines)
}

func (f Field[T]) Print() {
	for y := 0; y < f.Height(); y++ {
		for x := 0; x < f.Width(); x++ {
			switch val := any(f.Field[y][x]).(type) {
			case rune:
				fmt.Print(string(val))
			}
		}
		fmt.Println()
	}
}

func (f Field[T]) Width() int {
	if len(f.Field) == 0 {
		return 0
	}
	return len(f.Field[0])
}

func (f Field[T]) Height() int {
	return len(f.Field)
}

func (f Field[T]) MustFindOnce(val T) Vec2D[int] {
	found := false
	var pos Vec2D[int]
	for y := 0; y < f.Height(); y++ {
		for x := 0; x < f.Width(); x++ {
			if f.Field[y][x] == val {
				if found {
					ExitWithMessage("value %v found twice in field", val)
				}

				pos = Vec2D[int]{X: x, Y: y}
				found = true
			}
		}
	}
	if !found {
		ExitWithMessage("value %v not found in field", val)
	}
	return pos
}

func (f Field[T]) At(p Vec2D[int]) T {
	return f.Field[p.Y][p.X]
}

func (f Field[T]) AtXY(x, y int) T {
	return f.Field[y][x]
}

func (f Field[T]) Set(p Vec2D[int], val T) {
	f.Field[p.Y][p.X] = val
}

func (f Field[T]) SetXY(x, y int, val T) {
	f.Field[y][x] = val
}

func (f Field[T]) InBounds(p Vec2D[int]) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < f.Width() && p.Y < f.Height()
}

func (f Field[T]) SurroundingFields4(p Vec2D[int]) []Vec2D[int] {
	neighbours := make([]Vec2D[int], 0, 4)
	for _, d := range Dirs4() {
		sp := p.Add(d)
		if f.InBounds(sp) {
			neighbours = append(neighbours, sp)
		}
	}
	return neighbours
}

func (f Field[T]) SurroundingFields8(p Vec2D[int]) []Vec2D[int] {
	neighbours := make([]Vec2D[int], 0, 8)
	for _, d := range Dirs8() {
		sp := p.Add(d)
		if f.InBounds(sp) {
			neighbours = append(neighbours, sp)
		}
	}
	return neighbours
}

func (f Field[T]) Iterate(itFunc func(p Vec2D[int])) {
	for y := 0; y < f.Height(); y++ {
		for x := 0; x < f.Width(); x++ {
			itFunc(Vec2D[int]{X: x, Y: y})
		}
	}
}

func (f Field[T]) IterateNeighbours4(p Vec2D[int], itFunc func(p Vec2D[int])) {
	for _, np := range f.SurroundingFields4(p) {
		itFunc(np)
	}
}

func (f Field[T]) IterateNeighbours8(p Vec2D[int], itFunc func(p Vec2D[int])) {
	for _, np := range f.SurroundingFields8(p) {
		itFunc(np)
	}
}
