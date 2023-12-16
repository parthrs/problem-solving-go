package stacks

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	inputs := [][]int{
		{5, 10, -5},
		{8, -8},
		{10, 2, -5},
		{10, 8, -2, 5, -8, 12, 14, -3},
		{1, -1, -2, -2},
	}
	expected := [][]int{
		{5, 10},
		{},
		{10},
		{10, 12, 14},
		{-2, -2},
	}

	for i := range inputs {
		if result := AsteroidCollision(inputs[i]); !reflect.DeepEqual(result, expected[i]) {
			t.Errorf("For input %v, expected: %v, got: %v", inputs[i], expected[i], result)
		}
	}
}
