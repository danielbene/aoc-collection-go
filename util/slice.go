package util

// faster unordered solution: https://stackoverflow.com/a/37335777
func RemoveSliceElement(slice []string, idx int) []string {
	return append(slice[:idx], slice[idx+1:]...)
}
