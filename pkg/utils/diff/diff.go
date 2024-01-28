package diff

type Diff[E any] struct {
	Created []E
	Updated []E
	Deleted []E
}

func (d *Diff[E]) IsEmpty() bool {
	return len(d.Created) == 0 &&
		len(d.Updated) == 0 &&
		len(d.Deleted) == 0
}
