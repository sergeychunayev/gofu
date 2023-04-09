package iterable

type cycleIterable[T any] struct {
	*Slice[T]
}

func (v *cycleIterable[T]) HasNext() bool {
	return len(v.slice) > 0
}

func (v *cycleIterable[T]) Next() T {
	res := v.slice[v.i]
	v.i++
	if v.i == len(v.slice) {
		v.i = 0
	}

	return res
}
