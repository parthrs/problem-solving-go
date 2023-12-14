package linkedlist

import (
	. "github.com/parthrs/problem-solving-go/pkg"
)

/*
https://en.wikipedia.org/wiki/Cycle_detection#Tortoise_and_hare
The crux of this method uses Robert Floyd's (which as per wikipedia
should not be attributed to him) slow and fast pointer approach.
Same method to find if the linked list is looped.

The fast move 2x faster than slow, hence when fast is done,
slow is smack in the middle of the list.

Since we need to return head, we create a backup called
dummyHead.

Prev is needed because in singly list we don't have a prev
pointer, and we can't remove the middle node (slow) if we
can't change the prev node's pointer.

There are some smart approaches that:
- start slow, fast and prev to the new node
  before head, and,
- check for fast.Next && fast.Next.Next (crucial for this
  approach to work)
this way, slow points to one node behind of middle.
*/

func DeleteMiddle(head *SinglyNode[int]) *SinglyNode[int] {
	if head.Next == nil {
		return nil
	}
	prev := &SinglyNode[int]{
		Value: 0,
		Next:  head,
	}
	dummyHead, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev = prev.Next
		slow = slow.Next
		fast = fast.Next.Next
	}
	// truly remove the node
	// by removing it referencing the next
	// node
	prev.Next = slow.Next
	slow.Next = nil
	return dummyHead
}
