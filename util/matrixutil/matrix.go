package matrixutil

type Position struct {
	X int
	Y int
}

type Matrix[T any] struct {
	Matrix          [][]T
	CurrentPosition Position
	Size            int
}

// matrix.Matrix = make([][]string, length)
// var matrix Matrix[string]

func (mtx Matrix[T]) Move(x int, y int) (successful bool) {
	if x < 0 || y < 0 || x >= mtx.Size || y >= mtx.Size {
		return false
	}

	mtx.CurrentPosition.X = x
	mtx.CurrentPosition.Y = y
	return true
}
