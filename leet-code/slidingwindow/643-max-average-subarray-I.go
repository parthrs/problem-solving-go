package slidingwindow

func FindMaxAverage(nums []int, k int) float64 {
	start, end := 0, k-1
	runningSum := 0.0

	for i := start; i <= end; i++ {
		runningSum += float64(nums[i])
	}

	maxAvg := runningSum / float64(k)
	end++

	for end < len(nums) {
		runningSum -= float64(nums[start])
		start++
		runningSum += float64(nums[end])
		maxAvg = findMax(maxAvg, runningSum/float64(k))
		end++
	}
	return maxAvg
}

func findMax(i, j float64) float64 {
	if i < j {
		return j
	}
	return i
}
