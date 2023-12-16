package linkedlist

import (
	"github.com/parthrs/problem-solving-go/pkg"
)

/*
Maintain two pointers for our list, one for the last
odd element and other for the last even element.
Starting with 1 and 2 respectively.
The element pointed by the last even element is
the one we need to "pull up", by placing it at
the end of the last odd element.

On each iteration, if there's an element after the
last even elem (i.e. potential candidate), we move it up.
*/
func OddEvenList(head *pkg.SinglyNode[int]) *pkg.SinglyNode[int] {
	// Empty or single node list
	if head == nil || head.Next == nil {
		return head
	}

	oddEnd, evenEnd := head, head.Next
	for evenEnd != nil && evenEnd.Next != nil {
		curr := evenEnd.Next

		evenEnd.Next = curr.Next
		evenEnd = evenEnd.Next
		curr.Next = oddEnd.Next
		oddEnd.Next = curr
		oddEnd = curr
	}

	return head
}
