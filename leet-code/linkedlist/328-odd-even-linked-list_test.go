package linkedlist

import (
	"testing"

	"github.com/parthrs/problem-solving-go/pkg"
)

func TestOddEvenList1(t *testing.T) {
	singly := pkg.NewSinglyList[int]()
	for i := 5; i >= 1; i-- {
		singly.AddNodeHead(i)
	}
	head := OddEvenList(singly.Head)
	expected := []int{1, 3, 5, 2, 4}
	for _, val := range expected {
		if head.Value != val {
			t.Errorf("Expected %d, got %d", val, head.Value)
		}
		head = head.Next
	}
}

func TestOddEvenList2(t *testing.T) {
	singly := pkg.NewSinglyList[int]()
	input := []int{2, 1, 3, 5, 6, 4, 7}
	for i := len(input) - 1; i >= 0; i-- {
		singly.AddNodeHead(input[i])
	}

	head := OddEvenList(singly.Head)
	expected := []int{2, 3, 6, 7, 1, 5, 4}
	for _, val := range expected {
		if head.Value != val {
			t.Errorf("Expected %d, got %d", val, head.Value)
		}
		head = head.Next
	}
}
