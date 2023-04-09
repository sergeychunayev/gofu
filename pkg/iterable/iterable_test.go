package iterable_test

import (
	"github.com/sergeychunayev/gofu/pkg/iterable"
	"github.com/sergeychunayev/gofu/pkg/iterable/ord"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		check func(t *testing.T, itr iterable.Iterable[int])
	}{
		{
			"nil",
			nil,
			func(t *testing.T, itr iterable.Iterable[int]) {
				require.False(t, itr.HasNext())
			},
		},
		{
			"empty",
			[]int{},
			func(t *testing.T, itr iterable.Iterable[int]) {
				require.False(t, itr.HasNext())
			},
		},
		{
			"not empty",
			[]int{0},
			func(t *testing.T, itr iterable.Iterable[int]) {
				require.True(t, itr.HasNext())
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			tc.check(t, itr)
		})
	}
}

func TestSlice_Filter(t *testing.T) {
	itr := iterable.
		New([]int{1, 2, 3, 4, 5, 5, 5, 5, 8}).
		Filter(func(v int) bool {
			return v%2 == 0
		})
	require.True(t, itr.HasNext())
	require.Equal(t, 2, itr.Next())
	require.Equal(t, 4, itr.Next())
	require.Equal(t, 8, itr.Next())
}

func TestIterable(t *testing.T) {
	arr := []int{1, 2, 3}
	var itr iterable.Iterable[int]
	itr = iterable.New(arr)
	var res []int
	for itr.HasNext() {
		v := itr.Next()
		res = append(res, v)
	}
	require.Equal(t, arr, res)

	bools := iterable.Map(iterable.New(arr), func(v int) bool {
		return v%2 == 0
	}).ToSlice()
	require.Equal(t, []bool{false, true, false}, bools)

	allTrue, ok := iterable.New(bools).Reduce(func(acc bool, v bool) bool {
		return acc && v
	})
	require.True(t, ok)
	require.False(t, allTrue)

	someTrue, ok := iterable.New(bools).Reduce(func(acc bool, v bool) bool {
		return acc || v
	})
	require.True(t, ok)
	require.True(t, someTrue)

	filter := iterable.
		New(bools).
		Filter(func(v bool) bool {
			return v
		})
	trues := filter.ToSlice()
	require.Equal(t, []bool{true}, trues)

	falses := iterable.New(bools).Filter(func(v bool) bool {
		return !v
	}).ToSlice()
	require.Equal(t, []bool{false, false}, falses)

	cycle := iterable.New(arr).Cycle()
	require.Equal(t, 1, cycle.Next())
	require.Equal(t, 2, cycle.Next())
	require.Equal(t, 3, cycle.Next())
	require.Equal(t, 1, cycle.Next())
}

func TestFor(t *testing.T) {
	arr := []int{1, 2, 3}
	itr := iterable.New(arr)
	var res []iterable.Tuple[int, int]
	itr.For(func(v int, i int) {
		res = append(res, iterable.Tuple[int, int]{A: v, B: i})
	})
	expected := []iterable.Tuple[int, int]{
		{1, 0},
		{2, 1},
		{3, 2},
	}
	require.Equal(t, expected, res)
}

func TestAll(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		f        func(v int) bool
		expected bool
	}{
		{
			"Empty",
			[]int{},
			func(v int) bool {
				return v > 0
			},
			true,
		},
		{
			"All gt 0 true",
			[]int{1, 2, 3},
			func(v int) bool {
				return v > 0
			},
			true,
		},
		{
			"Not all gt 0",
			[]int{1, 2, 0},
			func(v int) bool {
				return v > 0
			},
			false,
		},
		{
			"All lte 0",
			[]int{-1, -2, 0},
			func(v int) bool {
				return v > 0
			},
			false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res := itr.All(tc.f)
			require.Equal(t, tc.expected, res)
		})
	}
}

func TestAny(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		f        func(v int) bool
		expected bool
	}{
		{
			"Empty",
			[]int{},
			func(v int) bool {
				return v > 0
			},
			false,
		},
		{
			"All gt 0 true",
			[]int{1, 2, 3},
			func(v int) bool {
				return v > 0
			},
			true,
		},
		{
			"Not all gt 0",
			[]int{0, 0, 1},
			func(v int) bool {
				return v > 0
			},
			true,
		},
		{
			"All lte 0",
			[]int{0, 0, -1},
			func(v int) bool {
				return v > 0
			},
			false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res := itr.Any(tc.f)
			require.Equal(t, tc.expected, res)
		})
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected iterable.Tuple[int, bool]
	}{
		{
			"Empty",
			[]int{},
			iterable.Tuple[int, bool]{},
		},
		{
			"Increasing",
			[]int{1, 2, 3},
			iterable.Tuple[int, bool]{A: 1, B: true},
		},
		{
			"Decreasing",
			[]int{3, 2, 1},
			iterable.Tuple[int, bool]{A: 1, B: true},
		},
		{
			"All 0",
			[]int{0, 0, 0},
			iterable.Tuple[int, bool]{B: true},
		},
		{
			"One 0",
			[]int{0},
			iterable.Tuple[int, bool]{B: true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res, ok := itr.Reduce(func(acc int, v int) int {
				if v < acc {
					return v
				}
				return acc
			})
			tup := iterable.Tuple[int, bool]{A: res, B: ok}
			require.Equal(t, tc.expected, tup)
		})
	}
}

func TestMax(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected iterable.Tuple[int, bool]
	}{
		{
			"Empty",
			[]int{},
			iterable.Tuple[int, bool]{},
		},
		{
			"Increasing",
			[]int{1, 2, 3},
			iterable.Tuple[int, bool]{A: 3, B: true},
		},
		{
			"Decreasing",
			[]int{3, 2, 1},
			iterable.Tuple[int, bool]{A: 3, B: true},
		},
		{
			"All 0",
			[]int{0, 0, 0},
			iterable.Tuple[int, bool]{B: true},
		},
		{
			"One 0",
			[]int{0},
			iterable.Tuple[int, bool]{B: true},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res, ok := itr.Reduce(func(acc int, v int) int {
				if v > acc {
					return v
				}
				return acc
			})
			tup := iterable.Tuple[int, bool]{A: res, B: ok}
			require.Equal(t, tc.expected, tup)
		})
	}
}

func TestSort(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		check func(t *testing.T, input []int, result []int)
	}{
		{
			"Empty",
			[]int{},
			func(t *testing.T, input []int, result []int) {
				require.Empty(t, input)
				require.Empty(t, result)
			},
		},
		{
			"Increasing",
			[]int{1, 2, 3},
			func(t *testing.T, input []int, result []int) {
				require.Equal(t, []int{1, 2, 3}, input)
				require.Equal(t, []int{1, 2, 3}, result)
			},
		},
		{
			"Decreasing",
			[]int{3, 2, 1},
			func(t *testing.T, input []int, result []int) {
				require.Equal(t, []int{1, 2, 3}, input)
				require.Equal(t, []int{1, 2, 3}, result)
			},
		},
		{
			"All 0",
			[]int{0, 0, 0},
			func(t *testing.T, input []int, result []int) {
				require.Equal(t, []int{0, 0, 0}, input)
				require.Equal(t, []int{0, 0, 0}, result)
			},
		},
		{
			"One 0",
			[]int{0},
			func(t *testing.T, input []int, result []int) {
				require.Equal(t, []int{0}, input)
				require.Equal(t, []int{0}, result)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res := itr.Sort(ord.Lt[int])
			tc.check(t, tc.input, res.ToSlice())
		})
	}
}

func TestSortStruct(t *testing.T) {
	type s struct {
		intV int
		strV string
	}
	testCases := []struct {
		name  string
		input []s
		check func(t *testing.T, input []s, result []s)
	}{
		{
			"Empty",
			[]s{},
			func(t *testing.T, input []s, result []s) {
				require.Empty(t, input)
				require.Empty(t, result)
			},
		},
		{
			"Increasing",
			[]s{
				{1, "1"},
				{2, "2"},
				{3, "3"},
			},
			func(t *testing.T, input []s, result []s) {
				expected := []s{
					{1, "1"},
					{2, "2"},
					{3, "3"},
				}
				require.Equal(t, expected, input)
				require.Equal(t, expected, result)
			},
		},
		{
			"Decreasing",
			[]s{
				{3, "3"},
				{2, "2"},
				{1, "1"},
			},
			func(t *testing.T, input []s, result []s) {
				expected := []s{
					{1, "1"},
					{2, "2"},
					{3, "3"},
				}
				require.Equal(t, expected, input)
				require.Equal(t, result, result)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res := itr.Sort(func(a s, b s) bool {
				return a.intV < b.intV
			})
			tc.check(t, tc.input, res.ToSlice())
		})
	}
}

func TestSortStructDesc(t *testing.T) {
	type s struct {
		intV int
		strV string
	}
	testCases := []struct {
		name  string
		input []s
		check func(t *testing.T, input []s, result []s)
	}{
		{
			"Empty",
			[]s{},
			func(t *testing.T, input []s, result []s) {
				require.Empty(t, input)
				require.Empty(t, result)
			},
		},
		{
			"Increasing",
			[]s{
				{1, "1"},
				{2, "2"},
				{3, "3"},
			},
			func(t *testing.T, input []s, result []s) {
				expected := []s{
					{3, "3"},
					{2, "2"},
					{1, "1"},
				}
				require.Equal(t, expected, input)
				require.Equal(t, expected, result)
			},
		},
		{
			"Decreasing",
			[]s{
				{3, "3"},
				{2, "2"},
				{1, "1"},
			},
			func(t *testing.T, input []s, result []s) {
				expected := []s{
					{3, "3"},
					{2, "2"},
					{1, "1"},
				}
				require.Equal(t, expected, input)
				require.Equal(t, result, result)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res := itr.Sort(func(a s, b s) bool {
				return a.intV > b.intV
			})
			tc.check(t, tc.input, res.ToSlice())
		})
	}
}

func TestCycle(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		check func(t *testing.T, input []int, result iterable.Iterable[int])
	}{
		{
			"Empty",
			[]int{},
			func(t *testing.T, input []int, result iterable.Iterable[int]) {
				require.Empty(t, input)
				require.False(t, result.HasNext())
			},
		},
		{
			"Loop",
			[]int{1, 2, 3},
			func(t *testing.T, input []int, result iterable.Iterable[int]) {
				require.Equal(t, []int{1, 2, 3}, input)
				var arr []int
				i := 0
				for result.HasNext() && i <= len(input)*2 {
					arr = append(arr, result.Next())
					i++
				}
				expected := []int{1, 2, 3, 1, 2, 3, 1}
				require.Equal(t, expected, arr)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			itr := iterable.New(tc.input)
			res := itr.Cycle()
			tc.check(t, tc.input, res)
		})
	}
}

func TestZip(t *testing.T) {
	testCases := []struct {
		name     string
		a        []int
		b        []int
		expected []iterable.Tuple[int, int]
	}{
		{
			"Empty",
			nil,
			nil,
			nil,
		},
		{
			"Same size",
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]iterable.Tuple[int, int]{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
		{
			"a has more elements",
			[]int{1, 2, 3, 4},
			[]int{4, 5, 6},
			[]iterable.Tuple[int, int]{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
		{
			"b has more elements",
			[]int{1, 2, 3},
			[]int{4, 5, 6, 7},
			[]iterable.Tuple[int, int]{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			aIt := iterable.New(tc.a)
			bIt := iterable.New(tc.b)
			res := iterable.Zip(aIt, bIt)
			require.Equal(t, tc.expected, res.ToSlice())
		})
	}
}

func TestMap(t *testing.T) {
	res := iterable.Map(iterable.New([]int{1, 2, 3}), func(v int) float64 {
		return float64(v)
	}).ToSlice()
	require.Equal(t, []float64{1.0, 2.0, 3.0}, res)
}

func TestFold(t *testing.T) {
	type s struct {
		strVal string
		intVal int
	}
	arr := []s{
		{"a", 1},
		{"b", 2},
	}
	itr := iterable.New(arr)
	res := iterable.Fold(itr, func(acc map[string]s, v s) map[string]s {
		acc[v.strVal] = v
		return acc
	}, make(map[string]s))
	expected := map[string]s{
		"a": {"a", 1},
		"b": {"b", 2},
	}
	require.Equal(t, expected, res)
}

func TestGroupBy(t *testing.T) {
	type s struct {
		strVal string
		intVal int
	}
	arr := []s{
		{"a", 1},
		{"b", 2},
		{"a", 3},
	}
	itr := iterable.New(arr)
	res := iterable.GroupBy(itr, func(v s) string {
		return v.strVal
	})
	expected := map[string][]s{
		"a": {{"a", 1}, {"a", 3}},
		"b": {{"b", 2}},
	}
	require.Equal(t, expected, res)
}
