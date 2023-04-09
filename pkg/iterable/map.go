package iterable

type mapIterable[T any, U any] struct {
	itr  Iterable[T]
	mapF func(v T) U
}

func (v *mapIterable[T, U]) HasNext() bool {
	return v.itr.HasNext()
}

func (v *mapIterable[T, U]) Next() U {
	return v.mapF(v.itr.Next())
}

func (v *mapIterable[T, U]) Filter(f func(v U) bool) Iterable[U] {
	return &filterIterable[U]{
		v,
		f,
		nil,
	}
}

func (v *mapIterable[T, U]) For(f func(v U, i int)) {
	doFor[U](v, f)
}

func (v *mapIterable[T, U]) All(f func(v U) bool) bool {
	return all[U](v, f)
}

func (v *mapIterable[T, U]) Any(f func(v U) bool) bool {
	return doAny[U](v, f)
}

func (v *mapIterable[T, U]) Min(cmp func(a U, b U) bool) (U, bool) {
	return min[U](v, cmp)
}

func (v *mapIterable[T, U]) Max(cmp func(a U, b U) bool) (U, bool) {
	return max[U](v, cmp)
}

func (v *mapIterable[T, U]) Reduce(f func(acc U, v U) U) (U, bool) {
	return reduce[U](v, f)
}

func (v *mapIterable[T, U]) Sort(less func(a U, b U) bool) Iterable[U] {
	return doSort[U](v, less)
}

func (v *mapIterable[T, U]) Cycle() Iterable[U] {
	var res = v.ToSlice()
	return New(res).Cycle()
}

func (v *mapIterable[T, U]) ToSlice() []U {
	return toSlice[U](v)
}
