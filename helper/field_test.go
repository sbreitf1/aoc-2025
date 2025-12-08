package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDirs4(t *testing.T) {
	dirs := Dirs4()
	require.Len(t, dirs, 4)
}

func TestDirs8(t *testing.T) {
	dirs := Dirs8()
	require.Len(t, dirs, 8)
}

func TestRuneField(t *testing.T) {
	field := NewRuneFieldFromLines([]string{
		"#######",
		"#     #",
		"# # ###",
		"# # # #",
		"# #   #",
		"# #####",
		"#  #  #",
		"#######",
	})

	t.Run("FieldSize", func(t *testing.T) {
		require.Equal(t, 7, field.Width())
		require.Equal(t, 8, field.Height())
	})
	t.Run("SurroundingFields4", func(t *testing.T) {
		sps := field.SurroundingFields4(Vec2D[int]{X: 2, Y: 4})
		require.Len(t, sps, 4)
	})
	t.Run("SurroundingFields4TopLeft", func(t *testing.T) {
		sps := field.SurroundingFields4(Vec2D[int]{X: 0, Y: 0})
		require.Len(t, sps, 2)
	})
	t.Run("SurroundingFields4TopRight", func(t *testing.T) {
		sps := field.SurroundingFields4(Vec2D[int]{X: field.Width() - 1, Y: 0})
		require.Len(t, sps, 2)
	})
	t.Run("SurroundingFields4BottomLeft", func(t *testing.T) {
		sps := field.SurroundingFields4(Vec2D[int]{X: 0, Y: field.Height() - 1})
		require.Len(t, sps, 2)
	})
	t.Run("SurroundingFields4BottomLeft", func(t *testing.T) {
		sps := field.SurroundingFields4(Vec2D[int]{X: field.Width() - 1, Y: field.Height() - 1})
		require.Len(t, sps, 2)
	})
}
