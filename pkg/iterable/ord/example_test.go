package ord_test

import (
	"fmt"
	"github.com/sergeychunayev/gofu/pkg/iterable"
	"github.com/sergeychunayev/gofu/pkg/iterable/ord"
)

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
