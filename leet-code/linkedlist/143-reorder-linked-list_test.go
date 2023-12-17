package linkedlist

import (
	"testing"

	"github.com/parthrs/problem-solving-go/pkg"
)

func TestReorderList(t *testing.T) {
	SinglyList := pkg.NewSinglyList[int]()
	for i := 5; i >= 1; i-- {
		SinglyList.AddNodeHead(i)
	}
	expected := []int{1, 5, 2, 4, 3}
	ReorderList(SinglyList.Head)

	head := SinglyList.Head
	for _, val := range expected {
		if val != head.Value {
			t.Errorf("Reorder list failed, expected %d, got %d", val, head.Value)
		}
		head = head.Next
	}
}
