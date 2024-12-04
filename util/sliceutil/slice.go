package sliceutil

// ----------------------------- MODIFY -----------------------------

// faster unordered solution: https://stackoverflow.com/a/37335777
func RemoveSliceElement(slice []string, idx int) []string {
	return append(slice[:idx], slice[idx+1:]...)
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

	return
}
