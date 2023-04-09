[![CodeQL](https://github.com/sergeychunayev/gofu/actions/workflows/codeql.yml/badge.svg)](https://github.com/sergeychunayev/gofu/actions/workflows/codeql.yml)
[![Dependency Review](https://github.com/sergeychunayev/gofu/actions/workflows/dependency-review.yml/badge.svg)](https://github.com/sergeychunayev/gofu/actions/workflows/dependency-review.yml)
[![Build](https://github.com/sergeychunayev/gofu/actions/workflows/build.yml/badge.svg)](https://github.com/sergeychunayev/gofu/actions/workflows/build.yml)

<img src="https://user-images.githubusercontent.com/57276805/230751843-c0972a6e-64ee-439f-a21c-054c678326bb.png" />

# GoFu
## Go Functional

Functional programming for Go.

# Examples

### Filter and sort
https://go.dev/play/p/82eV-chuxcf
```go
package main

import (
	"fmt"
	"github.com/sergeychunayev/gofu/pkg/iterable"
)

func main() {
	arr := iterable.New([]int{4, 3, 2, 1}).
		Filter(func(v int) bool {
			return v%2 == 0
		}).
		Sort(iterable.LtOrd[int]).
		ToSlice()
	fmt.Println(arr) // [2 4]
}
```

### Filter and sort a struct
https://go.dev/play/p/U4bge6gqGMH
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
https://go.dev/play/p/bAEcdAmreuM
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
https://go.dev/play/p/QeZSnWOAyOz
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
		Min(func(a s, b s) bool {
			return a.value < b.value
		})

	fmt.Printf("res: %+v\n", min) // res: {name:c value:1}

}
```

### Reduce
https://go.dev/play/p/JtUgM4qWB7q

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
