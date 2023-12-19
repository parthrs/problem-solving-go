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
