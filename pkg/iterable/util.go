package iterable

import "sort"

func doFor[T any](it Iterable[T], f func(v T, i int)) {
	i := 0
	for it.HasNext() {
		f(it.Next(), i)
		i++
	}
}

func all[T any](it Iterable[T], f func(v T) bool) bool {
	for it.HasNext() {
		if !f(it.Next()) {
			return false
		}
	}
	return true
}

func doAny[T any](it Iterable[T], f func(v T) bool) bool {
	for it.HasNext() {
		if f(it.Next()) {
			return true
		}
	}
	return false
}

func min[T any](it Iterable[T], cmp func(a T, b T) bool) (T, bool) {
	return it.Reduce(func(acc T, v T) T {
		if cmp(acc, v) {
			return acc
		}
		return v
	})
}

func max[T any](it Iterable[T], cmp func(a T, b T) bool) (T, bool) {
	return it.Reduce(func(acc T, v T) T {
		if cmp(acc, v) {
			return acc
		}
		return v
	})
}

func reduce[T any](it Iterable[T], f func(acc T, v T) T) (T, bool) {
	var res T
	if !it.HasNext() {
		return res, false
	}

	res = Fold[T, T](it, f, it.Next())
	return res, true
}

func doSort[T any](it Iterable[T], less func(a T, b T) bool) Iterable[T] {
	var res = it.ToSlice()
	sort.Slice(res, func(i int, j int) bool {
		return less(res[i], res[j])
	})
	return New(res)
}

func cycle[T any](it Iterable[T]) Iterable[T] {
	var res = it.ToSlice()
	return New(res).Cycle()
}

func toSlice[T any](it Iterable[T]) []T {
	var res []T
	for it.HasNext() {
		res = append(res, it.Next())
	}
	return res
}
