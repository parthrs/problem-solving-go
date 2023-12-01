package main

func max(i, j int32) int32 {
	if i > j {
		return i
	}
	return j
}

func LargestSum(inList []int32) int32 {
	// While looping over the input list, maintain the max sum
	// and the current continous sum
	maxSum := inList[0]
	currentSum := inList[0]

	// Logic: At the end of every iteration, take the max of the current (continous) sum and max sum
	//        If adding the current sum (sum of the ints thus far) reduces the current num,
	//        make the num as the current sum and proceed.
	for _, num := range inList[1:] {
		if num > num+currentSum {
			currentSum = num
		}
		currentSum += num
		maxSum = max(maxSum, currentSum)
	}
	return maxSum
}
