package diff

const Ints Comparable[int64] = 0

const Strings Comparable[string] = 0

type Comparable[T comparable] byte

func (Comparable[T]) GetUniqueID(a T) T {
	return a
}

func (Comparable[T]) PrepareToUpdate(_, _ T) {}

func (Comparable[T]) Equals(a, b T) bool {
	return a == b
}
