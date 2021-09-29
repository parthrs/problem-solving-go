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

func linkedListToInt(l *sll.SinglyNode) int {
	sum := 0
	decimalMultiplier := 1

	if l.Value == 0 {
		return 0
	}

	for n := l; n != nil; {
		sum += (decimalMultiplier * n.Value)
		decimalMultiplier *= 10
		n = n.Next
	}
	return sum
}

func intToList(i int) *sll.SinglyNode {
	var prevNode *sll.SinglyNode = nil
	var headNode *sll.SinglyNode = nil

	if i == 0 {
		return sll.NewSinglyNode(i)
	}

	for i > 0 {
		_n := sll.NewSinglyNode(i % 10)
		if prevNode == nil {
			headNode = _n
		} else {
			prevNode.Next = _n
		}
		prevNode = _n
		i = i / 10
	}
	return headNode
}

func AddTwoNumbers(m *sll.SinglyNode, n *sll.SinglyNode) *sll.SinglyNode {
	return intToList(linkedListToInt(m) + linkedListToInt(n))
}
