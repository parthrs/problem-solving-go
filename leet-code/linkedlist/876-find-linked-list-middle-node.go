package linkedlist

import (
	"github.com/parthrs/problem-solving-go/pkg"
)

/*
To-Do: What if we wish to know the last and the middle node?
Currently the last node points to nil.

Also To-Do: What are the various outcomes of starting slow and
fast at different positions - like slow at -1 and fast at 0.
*/
func MiddleNode(head *pkg.SinglyNode[int]) *pkg.SinglyNode[int] {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
