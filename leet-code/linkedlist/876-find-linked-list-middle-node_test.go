package linkedlist

import (
	"testing"

	. "github.com/parthrs/problem-solving-go/pkg"
)

func TestMiddleNode(t *testing.T) {
	SinglyList := NewSinglyList[int]()
	for i := 5; i >= 1; i-- {
		SinglyList.AddNodeHead(i)
	}
	middleNode := MiddleNode(SinglyList.Head)
	if middleNode.Value != 3 {
		t.Errorf("Expected 3 as the middle node, got %d", middleNode.Value)
	}
	SinglyList.AddNodeHead(0)
	middleNode = MiddleNode(SinglyList.Head)
	if middleNode.Value != 3 {
		t.Errorf("Expected 3 as the middle node, got %d", middleNode.Value)
	}
	SinglyList.AddNodeHead(-1)
	middleNode = MiddleNode(SinglyList.Head)
	if middleNode.Value != 2 {
		t.Errorf("Expected 4 as the middle node, got %d", middleNode.Value)
	}
}
