package godoc_generics_issue

type MyType[T any] struct {
	i T
}

func (m *MyType[T]) Set(i T) {
	m.i = i
}

func (m *MyType[T]) Get() T {
	return m.i
}
