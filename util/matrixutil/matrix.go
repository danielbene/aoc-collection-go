package matrixutil

import (
	"aoc/util/matrixutil/directions"
	"fmt"
	"strconv"
)

type Position struct {
	X int // Column (element index in current row)
	Y int // Row (line index in the matrix)
}

type Matrix[T any] struct {
	Matrix          [][]T    // Matrix[Y][X] - 2D slice representation of a matrix with generic type
	CurrentPosition Position // Cursor for seeking through the matrix
	RowCount        int      // Total number of rows in the matrix
	ColCount        int      // Total number of columns in a matrix row
}

func (mtx Matrix[T]) Print() {
	for _, row := range mtx.Matrix {
		fmt.Println(row)
	}
}

// initialize a new matrix object based on the input string
func Init[T ~string | ~int](inputString []string) Matrix[T] {
	var mtx Matrix[T]

	mtx.RowCount = len(inputString)
	mtx.ColCount = len(inputString[0])
	mtx.Matrix = make([][]T, mtx.RowCount)

	for y, line := range inputString {
		mtx.Matrix[y] = make([]T, mtx.ColCount)

		for x, ch := range line {
			switch any(mtx.Matrix[y][x]).(type) {
			case string:
				mtx.Matrix[y][x] = any(string(ch)).(T)
			case int:
				num, _ := strconv.Atoi(string(ch))
				mtx.Matrix[y][x] = any(int(num)).(T) // TODO: test
			}
		}
	}

	return mtx
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
