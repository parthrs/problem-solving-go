package pkg

/*
Here are the ways a deque could be implemented:
 - Using a dynamic array:
   - Insertions to the front are ammortized to O(1)
	 - Insertions/deletions in the middle are O(n)
	 - Random access by index is O(1)
 - Using a LinkedList:
   - All insertions are O(1)
	 - Insertions/deletions in the middle are O(1)
	 - Random access by index is O(n)
*/

/*
Implementing Deque using a Doubly linked list
  - O(1) push and pop both ends
  - PushFront and PopFront are at Tail
  - PushBack and PopBack are at Head
*/
type Deque[T comparable] struct {
	Cap  int
	Len  int
	Head *DoublyLinkedNode[T]
	Tail *DoublyLinkedNode[T]
}

func NewDeque[T comparable](size int) *Deque[T] {
	return &Deque[T]{
		Cap:  size,
		Len:  0,
		Head: nil,
		Tail: nil,
	}
}

func (d *Deque[T]) PushFront(val T) bool {
	if d.Len == d.Cap {
		return false
	}
	n := NewDoublyLinkedNode[T](val)
	if d.Head == nil {
		d.Tail = n
	} else {
		d.Head.Previous = n
	}
	n.Next = d.Head
	d.Head = n

	d.Len++
	return true
}

func (d *Deque[T]) PushBack(val T) bool {
	if d.Len == d.Cap {
		return false
	}
	n := NewDoublyLinkedNode[T](val)
	if d.Tail == nil {
		d.Head = n
	} else {
		d.Tail.Next = n
	}
	n.Previous = d.Tail
	d.Tail = n
	d.Len++
	return true
}

func (d *Deque[T]) PopFront() (retVal T, success bool) {
	if d.Len == 0 {
		success = false
		return
	}
	n := d.Head
	if d.Head.Next == nil {
		d.Tail = nil
	} else {
		d.Head.Next.Previous = nil
	}
	d.Head = d.Head.Next
	n.Next = nil
	d.Len--
	return n.Value, true
}

func (d *Deque[T]) PopBack() (retVal T, success bool) {
	if d.Len == 0 {
		success = false
		return
	}
	n := d.Tail
	if d.Tail.Previous == nil {
		d.Head = nil
	} else {
		d.Tail.Previous.Next = nil
	}
	d.Tail = d.Tail.Previous
	n.Previous = nil
	d.Len--
	return n.Value, true
}
