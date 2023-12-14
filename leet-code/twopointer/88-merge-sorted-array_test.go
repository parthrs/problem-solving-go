package twopointer

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	inputNums1 := [][]int{
		{1, 2, 3, 0, 0, 0},
		{1, 2, 3, 4, 0, 0, 0},
		{0},
		{1, 2, 0, 0, 0},
	}
	inputNums2 := [][]int{
		{2, 5, 6},
		{2, 5, 6},
		{1},
		{2, 5, 6},
	}
	expected := [][]int{
		{1, 2, 2, 3, 5, 6},
		{1, 2, 2, 3, 4, 5, 6},
		{1},
		{1, 2, 2, 5, 6},
	}
	inputM := []int{3, 4, 0, 2}
	inputN := []int{3, 3, 1, 3}
	for i := range inputNums1 {
		result := Merge(inputNums1[i], inputM[i], inputNums2[i], inputN[i])
		if equal := reflect.DeepEqual(result, expected[i]); !equal {
			t.Errorf("Expected: %v, Got: %v", expected[i], result)
		}
	}
}
