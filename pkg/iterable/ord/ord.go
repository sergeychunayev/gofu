package ord

import (
	"github.com/sergeychunayev/gofu/pkg/iterable"
)

type Ord interface {
	~string |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Iterable[T Ord] interface {
	iterable.Iterable[T]

	SumOrd() (T, bool)

	Min() (T, bool)

	Max() (T, bool)
}

type Slice[T Ord] struct {
	iterable.Iterable[T]
}

func New[T Ord](slice []T) Iterable[T] {
	return &Slice[T]{iterable.New[T](slice)}
}

func MinBy[T any, K Ord](it iterable.Iterable[T], keyF func(v T) K) (T, bool) {
	return it.Reduce(func(a T, b T) T {
		ak := keyF(a)
		bk := keyF(b)
		if bk < ak {
			return b
		}
		return a
	})
}

func MaxBy[T any, K Ord](it iterable.Iterable[T], keyF func(v T) K) (T, bool) {
	return it.Reduce(func(a T, b T) T {
		ak := keyF(a)
		bk := keyF(b)
		if bk > ak {
			return b
		}
		return a
	})
}

func Lt[T Ord](a T, b T) bool {
	return a < b
}

func Gt[T Ord](a T, b T) bool {
	return a > b
}

func Min[T Ord](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Ord](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func (v *Slice[T]) SumOrd() (T, bool) {
	return v.Reduce(func(acc T, v T) T {
		return acc + v
	})
}

func (v *Slice[T]) Min() (T, bool) {
	return v.Reduce(Min[T])
}

func (v *Slice[T]) Max() (T, bool) {
	return v.Reduce(Max[T])
}

func Sum[T Ord](it []T) (T, bool) {
	var res T
	if len(it) == 0 {
		return res, false
	}
	for _, v := range it {
		res += v
	}
	return res, true
}
