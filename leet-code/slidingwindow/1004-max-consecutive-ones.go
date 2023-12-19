package slidingwindow

// Brute force approach
func LlongestOnes(nums []int, k int) int {
	left, maxCount, count, borrow := 0, 0, 0, k
	for _, elem := range nums {
		if elem == 1 {
			count++
		} else {
			if borrow <= 0 {
				for nums[left] != 0 {
					left++
					count--
				}
				left++
			} else {
				borrow--
				count++
			}
		}
		maxCount = findMax(maxCount, count)
	}
	return maxCount
}

// The window size, i.e. the distance between the
// left and the right pointers, is the max consecutive
// ones with flipping max k 0's
func LongestOnes(nums []int, k int) int {
	left := 0

	// Right pointer for the window
	// is greedily moved right every iteration
	// left is only moved up if we've exhausted
	// k (i.e. the ability to flip 0's to 1's)
	for _, elem := range nums {
		if elem == 0 {
			k--
		}
		// Each time k is negative, move left up
		// i.e. push the window forward and not grow
		// it.
		// If while moving the window, the number leaving
		// is 0, increase k.
		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
	}
	return len(nums) - left
}
