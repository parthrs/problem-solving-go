package linkedlist

import (
	"github.com/parthrs/problem-solving-go/pkg"
)

/*
Solve lc-206 and lc-143 first
Time & Space: O(n)
*/

func PairSum(head *pkg.SinglyNode[int]) int {
	vals := []int{}
	for head != nil {
		vals = append(vals, head.Value)
		head = head.Next
	}
	left, right := 0, len(vals)-1
	maxSum := 0
	for left < right {
		if result := vals[left] + vals[right]; result > maxSum {
			maxSum = result
		}
		left++
		right--
	}
	return maxSum
}
