[![CodeQL](https://github.com/sergeychunayev/gofu/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/sergeychunayev/gofu/actions/workflows/github-code-scanning/codeql)
[![Dependency Review](https://github.com/sergeychunayev/gofu/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/sergeychunayev/gofu/actions/workflows/dependency-review.yml)
[![Build](https://github.com/sergeychunayev/gofu/actions/workflows/build.yml/badge.svg)](https://github.com/sergeychunayev/gofu/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sergeychunayev/gofu)](https://goreportcard.com/report/github.com/sergeychunayev/gofu)

![logo](https://user-images.githubusercontent.com/57276805/230751843-c0972a6e-64ee-439f-a21c-054c678326bb.png)

# GoFu
## Go Functional

Functional programming for Go. With generics.

# Examples

### Filter and sort

```go
package main

import (
	"fmt"
	"github.com/sergeychunayev/gofu/pkg/iterable"
	"github.com/sergeychunayev/gofu/pkg/iterable/ord"
)

func main() {
	arr := iterable.New([]int{4, 3, 2, 1}).
		Filter(func(v int) bool {
			return v%2 == 0
		}).
		Sort(ord.Lt[int]).
		ToSlice()
	fmt.Println(arr) // [2 4]
}
```

### Filter and sort a struct

```go
package main

import (
	"fmt"

	"github.com/sergeychunayev/gofu/pkg/iterable"
)

func main() {
	type s struct {
		name  string
		value int
	}

	res := iterable.
		New([]s{
			{"c", 1},
			{"b", 2},
			{"d", 3},
			{"a", 4},
		}).
		Filter(func(v s) bool {
			return v.value%2 == 0
		}).
		Sort(func(a s, b s) bool {
			return a.value < b.value
		}).
		ToSlice()

	fmt.Printf("res: %+v\n", res) // res: [{name:b value:2} {name:a value:4}]
}

```

### GroupBy

```go
package main

import (
	"fmt"

	"github.com/sergeychunayev/gofu/pkg/iterable"
)

func main() {
	type s struct {
		name  string
		value int
	}

	itr := iterable.
		New([]s{
			{"c", 1},
			{"b", 2},
			{"d", 3},
			{"a", 4},
			{"a", 4},
			{"a", 6},
			{"f", 4},
			{"f", 8},
		}).
		Filter(func(v s) bool {
			return v.value%2 == 0
		})

	res := iterable.GroupBy(itr, func(v s) string {
		return v.name
	})

	// indentation made for clarity
	// res: map[
	// 		a:[ {name:a value:4} {name:a value:4} {name:a value:6} ]
	//		b:[ {name:b value:2} ]
	//		f:[ {name:f value:4} {name:f value:8} ]
	// 	]
	fmt.Printf("res: %+v\n", res)
}
```

### Min

```go
package main

import (
	"fmt"

	"github.com/sergeychunayev/gofu/pkg/iterable"
)

func main() {
	type s struct {
		name  string
		value int
	}

	min, _ := iterable.
		New([]s{
			{"b", 2},
			{"d", 3},
			{"a", 4},
			{"a", 4},
			{"a", 6},
			{"f", 4},
			{"f", 8},
			{"c", 1},
		}).
		Reduce(func(a s, b s) s {
			if a.value < b.value {
				return a
			}
			return b
		})

	fmt.Printf("res: %+v\n", min) // res: {name:c value:1}
}
```

### MinBy

```go
package main

import (
	"fmt"

	"github.com/sergeychunayev/gofu/pkg/iterable"
	"github.com/sergeychunayev/gofu/pkg/iterable/ord"
)

func main() {
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
```

### MaxBy

```go
package main

import (
	"fmt"

	"github.com/sergeychunayev/gofu/pkg/iterable"
	"github.com/sergeychunayev/gofu/pkg/iterable/ord"
)

func main() {
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
```

### Reduce

```go
package main

import (
	"fmt"

	"github.com/sergeychunayev/gofu/pkg/iterable"
)

func main() {
	res, _ := iterable.
		New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}).
		Reduce(func(acc int, v int) int {
			return acc + v
		})

	fmt.Printf("res: %+v\n", res) // res: 55
}
```

## License

[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)
