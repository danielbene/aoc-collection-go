package matrixutil

import "fmt"

var matrix [][]string

func SetDimensions(length int) {
	matrix = make([][]string, length)
	for i := range matrix {
		//matrix[i] =
		fmt.Println(i)
	}
}
