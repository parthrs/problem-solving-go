package main

import (
	sll "github.com/parthrs/LetsGo/dsandalgo/pkg/singlylinkedlist"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := sll.NewSinglyNode(2)
	l1.Next = sll.NewSinglyNode(4)
	l1.Next.Next = sll.NewSinglyNode(3)

	l2 := sll.NewSinglyNode(5)
	l2.Next = sll.NewSinglyNode(6)
	l2.Next.Next = sll.NewSinglyNode(4)

	n := AddTwoNumbers(l1, l2)
	if n.Value != 7 || n.Next.Value != 0 || n.Next.Next.Value != 8 {
		t.Errorf("%d%d%d\n != 708", n.Value, n.Next.Value, n.Next.Next.Value)
	}

	l3 := sll.NewSinglyNode(0)
	l4 := sll.NewSinglyNode(1)

	n = AddTwoNumbers(l3, l4)
	if n.Value != 1 {
		t.Errorf("1 != %d", n.Value)
	}

	l5 := sll.NewSinglyNode(0)
	l6 := sll.NewSinglyNode(0)

	n = AddTwoNumbers(l5, l6)
	if n.Value != 0 {
		t.Errorf("0 != %d\n", n.Value)
	}

	/*
		ll1 := sll.NewSinglyNode(1)
		prev := ll1
		for i := 1; i < 30; i++ {
			ll := sll.NewSinglyNode(0)
			prev.Next = ll
			prev = ll
		}
		prev.Next = sll.NewSinglyNode(1)

		ll2 := sll.NewSinglyNode(5)
		ll2.Next = sll.NewSinglyNode(6)
		ll2.Next.Next = sll.NewSinglyNode(4)
	*/
}
