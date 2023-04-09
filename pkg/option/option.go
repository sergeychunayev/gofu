package option

type Option[T any] interface {
	IsNone() bool
	IsSome() bool
	Unwrap() T
	UnwrapOr(def T) T
}

type None[T any] struct{}

func (*None[T]) IsNone() bool {
	return true
}
func (*None[T]) IsSome() bool {
	return false
}
func (*None[T]) Unwrap() T {
	panic("None")
}
func (*None[T]) UnwrapOr(def T) T {
	return def
}

type Some[T any] struct {
	v T
}

func (*Some[T]) IsNone() bool {
	return false
}
func (*Some[T]) IsSome() bool {
	return true
}
func (v *Some[T]) Unwrap() T {
	return v.v
}
func (v *Some[T]) UnwrapOr(T) T {
	return v.v
}

func No[T any]() Option[T] {
	return &None[T]{}
}

func Of[T any](v T) Option[T] {
	return &Some[T]{v}
}
