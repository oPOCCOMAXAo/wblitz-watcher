package diff

const Ints ints = 0

type ints byte

func (ints) GetUniqueID(i int64) int64 {
	return i
}

func (ints) PrepareToUpdate(a, b int64) bool {
	return a == b
}
