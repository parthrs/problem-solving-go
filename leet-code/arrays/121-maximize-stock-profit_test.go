package arrays

import (
	"testing"
)

func TestMaximizeProfit(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3},
		{2, 1, 3},
		{3, 2, 1},
		{2, 1},
		{7, 1, 5, 3, 6, 4},
	}
	expected := []int{2, 2, 0, 0, 5}

	for i := range inputs {
		result := MaximizeProfit(inputs[i])
		if result != expected[i] {
			t.Errorf("Input: %v, Expected: %v, Got: %v", inputs[i], expected[i], result)
		}
	}
}
