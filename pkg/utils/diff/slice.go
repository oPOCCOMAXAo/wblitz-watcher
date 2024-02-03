package diff

// Slice calculates diff between two slices.
//
// It returns a Diff struct with created, updated and deleted elements.
func Slice[T ~[]E, E any, ID comparable](
	newList T,
	oldList T,
	getUniqueIDFn func(e E) ID,
	prepareToUpdateFn func(newE, oldE E),
	equals func(newE, oldE E) bool,
) Diff[E] {
	res := Diff[E]{
		Created: make([]E, 0, len(newList)),
		Updated: make([]E, 0, len(newList)),
		Deleted: make([]E, 0, len(oldList)),
	}

	oldMap := make(map[ID]E, len(oldList))
	newMap := make(map[ID]E, len(newList))

	for _, oldItem := range oldList {
		oldMap[getUniqueIDFn(oldItem)] = oldItem
	}

	for _, newItem := range newList {
		id := getUniqueIDFn(newItem)

		newMap[id] = newItem

		oldItem, ok := oldMap[id]
		if !ok {
			res.Created = append(res.Created, newItem)

			continue
		}

		prepareToUpdateFn(newItem, oldItem)

		if !equals(newItem, oldItem) {
			res.Updated = append(res.Updated, newItem)
		}
	}

	// we should preserve order of oldList.
	for _, oldItem := range oldList {
		_, ok := newMap[getUniqueIDFn(oldItem)]
		if ok {
			continue
		}

		res.Deleted = append(res.Deleted, oldItem)
	}

	return res
}
