package ord_test

import (
	"fmt"
	"github.com/sergeychunayev/gofu/pkg/iterable"
	"github.com/sergeychunayev/gofu/pkg/iterable/ord"
)

func ExampleSlice_SumOrd() {
	res, ok := ord.
		New[int]([]int{1, 2, 3}).
		SumOrd()
	fmt.Println(ok)
	fmt.Println(res)
	// Output:
	// true
	// 6
}

func ExampleSlice_Min() {
	res, ok := ord.
		New[int]([]int{3, 2, 1}).
		Min()
	fmt.Println(ok)
	fmt.Println(res)
	// Output:
	// true
	// 1
}

func ExampleSlice_Max() {
	res, ok := ord.
		New[int]([]int{1, 2, 3}).
		Max()
	fmt.Println(ok)
	fmt.Println(res)
	// Output:
	// true
	// 3
}

func ExampleMinBy() {
	type S struct {
		name  string
		value int
	}
	itr := iterable.New([]S{
		{"one", 1},
		{"two", 2},
		{"three", 3},
	})
	res, ok := ord.MinBy(itr, func(v S) int {
		return v.value
	})
	fmt.Println(ok)
	fmt.Println(res)
	// Output:
	// true
	// {one 1}
}

func ExampleMaxBy() {
	type S struct {
		name  string
		value int
	}
	itr := iterable.New([]S{
		{"one", 1},
		{"two", 2},
		{"three", 3},
	})
	res, ok := ord.MaxBy(itr, func(v S) int {
		return v.value
	})
	fmt.Println(ok)
	fmt.Println(res)
	// Output:
	// true
	// {three 3}
}
