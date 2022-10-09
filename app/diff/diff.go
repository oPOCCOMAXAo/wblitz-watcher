package diff

import (
	"fmt"
	"strings"

	"github.com/opoccomaxao-go/generic-collection/set"
	"github.com/opoccomaxao-go/generic-collection/slice"
)

type Type int

type Void struct{}

type Diff[T any] struct {
	Type Type
	Old  T
	New  T
}

func (d *Diff[T]) Pretty() string {
	return fmt.Sprintf("%d: %v -> %v", d.Type, d.Old, d.New)
}

func (Diff[T]) PrettyStringer(diff Diff[T]) string {
	return diff.Pretty()
}

type Total struct {
	Int    []Diff[int]
	String []Diff[string]
	Void   []Diff[Void]
}

func (total *Total) Len() int {
	return len(total.Int) + len(total.String) + len(total.Void)
}

func (total *Total) Pretty() string {
	return strings.Join([]string{
		slice.Join(total.Void, "\n", Diff[Void]{}.PrettyStringer),
		slice.Join(total.Int, "\n", Diff[int]{}.PrettyStringer),
		slice.Join(total.String, "\n", Diff[string]{}.PrettyStringer),
	}, "\n")
}

func DetectSingleValue[T comparable](
	typ Type,
	oldValue, newValue T,
	resArr *[]Diff[T],
) {
	if oldValue != newValue {
		*resArr = append(*resArr, Diff[T]{
			Type: typ,
			Old:  oldValue,
			New:  newValue,
		})
	}
}

func DetectSetNew[T comparable](
	typ Type,
	oldList, newList []T,
	resArr *[]Diff[T],
) {
	added := DiffNewOnly(set.FromSlice(oldList), newList)

	for _, value := range added {
		*resArr = append(*resArr, Diff[T]{
			Type: typ,
			Old:  value,
			New:  value,
		})
	}
}

func DiffNewOnly[T comparable](oldSet set.Set[T], newList []T) []T {
	res := []T{}

	for _, newKey := range newList {
		if !oldSet.Has(newKey) {
			res = append(res, newKey)
		}
	}

	return res
}
