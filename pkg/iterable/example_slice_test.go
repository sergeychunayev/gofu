package iterable_test

import (
	"fmt"
	"github.com/sergeychunayev/gofu/pkg/iterable"
)

func ExampleSlice_HasNext() {
	itr := iterable.New([]int{4, 3, 2, 1})
	for itr.HasNext() {
		fmt.Println("next:", itr.Next())
	}
	// Output:
	// next: 4
	// next: 3
	// next: 2
	// next: 1
}

func ExampleSlice_Next() {
	itr := iterable.New([]int{4, 3, 2, 1})
	for itr.HasNext() {
		fmt.Println("next:", itr.Next())
	}
	// Output:
	// next: 4
	// next: 3
	// next: 2
	// next: 1
}

func ExampleSlice_Filter() {
	arr := iterable.New([]int{4, 3, 2, 1}).
		Filter(func(v int) bool {
			return v%2 == 0
		}).
		ToSlice()
	fmt.Println(arr)
	// Output: [4 2]
}

func ExampleSlice_For() {
	iterable.
		New([]int{4, 3, 2, 1}).
		For(func(v int, i int) {
			fmt.Printf("v: %d, i: %d\n", v, i)
		})
	// Output:
	// v: 4, i: 0
	// v: 3, i: 1
	// v: 2, i: 2
	// v: 1, i: 3
}

func ExampleSlice_All() {
	type S struct {
		name  string
		value int
	}
	res := iterable.
		New([]S{
			{"one", 1},
			{"two", 1},
			{"three", -1},
		}).
		All(func(v S) bool {
			return v.value > 0
		})
	fmt.Println(res)
	// Output: false
}

func ExampleSlice_Any() {
	type S struct {
		name  string
		value int
	}
	res := iterable.
		New([]S{
			{"one", -1},
			{"two", -1},
			{"three", 1},
		}).
		Any(func(v S) bool {
			return v.value > 0
		})
	fmt.Println(res)
	// Output: true
}

func ExampleSlice_Min() {
	type S struct {
		name  string
		value int
	}
	res, ok := iterable.
		New([]S{
			{"one", -1},
			{"two", -1},
			{"three", -100},
		}).
		Min(func(a S, b S) bool {
			return a.value < b.value
		})
	fmt.Println(ok)
	fmt.Println(res)
	// Output:
	// true
	// {three -100}
}

func ExampleSlice_Max() {
	type S struct {
		name  string
		value int
	}
	res, ok := iterable.
		New([]S{
			{"one", -1},
			{"two", -1},
			{"three", 100},
		}).
		Max(func(a S, b S) bool {
			return a.value > b.value
		})
	fmt.Println(ok)
	fmt.Println(res)
	// Output:
	// true
	// {three 100}
}

func ExampleSlice_Reduce() {
	res, _ := iterable.
		New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Reduce(func(acc int, v int) int {
			return acc + v
		})

	fmt.Println(res)
	// Output: 55
}

func ExampleSlice_Sort() {
	type s struct {
		name  string
		value int
	}

	res := iterable.
		New([]s{
			{"c", 4},
			{"b", 3},
			{"d", 2},
			{"a", 1},
		}).
		Sort(func(a s, b s) bool {
			return a.value < b.value
		}).
		ToSlice()

	fmt.Printf("%v\n", res)
	// Output: [{a 1} {d 2} {b 3} {c 4}]
}

func ExampleSlice_Cycle() {
	arr := []int{1, 2, 3}
	itr := iterable.New(arr).Cycle()
	i := 0
	for itr.HasNext() && i <= len(arr)*2 {
		fmt.Println(itr.Next())
		i++
	}
	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
	// 1
}

func ExampleSlice_ToSlice() {
	res := iterable.New([]int{1, 2, 3}).ToSlice()
	fmt.Println(res)
	// Output: [1 2 3]
}

func ExampleMap() {
	type S struct {
		name  string
		value int
	}
	itr := iterable.
		New([]S{
			{"one", 1},
			{"two", 1},
			{"three", 1},
		})
	res := iterable.Map(itr, func(v S) string {
		return v.name
	}).ToSlice()
	fmt.Println(res)
	// Output: [one two three]
}

func ExampleFold() {
	itr := iterable.New([]int{1, 2, 3})
	res := iterable.Fold(itr, func(acc map[bool][]int, v int) map[bool][]int {
		even := v%2 == 0
		arr := acc[even]
		arr = append(arr, v)
		acc[even] = arr
		return acc
	}, make(map[bool][]int))
	fmt.Println(res)
	// Output: map[false:[1 3] true:[2]]
}

func ExampleZip() {
	iterable.
		Zip(
			iterable.New([]int{-1, -2, -3, 4}),
			iterable.New([]int{1, 2, 3}),
		).
		For(func(v iterable.Tuple[int, int], _ int) {
			fmt.Println(v)
		})
	// Output:
	// {-1 1}
	// {-2 2}
	// {-3 3}
}

func ExampleGroupBy() {
	res := iterable.GroupBy(
		iterable.New([]int{1, 2, 3}),
		func(v int) bool {
			return v%2 == 0
		},
	)
	fmt.Println(res)
	// Output: map[false:[1 3] true:[2]]
}
