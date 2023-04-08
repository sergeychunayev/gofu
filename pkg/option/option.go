package option

type Option[T any] interface {
	opt()
}

type opt[T any] struct {
	V *T
}

type None[T any] opt[T]

//var N = NewNone()

type Some[T any] opt[T]

func (v None[T]) opt() {
}

func (v Some[T]) opt() {
}

func New[T any](v *T) Option[T] {
	if v == nil {
		return NONE
	}
	return Some[T]{v}
}

var NONE = newNone[any]()

func newNone[T any]() Option[T] {
	return None[T]{}
}
