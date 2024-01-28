package diff

// UniqueIDFunc is a function that returns a unique ID for element.
type UniqueIDFunc[E any, ID comparable] func(e E) ID

// PrepareToUpdateFunc is a function that compares two elements and returns true if element should be updated.
// Also it must copy all required information from old element to new element.
type PrepareToUpdateFunc[E any] func(old, new E) bool

// Calculate calculates diff between two slices.
//
// It returns a Diff struct with created, updated and deleted elements.
//
// Params:
//   - getUniqueIDFn see UniqueIDFunc.
//   - prepareToUpdateFn see PrepareToUpdateFunc.
func Calculate[T ~[]E, E any, ID comparable](
	newList T,
	oldList T,
	getUniqueIDFn UniqueIDFunc[E, ID],
	prepareToUpdateFn PrepareToUpdateFunc[E],
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

		if prepareToUpdateFn(oldItem, newItem) {
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
