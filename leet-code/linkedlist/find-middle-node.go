package linkedlist

import (
	"fmt"

	"github.com/parthrs/problem-solving-go/pkg"
)

/*
The tortoise and hare algorithm to detect cycle in a LL.
Apparently its misattributed to R. Floyd
https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_tortoise_and_hare
*/

func PrintMiddleAndLastNodes[T any](slow, fast *pkg.SinglyNode[T]) {
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	fmt.Printf("Slow: %v\n", slow)
	fmt.Printf("Fast: %v\n", fast)
}
