package pkg

/*
Implementing a stack that does not involve memory
re-allocation (via an underlying array) for any
push and pop operations.

Implementing using a singly linked list

The difference between this and a standalone
singly linked list implementation: tracking
the size.

Push and pop are O(1)
*/

type Stack[T any] struct {
	Head *SinglyNode[T]
	Len  int
	Cap  int
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		Len:  0,
		Cap:  capacity,
		Head: nil,
	}
}

func (s *Stack[T]) Push(elem T) bool {
	if s.Len == s.Cap {
		return false // Element not added
	}
	node := NewSinglyNode[T](elem)
	node.Next = s.Head
	s.Head = node
	s.Len++
	return true
}

func (s *Stack[T]) Pop() (retVal T) {
	if s.Head == nil {
		return // return the default value for T
	}
	n := s.Head
	s.Head = s.Head.Next
	// Remove any reference this node is making to other nodes
	n.Next = nil
	retVal = n.Value
	s.Len--
	return
}
