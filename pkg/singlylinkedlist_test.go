package pkg

import "testing"

func TestSinglyLinkedNode(t *testing.T) {
	n := NewSinglyNode[int](5)
	if n.Next != nil {
		t.Error("Newly created node has non-nil pointer\n")
	}
	if n.Value != 5 {
		t.Error("Newly created node does not reflect the desired value")
	}
}

func TestSinglyLinkedList(t *testing.T) {
	l := NewSinglyList[int]()
	l.AddNodeHead(3)
	if l.Head.Value != 3 {
		t.Errorf("Expected 3 at head, got %d", l.Head.Value)
	}
	l.AddNodeHead(2)
	l.AddNodeHead(1)
	if l.Head.Value != 1 {
		t.Errorf("Expected 1 at head, got %d", l.Head.Value)
	}
	if l.Head.Next.Value != 2 {
		t.Errorf("Expected 2 next to head, got %d", l.Head.Next.Value)
	}
	l.AddNodeTail(4)
	if l.Head.Next.Next.Next.Value != 4 {
		t.Errorf("Expected 4 next to head, got %d", l.Head.Next.Next.Next.Value)
	}
}
