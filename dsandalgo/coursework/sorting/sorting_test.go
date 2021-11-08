package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	if res := BubbleSort([]int{8, 3, 2, 1, 5, 7, 7, 0}); !reflect.DeepEqual(res, []int{0, 1, 2, 3, 5, 7, 7, 8}) {
		t.Errorf("BubbleSort test failed (got %v\n)", res)
	}
	if res := BubbleSort([]int{0, 1, 2, 3, 5, 7, 7, 8}); !reflect.DeepEqual(res, []int{0, 1, 2, 3, 5, 7, 7, 8}) {
		t.Errorf("BubbleSort test failed (got %v\n)", res)
	}
	if res := BubbleSort([]int{7, 9, 100, 11, -5, -7, 0, 10}); !reflect.DeepEqual(res, []int{-7, -5, 0, 7, 9, 10, 11, 100}) {
		t.Errorf("BubbleSort test failed (got %v\n)", res)
	}
}

func TestSelectionSort(t *testing.T) {
	if res := SelectionSort([]int{8, 3, 2, 1, 5, 7, 7, 0}); !reflect.DeepEqual(res, []int{0, 1, 2, 3, 5, 7, 7, 8}) {
		t.Errorf("SelectionSort test failed (got %v\n)", res)
	}
	if res := SelectionSort([]int{0, 1, 2, 3, 5, 7, 7, 8}); !reflect.DeepEqual(res, []int{0, 1, 2, 3, 5, 7, 7, 8}) {
		t.Errorf("SelectionSort test failed (got %v\n)", res)
	}
	if res := SelectionSort([]int{7, 9, 100, 11, -5, -7, 0, 10}); !reflect.DeepEqual(res, []int{-7, -5, 0, 7, 9, 10, 11, 100}) {
		t.Errorf("SelectionSort test failed (got %v\n)", res)
	}
}

func TestInsertionSort(t *testing.T) {
	if res := InsertionSort([]int{8, 3, 2, 1, 5, 7, 7, 0}); !reflect.DeepEqual(res, []int{0, 1, 2, 3, 5, 7, 7, 8}) {
		t.Errorf("InsertionSort test failed (got %v\n)", res)
	}
	if res := InsertionSort([]int{0, 1, 2, 3, 5, 7, 7, 8}); !reflect.DeepEqual(res, []int{0, 1, 2, 3, 5, 7, 7, 8}) {
		t.Errorf("InsertionSort test failed (got %v\n)", res)
	}
	if res := InsertionSort([]int{7, 9, 100, 11, -5, -7, 0, 10}); !reflect.DeepEqual(res, []int{-7, -5, 0, 7, 9, 10, 11, 100}) {
		t.Errorf("InsertionSort test failed (got %v\n)", res)
	}
}
