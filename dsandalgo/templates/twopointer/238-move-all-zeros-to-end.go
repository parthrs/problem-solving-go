package twopointer

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Input:  [0,1,0,3,12]
Output: [1,3,12,0,0]

Input:  [2,1,0,0,3,2,4,0]
Output: [2,1,3,2,4,0,0,0]
*/

func MoveZerosToEnd(nums []int) []int {
	left, right := 0, 0
	for right < len(nums) {
		for nums[right] == 0 {
			right++
			if right >= len(nums) {
				return nums
			}
		}
		if left != right {
			t := nums[left]
			nums[left] = nums[right]
			nums[right] = t
		}
		left++
		right++
	}
	return nums
}
