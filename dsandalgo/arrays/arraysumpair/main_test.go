package main

import "testing"

func TestArraySumPair(t *testing.T) {
	if res := len(ArraySumPair([]int32{1, 5, 6, 8, 2, -2, 0, 9}, 7)); res != 3 {
		t.Errorf("Expected: two pairs but got %v\n", res)
	}
	if res := len(ArraySumPair([]int32{1, 9, 2, 8, 3, 7, 4, 6, 5, 5, 13, 14, 11, 13, -1}, 10)); res != 6 {
		t.Errorf("Expected: six pairs but got %v\n", res)
	}
	if res := len(ArraySumPair([]int32{1, 2, 3, 1}, 3)); res != 1 {
		t.Errorf("Expected: one pair but got %v\n", res)
	}
	if res := len(ArraySumPair([]int32{1, 3, 2, 2}, 4)); res != 2 {
		t.Errorf("Expected: two pairs but got %v\n", res)
	}
}
