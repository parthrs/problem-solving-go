package twopointer

import (
	"testing"
)

func TestRemoveSpecificElem(t *testing.T) {
	inputNums := [][]int{
		{3, 2, 2, 3},
		{0, 1, 2, 2, 3, 0, 4, 2},
	}
	inputK := []int{3, 2}
	output := []int{2, 5}
	for i := range inputNums {
		result := RemoveSpecificElem(inputNums[i], inputK[i])
		if result != output[i] {
			t.Errorf("Expected: %v, Got: %v\n", output[i], result)
		}
	}
}
