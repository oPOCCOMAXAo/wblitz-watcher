package utils

func GetFromMap[K comparable, V any, M ~map[K]V](m M, key, defaultKey K) V {
	v, ok := m[key]
	if ok {
		return v
	}

	return m[defaultKey]
}
