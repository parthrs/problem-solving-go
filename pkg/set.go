package pkg

/*
	Go inherently does not have a set data structure
	In this implementation,
	1. Order is not maintained
	2. The type constraint is comparable as that is a
	   underlying map constraint (Map can't lookup elem i.e.
		 use as a key if not comparable)

	Notes: In any func definition - if T is referenced
	       its type (comparable) must be declared at that definition
				 level. For instance in func NewSetDS. This makes sense since
				 the func cannot know what the type constraint is..
				 But the methods on the defined struct can know based on the
				 struct definition. So no need to specify constraint here.
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
	if _, found := s.Store[elem]; !found {
		s.Store[elem] = true
	}
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
