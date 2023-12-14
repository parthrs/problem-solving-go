package pkg

/*
Implements a singly node and a list
The singly list only has one pointer, at head
and can insert nodes at head and tail
*/

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

type SinglyLinkedList[T any] struct {
	Head *SinglyNode[T]
}

func NewSinglyList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		Head: nil,
	}
}

func (l *SinglyLinkedList[T]) AddNodeHead(val T) {
	n := NewSinglyNode[T](val)
	n.Next = l.Head
	l.Head = n
}

func (l *SinglyLinkedList[T]) AddNodeTail(val T) {
	n := NewSinglyNode[T](val)
	if l.Head == nil {
		l.Head = n
	} else {
		// p points to the last node in the list
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = n
	}
}
