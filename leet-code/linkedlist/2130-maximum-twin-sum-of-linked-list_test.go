package linkedlist

import (
	"testing"

	"github.com/parthrs/problem-solving-go/pkg"
)

func TestPairSum(t *testing.T) {
	SinglyList := pkg.NewSinglyList[int]()
	for i := 1; i <= 5; i++ {
		SinglyList.AddNodeHead(i)
	}

	if result := PairSum(SinglyList.Head); result != 6 {
		t.Errorf("Reorder list failed, expected %d, got %d", 6, result)
	}

	SinglyList = pkg.NewSinglyList[int]()
	for _, val := range []int{3, 2, 2, 4} {
		SinglyList.AddNodeHead(val)
	}
	if result := PairSum(SinglyList.Head); result != 7 {
		t.Errorf("Reorder list failed, expected %d, got %d", 7, result)
	}
}
