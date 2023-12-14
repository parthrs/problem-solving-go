package twopointer

/*
Given n non-negative integers a1, a2, …, an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i , 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note : You may not slant the container and n is at least 2.

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

/*
To maximize the area, after calculating the trapped water at each step and comparing it to the maximum we've seen so far,
we move the pointer at the shorter line towards the other pointer. This is because keeping the pointer at the taller line stationary and
moving the shorter one might lead us to find a taller line and thus a larger area. There's no advantage in moving the taller pointer first,
as it would only reduce the potential width without guaranteeing a taller line to increase height. We repeat this process of calculating,
updating the maximum water area, and moving the shorter line pointer towards the other pointer until the two pointers meet,
at which point we've considered every possible container and the maximum stored water has been found.
*/

// maxArea calculates the area by comparing
// two heights, and moving the lesser one,
// to the next one, in search of a new high
func MaxArea(height []int) (area int) {
	left := 0
	right := len(height) - 1
	for left < right {
		area = getMax(area, (right-left)*getMin(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return
}

// Need min to take the lesser
// height of two lines to calculate
// volume
func getMin(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

// Needed to maintain a running
// max area on each iteration
func getMax(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
