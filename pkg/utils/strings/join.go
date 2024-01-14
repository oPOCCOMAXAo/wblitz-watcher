package strings

func CopyJoin(
	value string,
	separator string,
	copiesCount int,
) string {
	if copiesCount == 0 {
		return ""
	}

	res := value
	for i := 1; i < copiesCount; i++ {
		res += separator + value
	}

	return res
}
