package du

func GetFromAnyMap[R any](m map[string]any, key string) R {
	var res R

	v, ok := m[key]
	if !ok {
		return res
	}

	res, _ = v.(R)

	return res
}
