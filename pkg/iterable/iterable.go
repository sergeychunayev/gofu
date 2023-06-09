package iterable

type Iterable[T any] interface {
	HasNext() bool

	Next() T

	Filter(f func(v T) bool) Iterable[T]

	For(func(v T, i int))

	All(f func(v T) bool) bool

	Any(f func(v T) bool) bool

	Reduce(f func(acc T, v T) T) (T, bool)

	Sort(less func(a T, b T) bool) Iterable[T]

	Cycle() Iterable[T]

	ToSlice() []T
}

type Slice[T any] struct {
	slice []T
	i     int
}

type Tuple[A any, B any] struct {
	A A
	B B
}

func (v *Slice[T]) HasNext() bool {
	return v.i < len(v.slice)
}

func (v *Slice[T]) Next() T {
	res := v.slice[v.i]
	v.i++
	return res
}

func (v *Slice[T]) Filter(f func(v T) bool) Iterable[T] {
	return &filterIterable[T]{
		v,
		f,
		nil,
	}
}

func (v *Slice[T]) For(f func(v T, i int)) {
	doFor[T](v, f)
}

func (v *Slice[T]) All(f func(v T) bool) bool {
	return all[T](v, f)
}

func (v *Slice[T]) Any(f func(v T) bool) bool {
	return doAny[T](v, f)
}

func (v *Slice[T]) Reduce(f func(acc T, v T) T) (T, bool) {
	return reduce[T](v, f)
}

func (v *Slice[T]) Sort(less func(a T, b T) bool) Iterable[T] {
	return doSort[T](v, less)
}

func (v *Slice[T]) Cycle() Iterable[T] {
	return &cycleIterable[T]{v}
}

func (v *Slice[T]) ToSlice() []T {
	return v.slice
}

func New[T any](slice []T) Iterable[T] {
	return &Slice[T]{slice, 0}
}

func Map[T any, U any](v Iterable[T], f func(v T) U) Iterable[U] {
	return &mapIterable[T, U]{v, f}
}

func Fold[T any, U any](v Iterable[T], f func(acc U, v T) U, initial U) U {
	res := initial
	for v.HasNext() {
		n := v.Next()
		res = f(res, n)
	}
	return res
}

func Zip[A any, B any](aIt Iterable[A], bIt Iterable[B]) Iterable[Tuple[A, B]] {
	var res []Tuple[A, B]
	for aIt.HasNext() && bIt.HasNext() {
		res = append(res, Tuple[A, B]{aIt.Next(), bIt.Next()})
	}
	return New(res)
}

func GroupBy[T any, K comparable](it Iterable[T], keyF func(v T) K) map[K][]T {
	return Fold(it, func(acc map[K][]T, v T) map[K][]T {
		key := keyF(v)
		arr, ok := acc[key]
		if !ok {
			arr = []T{}
		}
		arr = append(arr, v)
		acc[key] = arr
		return acc
	}, make(map[K][]T))
}
