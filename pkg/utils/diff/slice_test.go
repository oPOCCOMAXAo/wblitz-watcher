package diff

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculate_int64(t *testing.T) {
	testCases := []struct {
		desc    string
		newList []int64
		oldList []int64
		diff    Diff[int64]
	}{
		{
			desc:    "same lists",
			newList: []int64{1, 2, 3},
			oldList: []int64{1, 2, 3},
			diff: Diff[int64]{
				Created: []int64{},
				Updated: []int64{},
				Deleted: []int64{},
			},
		},
		{
			desc:    "same lists with different order",
			newList: []int64{1, 2, 3},
			oldList: []int64{3, 2, 1},
			diff: Diff[int64]{
				Created: []int64{},
				Updated: []int64{},
				Deleted: []int64{},
			},
		},
		{
			desc:    "default",
			newList: []int64{1, 2, 3},
			oldList: []int64{1, 2, 4},
			diff: Diff[int64]{
				Created: []int64{3},
				Updated: []int64{},
				Deleted: []int64{4},
			},
		},
		{
			desc:    "delete only",
			newList: []int64{1, 2, 3},
			oldList: []int64{1, 2, 3, 4},
			diff: Diff[int64]{
				Created: []int64{},
				Updated: []int64{},
				Deleted: []int64{4},
			},
		},
		{
			desc:    "create only",
			newList: []int64{1, 2, 3, 4},
			oldList: []int64{1, 2, 3},
			diff: Diff[int64]{
				Created: []int64{4},
				Updated: []int64{},
				Deleted: []int64{},
			},
		},
		{
			desc:    "new empty list",
			newList: []int64{},
			oldList: []int64{1, 2, 3},
			diff: Diff[int64]{
				Created: []int64{},
				Updated: []int64{},
				Deleted: []int64{1, 2, 3},
			},
		},
		{
			desc:    "old empty list",
			newList: []int64{1, 2, 3},
			oldList: []int64{},
			diff: Diff[int64]{
				Created: []int64{1, 2, 3},
				Updated: []int64{},
				Deleted: []int64{},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res := Slice(
				tC.newList,
				tC.oldList,
				Ints.GetUniqueID,
				Ints.PrepareToUpdate,
				Ints.Equals,
			)

			require.Equal(t, tC.diff, res)
		})
	}
}

func TestCalculate_struct(t *testing.T) {
	type element struct {
		UniqueID   int64
		InternalID int64
		Value      string
	}

	type list []*element

	testCases := []struct {
		newList list
		oldList list
		diff    Diff[*element]
	}{
		{
			newList: list{
				{UniqueID: 1, Value: "a2"},
				{UniqueID: 2, Value: "b"},
				{UniqueID: 4, Value: "c"},
			},
			oldList: list{
				{UniqueID: 1, InternalID: 1, Value: "a"},
				{UniqueID: 2, InternalID: 2, Value: "b"},
				{UniqueID: 3, InternalID: 3, Value: "c"},
			},
			diff: Diff[*element]{
				Created: []*element{
					{UniqueID: 4, Value: "c"},
				},
				Updated: []*element{
					{UniqueID: 1, InternalID: 1, Value: "a2"},
				},
				Deleted: []*element{
					{UniqueID: 3, InternalID: 3, Value: "c"},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run("", func(t *testing.T) {
			res := Slice(
				tC.newList,
				tC.oldList,
				func(e *element) int64 { return e.UniqueID },
				func(newE, oldE *element) { newE.InternalID = oldE.InternalID },
				func(a, b *element) bool { return *a == *b },
			)

			require.Equal(t, tC.diff, res)
		})
	}
}
