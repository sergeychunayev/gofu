package iterable

type filterIterable[T any] struct {
	Iterable[T]
	filter func(v T) bool
	cur    *T
}

func (v *filterIterable[T]) HasNext() bool {
	if v.cur != nil {
		return true
	}

	itr := v.Iterable
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

func (v *filterIterable[T]) ToSlice() []T {
	return toSlice[T](v)
}
