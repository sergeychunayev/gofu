package ord_test

import (
	"github.com/sergeychunayev/gofu/pkg/iterable"
	"github.com/sergeychunayev/gofu/pkg/iterable/ord"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMinFunc(t *testing.T) {
	require.Equal(t, 1, ord.Min(1, 2))
	require.Equal(t, 1, ord.Min(2, 1))
	require.Equal(t, 1, ord.Min(1, 1))
}

func TestMaxFunc(t *testing.T) {
	require.Equal(t, 2, ord.Max(1, 2))
	require.Equal(t, 2, ord.Max(2, 1))
	require.Equal(t, 2, ord.Max(2, 2))
}

func TestMinBy(t *testing.T) {
	type S struct {
		name  string
		value int
	}

	testCases := []struct {
		name     string
		input    []S
		expected iterable.Tuple[S, bool]
	}{
		{
			"Empty",
			[]S{},
			iterable.Tuple[S, bool]{},
		},
		{
			"Increasing",
			[]S{
				{"one", 1},
				{"two", 2},
				{"three", 3},
			},
			iterable.Tuple[S, bool]{A: S{"one", 1}, B: true},
		},
		{
			"Decreasing",
			[]S{
				{"three", 3},
				{"two", 2},
				{"one", 1},
			},
			iterable.Tuple[S, bool]{A: S{"one", 1}, B: true},
		},
		{
			"All same",
			[]S{
				{"one", 1},
				{"two", 1},
				{"three", 1},
			},
			iterable.Tuple[S, bool]{A: S{"one", 1}, B: true},
		},
		{
			"One",
			[]S{{"one", 1}},
			iterable.Tuple[S, bool]{A: S{"one", 1}, B: true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res, ok := ord.MinBy(itr, func(v S) int {
				return v.value
			})
			tup := iterable.Tuple[S, bool]{A: res, B: ok}
			require.Equal(t, tc.expected, tup)
		})
	}
}

func TestMaxBy(t *testing.T) {
	type S struct {
		name  string
		value int
	}

	testCases := []struct {
		name     string
		input    []S
		expected iterable.Tuple[S, bool]
	}{
		{
			"Empty",
			[]S{},
			iterable.Tuple[S, bool]{},
		},
		{
			"Increasing",
			[]S{
				{"one", 1},
				{"two", 2},
				{"three", 3},
			},
			iterable.Tuple[S, bool]{A: S{"three", 3}, B: true},
		},
		{
			"Decreasing",
			[]S{
				{"three", 3},
				{"two", 2},
				{"one", 1},
			},
			iterable.Tuple[S, bool]{A: S{"three", 3}, B: true},
		},
		{
			"All same",
			[]S{
				{"one", 1},
				{"two", 1},
				{"three", 1},
			},
			iterable.Tuple[S, bool]{A: S{"one", 1}, B: true},
		},
		{
			"One",
			[]S{{"one", 1}},
			iterable.Tuple[S, bool]{A: S{"one", 1}, B: true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res, ok := ord.MaxBy(itr, func(v S) int {
				return v.value
			})
			tup := iterable.Tuple[S, bool]{A: res, B: ok}
			require.Equal(t, tc.expected, tup)
		})
	}
}

func TestLt(t *testing.T) {
	require.Equal(t, true, ord.Lt[int](1, 2))
	require.Equal(t, false, ord.Lt[int](2, 1))
}

func TestGt(t *testing.T) {
	require.Equal(t, false, ord.Gt[int](1, 2))
	require.Equal(t, true, ord.Gt[int](2, 1))
}

func TestSlice_Sum(t *testing.T) {
	res, ok := ord.New([]int{1, 2, 3}).SumOrd()
	require.True(t, ok)
	require.Equal(t, 6, res)
}

func TestSlice_Min(t *testing.T) {
	res, ok := ord.New([]int{1, 2, 3}).Min()
	require.True(t, ok)
	require.Equal(t, 1, res)
}

func TestSlice_Max(t *testing.T) {
	res, ok := ord.New([]int{1, 2, 3}).Max()
	require.True(t, ok)
	require.Equal(t, 3, res)
}

func TestSum(t *testing.T) {
	res, ok := ord.Sum([]int{1, 2, 3})
	require.True(t, ok)
	require.Equal(t, 6, res)

	res, ok = ord.Sum([]int{})
	require.False(t, ok)
	require.Equal(t, 0, res)
}
