package singlylinkedlist

import "testing"

func TestSinglyLinkedNode(t *testing.T) {
	n := NewSinglyNode(5)
	if n.Next != nil {
		t.Error("Newly created node has non-nil pointer\n")
	}
	if n.Value != 5 {
		t.Error("Newly created node does not reflect the desired value")
	}
}
