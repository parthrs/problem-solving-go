package arrays

func ProductExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	// Calculate products L to R
	// Start from index 1 because, there no
	// product to left of first elem
	i := 1
	runningProduct := nums[0]
	for i < len(nums) {
		products[i] = runningProduct
		runningProduct *= nums[i]
		i++
	}

	// Calculate products from R to L
	// Use the same existing array of earlier
	// products
	i = len(nums) - 2
	runningProduct = nums[len(nums)-1]
	for i >= 1 {
		products[i] *= runningProduct
		runningProduct *= nums[i]
		i--
	}
	// Since we skipped the first element
	products[0] = runningProduct
	return products
}
