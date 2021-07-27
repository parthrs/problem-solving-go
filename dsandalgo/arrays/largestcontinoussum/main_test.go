package main

import "testing"

func TestArraySumPair(t *testing.T) {
	if res := LargestSum([]int32{1, 2, -1, 3, 4, -1}); res != 9 {
		t.Errorf("Expected: 9 but got %v\n", res)
	}
	if res := LargestSum([]int32{1, 2, -1, 3, 4, 10, 10, -10, -1}); res != 29 {
		t.Errorf("Expected: 29 but got %v\n", res)
	}
	if res := LargestSum([]int32{1, -1}); res != 1 {
		t.Errorf("Expected: 1 but got %v\n", res)
	}
}
