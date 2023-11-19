package twopointer

/*
Given an array nums and a value val, remove all instances of that value in-place and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

The order of elements can be changed. It doesn't matter what you leave beyond the new length.

Given nums = [3,2,2,3], val = 3
length = 2

Given nums = [0,1,2,2,3,0,4,2], val = 2
length = 5
*/

// My attempt. Passed tests.
// But found a better solution online
func MyRemoveSpecificElem(nums []int, k int) (length int) {
	for i := range nums {
		if nums[i] != k {
			continue
		}
		left, right := i, i
		length = i
		for nums[right] == k {
			right++
			if right >= len(nums) {
				return
			}
		}
		t := nums[left]
		nums[left] = nums[right]
		nums[right] = t
	}
	return
}

// The core logic is is to "bubble up" the unwanted elem
// to the top
// Essentially both pointers should move in tandem
// This is done by moving right by one
// Moving left by one (to right), only if elem right is pointing
// to is not k.
// If thats the case swap first, and then increment left.
// Repeat
func RemoveSpecificElem(nums []int, k int) (length int) {
	if len(nums) < 1 {
		return 0
	}
	left := 0
	for right := range nums {
		if nums[right] != k {
			if left != right {
				nums[left], nums[right] = nums[right], nums[left]
			}
			left++
		}
	}
	return left
}
