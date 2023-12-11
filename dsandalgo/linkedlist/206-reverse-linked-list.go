package linkedlist

import (
	"github.com/parthrs/LetsGo/dsandalgo/pkg"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode pkg.SinglyNode[int]

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil { // Start with catching empty Lists
		head.Next, prev, head = prev, head, head.Next
		// t := head.Next
		// head.Next = prev
		// prev = head
		// head = t
	}
	return prev
}
