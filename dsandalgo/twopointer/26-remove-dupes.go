package twopointer

/*
Given a sorted array nums, remove the duplicates in-place such that each element appears only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Given nums = [1,1,2]
Len = 2

Given nums = [0,0,1,1,1,2,2,3,3,4]
Len = 5

Given nums = [0,1,2,2,3,3,4]
Len = 5

*/

func DeDuplicateArray(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	index := 0                 // Track the "current elem", or where to insert the new elem that finder finds
	for finder := range nums { // Find a new elem, thats different from what index is pointing
		if nums[index] == nums[finder] {
			continue
		}
		// Finder found a new elem, only then we arrive here
		index++
		nums[index] = nums[finder] // I know this overwrites, haven't found a solution that doesn't
	}
	return index + 1
}
