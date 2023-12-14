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

func MiddleNode(head *SinglyNode[int]) *SinglyNode[int] {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
