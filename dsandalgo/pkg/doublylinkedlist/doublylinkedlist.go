package doublylinkedlist

type DoublyLinkedNode struct {
	Previous *DoublyLinkedNode
	Next     *DoublyLinkedNode
	Value    int32
}

type DoublyLinkedList struct {
	Head *DoublyLinkedNode
	Tail *DoublyLinkedNode
	//length int32
}

func NewDoublyLinkedNode(val int32) *DoublyLinkedNode {
	return &DoublyLinkedNode{
		Previous: nil,
		Next:     nil,
		Value:    val,
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
		//length: int32(0),
	}
}

func (l *DoublyLinkedList) AddNodeHead(val int32) {
	n := NewDoublyLinkedNode(val)
	n.Next = l.Head
	n.Previous = nil

	// Check if list is empty
	if l.Head == nil {
		l.Tail = n
	} else {
		l.Head.Previous = n
	}

	l.Head = n
}

func (l *DoublyLinkedList) AddNodeTail(val int32) {
	n := NewDoublyLinkedNode(val)
	n.Next = nil
	n.Previous = l.Tail

	if l.Tail == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func (l *DoublyLinkedList) RemoveNodeHead() {
	if l.Head != nil {
		if l.Head.Next == nil {
			l.Head, l.Tail = nil, nil
		} else {
			l.Head = l.Head.Next
			l.Head.Previous = nil
		}
	}
}

func (l *DoublyLinkedList) RemoveNodeTail() {
	if l.Tail != nil {
		if l.Tail.Previous == nil {
			l.Tail, l.Head = nil, nil
		} else {
			l.Tail = l.Tail.Previous
			l.Tail.Next = nil
		}
	}
}

func (l *DoublyLinkedList) RemoveNode(val int32) {

	// List with one node or if the first node is what we seek
	if l.Head.Value == val {
		l.RemoveNodeHead()
		return
	}

	if l.Tail.Value == val {
		l.RemoveNodeTail()
		return
	}

	n := l.Head.Next
	// Other cases
	for n.Value != val {
		n = n.Next
	}
	n.Previous.Next = n.Next
	n.Next.Previous = n.Previous
}
