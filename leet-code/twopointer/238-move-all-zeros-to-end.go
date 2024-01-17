package twopointer

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Input:  [0,1,0,3,12]
Output: [1,3,12,0,0]

Input:  [2,1,0,0,3,2,4,0]
Output: [2,1,3,2,4,0,0,0]
*/

// MoveZerosToEnd gathers all zeros and moves
// them forward, collecting them.
// To maintain order we start both ptrs from LHS.
// One ptr moves rightward and finds non-zero elems,
// if found, it swaps with elem pointed by second ptr.
// Second ptr only moves if swap was done.
// What happens if the second ptr is pointing to a non-zero elem?
// A swap happens in place, and it is moved rightward.
// Ultimately both ptrs keep pace, until the first finds a zero,
// no swap happens till it finds a non-zero elem and it keeps moving
// rightward.
func MoveZerosToEnd(nums []int) []int {
	index := 0
	for finder := range nums {
		if nums[finder] != 0 { // If nums[finder] == 0, index stays there
			//if index != finder { // Avoids unnecessary swap if both point to same elem
			nums[index], nums[finder] = nums[finder], nums[index]
			//}
			index++
		}
	}
	return nums
}
