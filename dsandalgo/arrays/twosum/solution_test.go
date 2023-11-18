package twosum

import (
	"reflect"
	"testing"
)

func TestFindTwoSum(t *testing.T) {
	// Type int32
	inputElems := [][]int32{
		{2, 7, 11, 15},
		{2, 7, 11, 15},
	}
	inputK := []int32{9, 26}
	output := [][]int{
		{0, 1},
		{2, 3},
	}
	for i := range inputElems {
		retVal := FindTwoSum(inputElems[i], inputK[i])
		if same := reflect.DeepEqual(retVal, output[i]); !same {
			t.Errorf("Expected %v, got %v\n", output[i], retVal)
		}
	}

	// Type float32
	inputElemsFloat := [][]float64{
		{2.6, 7.1, 11.9, 15.3},
	}
	inputKFloat := []float64{19.0}
	outputFloat := [][]int{
		{1, 2},
	}
	for i := range inputElemsFloat {
		retVal := FindTwoSum(inputElemsFloat[i], inputKFloat[i])
		if same := reflect.DeepEqual(retVal, outputFloat[i]); !same {
			t.Errorf("Expected %v, got %v\n", outputFloat[i], retVal)
		}
	}
}
