package maputil

// note: compiler can infer types so generally you dont have to specifi them during usage
func GetKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
