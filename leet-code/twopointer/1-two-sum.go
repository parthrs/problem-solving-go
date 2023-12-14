package twopointer

import "sort"

// For the hashmap based implementation check that package
// This first sorts the slice in ascending order. Then initialize
// two pointers poles apart. Then compare their sum with target.
// If sum is more, reduce the vars by moving the right ptr in,
// otherwise opposite.
func TwoSum(nums []int, target int) []int {
	sortedSlice := [][]int{}
	for i, elem := range nums {
		sortedSlice = append(sortedSlice, []int{elem, i})
	}
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i][0] < sortedSlice[j][0]
	})

	left := 0
	right := len(nums) - 1
	for left < right {
		if sortedSlice[left][0]+sortedSlice[right][0] == target {
			return []int{sortedSlice[left][1], sortedSlice[right][1]}
		}
		if sortedSlice[left][0]+sortedSlice[right][0] > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}
