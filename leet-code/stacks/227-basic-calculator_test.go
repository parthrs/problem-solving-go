package stacks

import "testing"

func TestCalculate(t *testing.T) {
	expressions := []string{
		"3+2*2",
		"3*2+2",
		" 3/2 ",
		" 3+5 / 2 ",
		"42",
	}
	expected := []int{
		7,
		8,
		1,
		5,
		42,
	}

	for i := range expressions {
		if result := Calculate(expressions[i]); result != expected[i] {
			t.Errorf("For %s, expected %d got %d", expressions[i], expected[i], result)
		}
	}
}
