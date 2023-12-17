package linkedlist

/*
Solve lc-2130 next

You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.
*/

import (
	"github.com/parthrs/problem-solving-go/pkg"
)

func ReorderList(head *pkg.SinglyNode[int]) {
	if head.Next == nil || head.Next.Next == nil {
		return
	}
	nodes := []*pkg.SinglyNode[int]{}
	next := head
	for next != nil {
		nodes = append(nodes, next)
		next = next.Next
	}

	nodePtr := len(nodes) - 1
	next = head.Next
	for i := 1; i <= (len(nodes) / 2); i++ {
		nodes[nodePtr-1].Next = nil
		nodes[nodePtr].Next = head.Next
		head.Next = nodes[nodePtr]

		head = next
		nodePtr--
		next = head.Next
	}
}
