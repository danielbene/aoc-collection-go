package sliceutil

// ----------------------------- MODIFY -----------------------------

// faster unordered solution: https://stackoverflow.com/a/37335777
func RemoveStringSliceElement(slice []string, idx int) []string {
	return append(slice[:idx], slice[idx+1:]...)
}

func RemoveIntSliceElement(slice []int, idx int) []int {
	return append(slice[:idx], slice[idx+1:]...)
}

func SwapSliceElements(slice []int, a int, b int) []int {
	slice[a], slice[b] = slice[b], slice[a]
	return slice
}

// ----------------------------- COMPARE ----------------------------

func CountCommonSliceElements(left []string, right []string) int {
	return len(GetCommonSliceElements(left, right))
}

func GetCommonSliceElements(left []string, right []string) (commonElements []string) {
	for _, a := range left {
		for _, b := range right {
			if a == b {
				commonElements = append(commonElements, b)
			}
		}
	}

	return commonElements
}

func IntContainsElement(slice []int, target int) (contains bool, idx int) {
	for i, val := range slice {
		if val == target {
			return true, i
		}
	}

	return false, 0
}
