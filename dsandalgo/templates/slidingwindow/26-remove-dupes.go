package slidingwindow

/*
Given a sorted array nums, remove the duplicates in-place such that each element appears only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Given nums = [1,1,2]
Len = 2

Given nums = [0,0,1,1,1,2,2,3,3,4]
Len = 5

*/

func DeDuplicateArray(nums []int) (length int) {
	left, right := 0, 0
	for right < len(nums) {
		for nums[left] == nums[right] {
			right++
			if right >= len(nums) {
				return
			}
		}
		left++
		nums[left] = nums[right]
		length = left + 1
	}
	return
}
