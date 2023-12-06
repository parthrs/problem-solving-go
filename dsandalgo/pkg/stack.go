package pkg

type Stack[T any] struct {
	S []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{S: []T{}}
}

func (s *Stack[T]) Add(elem T) {
	s.S = append(s.S, elem)
}

func (s *Stack[T]) Pop() T {
	elem := s.S[len(s.S)-1]
	s.S = s.S[:len(s.S)-1]
	return elem
}
