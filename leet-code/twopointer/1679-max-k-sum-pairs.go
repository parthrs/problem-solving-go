package twopointer

import "sort"

// Same as the two sum implementation in this
// package

func MaxOperations(nums []int, k int) (count int) {
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
		if sortedSlice[left][0]+sortedSlice[right][0] == k {
			count++
			left++
			right--
		} else if sortedSlice[left][0]+sortedSlice[right][0] > k {
			right--
		} else {
			left++
		}
	}
	return
}
