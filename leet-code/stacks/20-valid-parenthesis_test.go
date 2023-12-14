package stacks

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	inputs := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"(])",
	}

	expecteds := []bool{
		true,
		true,
		false,
		false,
		true,
		false,
	}

	for i := range inputs {
		if result := IsValid(inputs[i]); result != expecteds[i] {
			t.Errorf("Expected %v, got %v", expecteds[i], result)
		}
	}
}
