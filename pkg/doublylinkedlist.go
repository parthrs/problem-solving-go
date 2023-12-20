package pkg

type DoublyLinkedNode[T comparable] struct {
	Previous *DoublyLinkedNode[T]
	Next     *DoublyLinkedNode[T]
	Value    T
}

type DoublyLinkedList[T comparable] struct {
	Head *DoublyLinkedNode[T]
	Tail *DoublyLinkedNode[T]
}

func NewDoublyLinkedNode[T comparable](val T) *DoublyLinkedNode[T] {
	return &DoublyLinkedNode[T]{
		Value: val,
	}
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Head: nil,
		Tail: nil,
	}
}

func (l *DoublyLinkedList[T]) AddNodeHead(val T) {
	n := NewDoublyLinkedNode[T](val)
	n.Next = l.Head
	if l.Head == nil {
		l.Tail = n
	} else {
		l.Head.Previous = n
	}
	l.Head = n
}

func (l *DoublyLinkedList[T]) AddNodeTail(val T) {
	n := NewDoublyLinkedNode[T](val)
	n.Previous = l.Tail
	if l.Tail == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func (l *DoublyLinkedList[T]) RemoveNodeHead() {
	// This returns if there's only one elem
	// Its better this way than having a clause that
	// only checks for nil and returns.
	if l.Head != nil {
		// Its better to do this rather than check
		// if head and tail point to the same elem
		// (list with single node), because it helps
		// in checking the head.next and setting
		// head.next.Previous to nil if it exists (which)
		// we'll have to do anyways.
		if l.Head.Next == nil {
			l.Tail, l.Head = nil, nil
		} else {
			n := l.Head
			l.Head = l.Head.Next
			l.Head.Previous = nil
			n.Next = nil
		}
	}
}

func (l *DoublyLinkedList[T]) RemoveNodeTail() {
	if l.Tail != nil {
		if l.Tail.Previous == nil {
			l.Head, l.Tail = nil, nil
		} else {
			n := l.Tail
			l.Tail = l.Tail.Previous
			l.Tail.Next = nil
			n.Previous = nil
		}
	}
}

func (l *DoublyLinkedList[T]) RemoveNode(val T) {
	if l.Head.Value == val {
		l.RemoveNodeHead()
		return
	}
	if l.Tail.Value == val {
		l.RemoveNodeTail()
		return
	}
	n := l.Head.Next
	for n.Value != val {
		n = n.Next
	}
	n.Previous.Next = n.Next
	n.Next.Previous = n.Previous
}
