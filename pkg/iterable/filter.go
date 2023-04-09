package iterable

type filterIterable[T any] struct {
	itr    Iterable[T]
	filter func(v T) bool
	cur    *T
}

func (v *filterIterable[T]) HasNext() bool {
	if v.cur != nil {
		return true
	}

	itr := v.itr
	for itr.HasNext() {
		el := itr.Next()
		if v.filter(el) {
			v.cur = &el
			return true
		}
	}

	return false
}

func (v *filterIterable[T]) Next() T {
	res := *v.cur
	v.cur = nil
	v.HasNext()
	return res
}

func (v *filterIterable[T]) Filter(f func(v T) bool) Iterable[T] {
	return &filterIterable[T]{
		v,
		f,
		nil,
	}
}

func (v *filterIterable[T]) For(f func(v T, i int)) {
	doFor[T](v, f)
}

func (v *filterIterable[T]) All(f func(v T) bool) bool {
	return all[T](v, f)
}

func (v *filterIterable[T]) Any(f func(v T) bool) bool {
	return doAny[T](v, f)
}

func (v *filterIterable[T]) Min(cmp func(a T, b T) bool) (T, bool) {
	return min[T](v, cmp)
}

func (v *filterIterable[T]) Max(cmp func(a T, b T) bool) (T, bool) {
	return max[T](v, cmp)
}

func (v *filterIterable[T]) Reduce(f func(acc T, v T) T) (T, bool) {
	return reduce[T](v, f)
}

func (v *filterIterable[T]) Sort(less func(a T, b T) bool) Iterable[T] {
	return doSort[T](v, less)
}

func (v *filterIterable[T]) Cycle() Iterable[T] {
	var res = v.ToSlice()
	return New(res).Cycle()
}

func (v *filterIterable[T]) ToSlice() []T {
	return toSlice[T](v)
}
