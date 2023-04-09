package iterable

import "sort"

type Iterable[T any] interface {
	HasNext() bool

	Next() T

	Filter(f func(v T) bool) Iterable[T]

	For(func(v T, i int))

	All(f func(v T) bool) bool

	Any(f func(v T) bool) bool

	Min(cmp func(a T, b T) bool) (T, bool)

	Max(cmp func(a T, b T) bool) (T, bool)

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
	var res []T
	for v.HasNext() {
		val := v.Next()
		if f(val) {
			res = append(res, val)
		}
	}
	return New(res)
}

func (v *Slice[T]) For(f func(v T, i int)) {
	i := 0
	for v.HasNext() {
		f(v.Next(), i)
		i++
	}
}

func (v *Slice[T]) All(f func(v T) bool) bool {
	for v.HasNext() {
		if !f(v.Next()) {
			return false
		}
	}
	return true
}

func (v *Slice[T]) Any(f func(v T) bool) bool {
	for v.HasNext() {
		if f(v.Next()) {
			return true
		}
	}
	return false
}

func (v *Slice[T]) Min(cmp func(a T, b T) bool) (T, bool) {
	return v.Reduce(func(acc T, v T) T {
		if cmp(acc, v) {
			return acc
		}
		return v
	})
}

func (v *Slice[T]) Max(cmp func(a T, b T) bool) (T, bool) {
	return v.Reduce(func(acc T, v T) T {
		if cmp(acc, v) {
			return acc
		}
		return v
	})
}

func (v *Slice[T]) Reduce(f func(acc T, v T) T) (T, bool) {
	var res T
	if !v.HasNext() {
		return res, false
	}

	res = Fold[T, T](v, f, v.Next())
	return res, true
}

func (v *Slice[T]) Sort(less func(a T, b T) bool) Iterable[T] {
	sort.Slice(v.slice, func(i int, j int) bool {
		return less(v.slice[i], v.slice[j])
	})
	return New(v.slice)
}

func (v *Slice[T]) Cycle() Iterable[T] {
	return &Cycle[T]{v}
}

func (v *Slice[T]) ToSlice() []T {
	return v.slice
}

func New[T any](slice []T) Iterable[T] {
	return &Slice[T]{slice, 0}
}

type Cycle[T any] struct {
	*Slice[T]
}

func (v *Cycle[T]) HasNext() bool {
	return len(v.slice) > 0
}

func (v *Cycle[T]) Next() T {
	res := v.slice[v.i]
	v.i++
	if v.i == len(v.slice) {
		v.i = 0
	}

	return res
}

func Map[T any, U any](v Iterable[T], f func(v T) U) Iterable[U] {
	var res []U
	for v.HasNext() {
		u := f(v.Next())
		res = append(res, u)
	}
	return New(res)
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
