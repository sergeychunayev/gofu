package iterable

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMapIterable_Filter(t *testing.T) {
	itr := mapIterable[int, float64]{
		New([]int{1, 2, 3, 4}),
		func(v int) float64 {
			return float64(v)
		},
	}
	res := itr.Filter(func(v float64) bool {
		return v > 2
	}).ToSlice()
	require.Equal(t, []float64{3.0, 4.0}, res)
}

func TestMapIterable_For(t *testing.T) {
	it := mapIterable[int, float64]{
		New([]int{1, 2, 3, 4, 0}),
		func(v int) float64 {
			return float64(v)
		},
	}
	var res []float64
	it.For(func(v float64, i int) {
		res = append(res, v)
	})
	require.Equal(t, []float64{1.0, 2.0, 3.0, 4.0, 0.0}, res)
}

func TestMapIterable_All(t *testing.T) {
	res := (&mapIterable[int, float64]{
		New([]int{1, 2, 3, 4, 0}),
		func(v int) float64 {
			return float64(v)
		},
	}).All(func(v float64) bool {
		return v == 0.0
	})
	require.False(t, res)
}

func TestMapIterable_Any(t *testing.T) {
	res := (&mapIterable[int, float64]{
		New([]int{1, 2, 3, 4, 0}),
		func(v int) float64 {
			return float64(v)
		},
	}).Any(func(v float64) bool {
		return v == 0.0
	})
	require.True(t, res)
}

func TestMapIterable_Min(t *testing.T) {
	res, ok := (&mapIterable[int, float64]{
		New([]int{1, 2, 3, 4, 0}),
		func(v int) float64 {
			return float64(v)
		},
	}).Min(func(a float64, b float64) bool {
		return a < b
	})
	require.True(t, ok)
	require.Equal(t, 0.0, res)
}

func TestMapIterable_Max(t *testing.T) {
	res, ok := (&mapIterable[int, float64]{
		New([]int{1, 2, 3, 4}),
		func(v int) float64 {
			return float64(v)
		},
	}).Max(func(a float64, b float64) bool {
		return a > b
	})
	require.True(t, ok)
	require.Equal(t, 4.0, res)
}

func TestMapIterable_Reduce(t *testing.T) {
	res, ok := (&mapIterable[int, float64]{
		New([]int{1, 2, 3, 4}),
		func(v int) float64 {
			return float64(v)
		},
	}).Reduce(func(acc float64, v float64) float64 {
		return acc + v
	})
	require.True(t, ok)
	require.Equal(t, 10.0, res)
}

func TestMapIterable_Sort(t *testing.T) {
	itr := mapIterable[int, float64]{
		New([]int{4, 3, 2, 1}),
		func(v int) float64 {
			return float64(v)
		},
	}
	res := itr.Sort(func(a float64, b float64) bool {
		return a < b
	}).ToSlice()
	require.Equal(t, []float64{1.0, 2.0, 3.0, 4.0}, res)
}

func TestMapIterable_Sort_Filter(t *testing.T) {
	itr := mapIterable[int, float64]{
		New([]int{4, 3, 2, 1}),
		func(v int) float64 {
			return float64(v)
		},
	}
	res := itr.
		Sort(func(a float64, b float64) bool {
			return a < b
		}).
		Filter(func(v float64) bool {
			return v > 3.0
		}).
		ToSlice()
	require.Equal(t, []float64{4.0}, res)
}

func TestMapIterable_Filter_Sort(t *testing.T) {
	itr := mapIterable[int, float64]{
		New([]int{4, 3, 2, 1, 5}),
		func(v int) float64 {
			return float64(v)
		},
	}
	res := itr.
		Filter(func(v float64) bool {
			return v > 3.0
		}).
		Sort(func(a float64, b float64) bool {
			return a < b
		}).
		ToSlice()
	require.Equal(t, []float64{4.0, 5.0}, res)
}

func TestMapIterable_Cycle(t *testing.T) {
	itr := mapIterable[int, float64]{
		New([]int{1, 2, 3}),
		func(v int) float64 {
			return float64(v)
		},
	}
	cyc := itr.Cycle()
	var res []float64
	for i := 0; i <= 7; i++ {
		res = append(res, cyc.Next())
	}
	require.Equal(t, []float64{1.0, 2.0, 3.0, 1.0, 2.0, 3.0, 1.0, 2.0}, res)
}
