package dijkstra

import (
	"aoc/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMustFindPath(t *testing.T) {
	maze := [][]rune{
		[]rune("#######"),
		[]rune("#     #"),
		[]rune("# # # #"),
		[]rune("# #   #"),
		[]rune("# ## ##"),
		[]rune("#  #  #"),
		[]rune("#######"),
	}
	path, dist := MustFindPath(helper.Vec2D[int]{X: 1, Y: 1}, helper.Vec2D[int]{X: 5, Y: 5}, Params[int, helper.Vec2D[int]]{
		SuccessorGenerator: NewRuneFieldSuccessorGenerator(maze, []rune{'#'}),
	})

	require.Equal(t, []helper.Vec2D[int]{{X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1}, {X: 3, Y: 2}, {X: 3, Y: 3}, {X: 4, Y: 3}, {X: 4, Y: 4}, {X: 4, Y: 5}, {X: 5, Y: 5}}, path)
	require.Equal(t, 8, dist)
}

func TestMustFindPathNoDiagonal(t *testing.T) {
	maze := [][]rune{
		[]rune("####"),
		[]rune("#  #"),
		[]rune("#  #"),
		[]rune("####"),
	}
	path, dist := MustFindPath(helper.Vec2D[int]{X: 1, Y: 1}, helper.Vec2D[int]{X: 2, Y: 2}, Params[int, helper.Vec2D[int]]{
		SuccessorGenerator: NewRuneFieldSuccessorGenerator(maze, []rune{'#'}),
	})

	require.Len(t, path, 3)
	require.Equal(t, helper.Vec2D[int]{X: 1, Y: 1}, path[0])
	require.Contains(t, []helper.Vec2D[int]{{X: 2, Y: 1}, {X: 1, Y: 2}}, path[1])
	require.Equal(t, helper.Vec2D[int]{X: 2, Y: 2}, path[2])
	require.Equal(t, 2, dist)
}

func TestMustFindPathOneField(t *testing.T) {
	maze := [][]rune{
		[]rune("###"),
		[]rune("# #"),
		[]rune("###"),
	}
	path, dist := MustFindPath(helper.Vec2D[int]{X: 1, Y: 1}, helper.Vec2D[int]{X: 1, Y: 1}, Params[int, helper.Vec2D[int]]{
		SuccessorGenerator: NewRuneFieldSuccessorGenerator(maze, []rune{'#'}),
	})

	require.Equal(t, []helper.Vec2D[int]{{X: 1, Y: 1}}, path)
	require.Equal(t, 0, dist)
}

func TestFindPathNoPath(t *testing.T) {
	maze := [][]rune{
		[]rune("#######"),
		[]rune("#     #"),
		[]rune("# # # #"),
		[]rune("# #   #"),
		[]rune("# #####"),
		[]rune("#  #  #"),
		[]rune("#######"),
	}
	_, _, ok := FindPath(helper.Vec2D[int]{X: 1, Y: 1}, helper.Vec2D[int]{X: 5, Y: 5}, Params[int, helper.Vec2D[int]]{
		SuccessorGenerator: NewRuneFieldSuccessorGenerator(maze, []rune{'#'}),
	})
	require.False(t, ok)
}

func TestFindPathMissingSuccessorGenerator(t *testing.T) {
	require.Panics(t, func() {
		FindPath(helper.Vec2D[int]{X: 0, Y: 0}, helper.Vec2D[int]{X: 0, Y: 0}, Params[int, helper.Vec2D[int]]{
			SuccessorGenerator: nil,
		})
	})
}
