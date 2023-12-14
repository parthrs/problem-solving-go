package linkedlist

import (
	. "github.com/parthrs/problem-solving-go/pkg"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
The brute force way is to reach the end, set head
to that and then traverse back. But that ends up us
traversing twice.

Instead traverse once, when you are on a node,
take a "backup" of the next pointer. Then set its next
to the prev pointer (which starts with nil). Set the prev
pointer to head (current node, so set prev = current node)
and finally set head to the next value (backed up).
*/

func ReverseList(head *SinglyNode[int]) *SinglyNode[int] {
	var prev *SinglyNode[int]
	for head != nil { // Start with catching empty Lists
		head.Next, prev, head = prev, head, head.Next
		// t := head.Next
		// head.Next = prev
		// prev = head
		// head = t
	}
	return prev
}
