package slidingwindow

import (
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	inputs := [][]int{
		{1, 12, -5, -6, 50, 3},
		{4},
		{5},
		{1},
	}

	// Every odd elem (i.e. 0.0) is dummy input so 'i'
	// can track both inputs and expected elems
	expected := []float64{12.75, 0.0, 5.00, 0.0}

	for i := 0; i < len(inputs); i += 2 {
		if result := FindMaxAverage(inputs[i], inputs[i+1][0]); result != expected[i] {
			t.Errorf("For input %v k=%v, expected: %v, got: %v", inputs[i], inputs[i+1][0], expected[i], result)
		}
	}
}
