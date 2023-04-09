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

func TestFilterIterable_Cycle(t *testing.T) {
	arr := []int{1, 2, 3}
	itr := filterIterable[int]{
		New(arr),
		func(v int) bool {
			return v > 1
		},
		nil,
	}
	cyc := itr.Cycle()
	var res []int
	for i := 0; i <= 4; i++ {
		res = append(res, cyc.Next())
	}
	require.Equal(t, []int{2, 3, 2, 3, 2}, res)
}

func TestFilterIterable_Sort(t *testing.T) {
	arr := []int{3, 2, 1}
	itr := filterIterable[int]{
		New(arr),
		func(v int) bool {
			return v > 1
		},
		nil,
	}
	res := itr.Sort(func(a int, b int) bool {
		return a < b
	}).ToSlice()
	require.Equal(t, []int{2, 3}, res)
}
