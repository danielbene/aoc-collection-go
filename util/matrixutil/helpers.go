package matrixutil

func DeepCopy(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for idx := range src {
		dst[idx] = make([]string, len(src[idx]))
		copy(dst[idx], src[idx])
	}

	return dst
}
