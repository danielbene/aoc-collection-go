package matrixutil

func RemoveMatrixRow[T any](slice [][]T, idx int) [][]T {
	return append(slice[:idx], slice[idx+1:]...)
}

func DeepCopy(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for idx := range src {
		dst[idx] = make([]string, len(src[idx]))
		copy(dst[idx], src[idx])
	}

	return dst
}
