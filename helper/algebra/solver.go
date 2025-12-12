package algebra

import "aoc/helper"

type Mat[T helper.Number] [][]T

func NewMat[T helper.Number](rows, cols int) Mat[T] {
	mat := make(Mat[T], rows)
	for r := range rows {
		mat[r] = make([]T, cols)
	}
	return mat
}

func (m *Mat[T]) Rows() int {
	return len(*m)
}

func (m *Mat[T]) Cols() int {
	if len(*m) == 0 {
		return 0
	}
	return len((*m)[0])
}

func (m *Mat[T]) Set(row, col int, value T) {
	(*m)[row][col] = value
}

func (m *Mat[T]) At(row, col int) T {
	return (*m)[row][col]
}

func (m *Mat[T]) SwapRows(row1, row2 int) {
	for col := 0; col < m.Cols(); col++ {
		(*m)[row1][col], (*m)[row2][col] = (*m)[row2][col], (*m)[row1][col]
	}
}

type Vec[T helper.Number] []T

func NewVecFromSlice[T helper.Number](arr []T) Vec[T] {
	return Vec[T](arr)
}

func MustSolveInteger[T helper.Number, I helper.Integer](mat Mat[T], vec Vec[T]) Vec[I] {
	result, ok := SolveInteger[T, I](mat, vec)
	if !ok {
		helper.ExitWithMessage("no solution found")
	}
	return result
}

func SolveInteger[T helper.Number, I helper.Integer](mat Mat[T], vec Vec[T]) (Vec[I], bool) {
	return nil, false
}
