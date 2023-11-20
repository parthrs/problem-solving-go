package twopointer

import (
	"testing"
)

func TestDeDuplicateArray(t *testing.T) {
	inputNums := [][]int{
		{1, 1, 2},
		{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		{0, 1, 2, 2, 3, 3, 4},
	}
	expected := []int{2, 5, 5}
	for i := range inputNums {
		result := DeDuplicateArray(inputNums[i])
		if result != expected[i] {
			t.Errorf("Expected: %v, Got: %v", expected[i], result)
		}
	}
}
