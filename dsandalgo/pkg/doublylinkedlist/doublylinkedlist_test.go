package doublylinkedlist

import "testing"

func TestDoublyLinkedNode(t *testing.T) {
	n := NewDoublyLinkedNode(5)
	if n.Next != nil && n.Previous != nil {
		t.Errorf("Newly created node has non-nil pointers\n")
	}
	if n.Value != int32(5) {
		t.Errorf("Newly created node does not reflect the desired Value")
	}
}

func TestDoublyLinkedList(t *testing.T) {
	l := NewDoublyLinkedList()
	l.AddNodeHead(1)
	l.AddNodeHead(0)
	l.AddNodeTail(2)
	l.AddNodeTail(3)

	if l.Head.Value != 0 {
		t.Errorf("List Head Value check failed.\n")
	}
	if l.Head.Next.Value != 1 {
		t.Errorf("Second node Value check failed.\n")
	}
	if l.Tail.Value != 3 {
		t.Errorf("List Tail Value check failed.\n")
	}
	if l.Tail.Previous.Value != 2 {
		t.Errorf("Second from the Tail node Value check failed.\n")
	}

	l.RemoveNodeTail()
	if l.Tail.Value != 2 {
		t.Errorf("List Tail Value check failed after removing node from Tail.\n")
	}

	l.RemoveNodeHead()
	if l.Head.Value != 1 {
		t.Errorf("List Head Value check failed after removing node from Head.\n")
	}

	if l.Head.Previous != nil {
		t.Errorf("First node Previous Value not nil.\n")
	}
	if l.Tail.Next != nil {
		t.Errorf("Last node Next Value not nil.\n")
	}

	if l.Head.Next != l.Tail {
		t.Errorf("First node Next not pointing to Tail node in a 2-node list.\n")
	}
	if l.Tail.Previous != l.Head {
		t.Errorf("Last node Previous not pointing to Head node in a 2-node list.\n")
	}

	l.AddNodeTail(3)
	l.RemoveNode(2)
	if l.Tail.Value != 3 {
		t.Errorf("List Tail Value check failed after removing center node.\n")
	}
}
