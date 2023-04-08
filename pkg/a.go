package main

import (
	"fmt"
	"github.com/sergeychunayev/gofu/pkg/option"
)

func main() {
	v := 1
	fmt.Println("v:", v)
	val := option.New[int](nil)
	fmt.Printf("t: %T\n", val)
	switch val.(type) {
	case option.Some[int]:
		fmt.Println("Some")
		v := val.(option.Some[int]).V
		fmt.Println("v:", v)
	case option.None[any]:
		fmt.Println("None")
	default:
		fmt.Println("a")
	}
}
