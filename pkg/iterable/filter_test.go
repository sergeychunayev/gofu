package iterable

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFilterIterable_HasNext(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 3}},
		{[]int{3, 2, 1}, []int{3, 2}},
		{[]int{1, 2, 1}, []int{2}},
		{[]int{1, 2, 1, 1, 1, 1, 2}, []int{2, 2}},
		{[]int{1, 1, 1}, nil},
		{[]int{}, nil},
	} {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			itr := filterIterable[int]{
				New(tc.input),
				func(v int) bool {
					return v > 1
				},
				nil,
			}
			res := itr.ToSlice()
			require.Equal(t, tc.expected, res)
		})
	}
}
