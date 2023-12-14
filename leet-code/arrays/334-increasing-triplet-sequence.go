package arrays

import "math"

func IncreasingTriplet(nums []int) bool {
	first, second := math.MaxInt, math.MaxInt
	for _, n := range nums {
		if n <= first { // If < is used, this if cond might skip and set it to second
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}
	return false
}
