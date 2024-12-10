package sliceutil

// ----------------------------- MODIFY -----------------------------

// faster unordered solution: https://stackoverflow.com/a/37335777
func RemoveStringSliceElement[T any](slice []T, idx int) []T {
	return append(slice[:idx], slice[idx+1:]...)
}

func SwapSliceElements[T ~[]int](slice T, a int, b int) T {
	slice[a], slice[b] = slice[b], slice[a]
	return slice
}

// ----------------------------- COMPARE ----------------------------

func CountCommonSliceElements[T comparable](left []T, right []T) int {
	return len(GetCommonSliceElements(left, right))
}

func GetCommonSliceElements[T comparable](left []T, right []T) (commonElements []T) {
	for _, a := range left {
		for _, b := range right {
			if a == b {
				commonElements = append(commonElements, b)
			}
		}
	}

	return commonElements
}

// slices.Index is the same as this but generic!
/*func IntContainsElement(slice []int, target int) (contains bool, idx int) {
	for i, val := range slice {
		if val == target {
			return true, i
		}
	}

	return false, 0
}*/
