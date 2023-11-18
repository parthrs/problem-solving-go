package internal

/*
	Go inherently does not have a set data structure
	1. Order is not maintained
	2. The type constraint is comparable as that is a
	   underlying map constraint (Map can't lookup elem if
		 not comparable)
*/

type Set[T comparable] interface {
	Contains(T) bool
	Add(T)
	Remove() bool
}

type SetDS[T comparable] struct {
	Store map[T]bool
}

func NewSetDS[T comparable]() *SetDS[T] {
	return &SetDS[T]{
		Store: map[T]bool{},
	}
}

func (s *SetDS[T]) Contains(elem T) (found bool) {
	_, found = s.Store[elem]
	return
}

func (s *SetDS[T]) Add(elem T) {
	if _, found := s.Store[elem]; found {
		return
	}
	s.Store[elem] = true
}

func (s *SetDS[T]) Remove(elem T) (found bool) {
	if _, found = s.Store[elem]; found {
		delete(s.Store, elem)
	}
	return
}

func Union[T comparable](s1, s2 *SetDS[T]) (result *SetDS[T]) {
	result = NewSetDS[T]()
	for k := range s1.Store {
		result.Add(k)
	}
	for k := range s2.Store {
		result.Add(k)
	}
	return
}

func Intersection[T comparable](s1, s2 *SetDS[T]) (result *SetDS[T]) {
	result = NewSetDS[T]()
	shorter := s1 // Pointer to the shorter set
	longer := s2
	if len(s2.Store) < len(s1.Store) {
		shorter, longer = s2, s1 // Swap pointers
	}

	for k := range shorter.Store {
		if found := longer.Store[k]; found {
			result.Add(k)
		}
	}
	return
}
