package twopointer

import (
	"reflect"
	"testing"
)

func TestMoveZerosToEnd(t *testing.T) {
	inputNums := [][]int{
		{0, 1, 0, 3, 12},
	}
	expectedNums := [][]int{
		{1, 3, 12, 0, 0},
	}
	for i := range inputNums {
		result := MoveZerosToEnd(inputNums[i])
		if equal := reflect.DeepEqual(result, expectedNums[i]); !equal {
			t.Errorf("Expected: %v, Got: %v\n", expectedNums[i], result)
		}
	}
}
