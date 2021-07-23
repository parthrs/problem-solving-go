package main

import "testing"

func TestArraySumPair(t *testing.T) {
	if res := len(ArraySumPair([]int32{1, 5, 6, 8, 2, -2, 0, 9}, 7)); res != 2 {
		t.Errorf("Expected: two pairs but got %v\n", res)
	}
}
