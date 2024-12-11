package matrixutil

import "aoc/util/matrixutil/directions"

type Position struct {
	X int // Column (element index in current row)
	Y int // Row (line index in the matrix)
}

type Matrix[T any] struct {
	Matrix          [][]T    // 2D slice representation of a matrix with generic type
	CurrentPosition Position // Cursor for seeking through the matrix
	RowCount        int      // Total number of rows in the matrix
	ColCount        int      // Total number of columns in a matrix row
}

func (mtx Matrix[T]) GetValueDirection(dir directions.Direction) (value T, successful bool) {
	return mtx.GetValue(mtx.CurrentPosition.X+dir.X, mtx.CurrentPosition.Y+dir.Y)
}

func (mtx Matrix[T]) GetValue(x int, y int) (value T, successful bool) {
	if mtx.isInsideBoundaries(x, y) {
		var defValue T
		return defValue, false
	}

	return mtx.Matrix[y][x], true
}

func (mtx Matrix[T]) MoveDirection(dir directions.Direction) (successful bool) {
	return mtx.Move(mtx.CurrentPosition.X+dir.X, mtx.CurrentPosition.Y+dir.Y)
}

func (mtx Matrix[T]) Move(x int, y int) (successful bool) {
	if mtx.isInsideBoundaries(x, y) {
		return false
	}

	mtx.CurrentPosition.X = x
	mtx.CurrentPosition.Y = y
	return true
}

func (mtx Matrix[T]) isInsideBoundaries(x int, y int) bool {
	if x < 0 || y < 0 || x >= mtx.ColCount || y >= mtx.RowCount {
		return false
	}

	return true
}
