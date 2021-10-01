package main

/*
https://leetcode.com/problems/add-two-numbers/
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807
*/

import (
	sll "github.com/parthrs/LetsGo/dsandalgo/pkg/singlylinkedlist"
)

func AddTwoNumbers(m *sll.SinglyNode, n *sll.SinglyNode) *sll.SinglyNode {
	carryForward := 0
	var prevNode *sll.SinglyNode
	var headNode *sll.SinglyNode

	for m != nil || n != nil {
		a := 0
		b := 0

		if m != nil {
			a = m.Value
		}

		if n != nil {
			b = n.Value
		}

		summation := a + b + carryForward
		carryForward = summation / 10
		if summation >= 10 {
			summation %= 10
		}
		_n := sll.NewSinglyNode(summation)
		if prevNode == nil {
			headNode = _n
		} else {
			prevNode.Next = _n
		}

		prevNode = _n
		if m != nil {
			m = m.Next
		}
		if n != nil {
			n = n.Next
		}
	}

	if carryForward != 0 {
		prevNode.Next = sll.NewSinglyNode(1)
	}

	return headNode
}
