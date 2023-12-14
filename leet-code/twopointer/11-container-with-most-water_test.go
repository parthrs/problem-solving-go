package twopointer

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	input := [][]int{
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{1, 5, 4, 3},
		{3, 1, 2, 4, 5},
		{1, 1},
		{2, 3, 4, 5, 18, 17, 6},
	}
	expected := []int{49, 6, 12, 1, 17}

	for i := range input {
		result := MaxArea(input[i])
		if result != expected[i] {
			t.Errorf("Input: %v, Expected: %d, Got: %d", input[i], expected[i], result)
		}
	}
}
