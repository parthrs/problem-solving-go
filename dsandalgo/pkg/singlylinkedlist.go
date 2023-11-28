package pkg

type SinglyNode[T any] struct {
	Value T
	Next  *SinglyNode[T]
}

func NewSinglyNode[T any](val T) *SinglyNode[T] {
	return &SinglyNode[T]{
		Value: val,
		Next:  nil,
	}
}
