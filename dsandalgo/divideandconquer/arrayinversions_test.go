package main

import (
	"reflect"
	"testing"
)

func TestDiscoverInversions(t *testing.T) {
	if res, cn := DiscoverInversions([]int{1, 3, 5, 2, 4, 6}); !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Array returned unsorted (got %v\n)", res)
	} else if cn != 3 {
		t.Errorf("Incorrect number of inversions (expected 3, got %v\n", cn)
	}

	if res, cn := DiscoverInversions([]int{6, 5, 4, 3, 2, 1}); !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6}) {
		t.Errorf("Array returned unsorted (got %v\n)", res)
	} else if cn != 15 {
		t.Errorf("Incorrect number of inversions (expected 3, got %v\n", cn)
	}
}
