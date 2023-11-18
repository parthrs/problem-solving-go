package twosum

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]
*/

// FindTwoSum returns the indices of the elements in elems
// that add up to k
// Regardless of the type of the input, indices are always
// type int
func FindTwoSum[T int32 | int64 | float32 | float64](elems []T, k T) (indices []int) {
	// Ensure there are atleast two elements in input
	if len(elems) < 2 {
		return
	}
	// Map to keep track of seen elems
	// key -> elem
	// value -> index(elem)
	seenMap := map[T]int{}
	// Loop over input and check if the
	// difference is in seen map
	for i, val := range elems {
		lookup := k - val
		if index, found := seenMap[lookup]; found {
			indices = []int{index, i}
			return
		}
		seenMap[val] = i
	}
	return
}
