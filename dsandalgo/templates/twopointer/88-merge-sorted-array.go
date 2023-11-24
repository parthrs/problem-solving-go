package twopointer

/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note :

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.

Input:
nums1 = [1,2,3,0,0,0], m = 3 (len m, cap m+n)
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]

==

nums1 = [1,2,3,4,0,0,0], m = 4 (len m, cap m+n)
nums2 = [2,5,6],         n = 3

Output: [1,2,2,3,4,5,6]

==

nums1 = [1,2,0,0,0], m = 2 (len m, cap m+n)
nums2 = [2,5,6],     n = 3

Output: [1,2,2,5,6]

==

[1,2,3,4,0,0,0]
*/

// nums1 will have "buffer" space at the end, equalling to
// len(nums2). So nums1 will be the result array of the merge.
// The core logic is to run a ptr on nums1, which starts from
// the last elem (of the buffer part).
// We then loop in reverse on both nums1 and nums2.
// And apply the 'merge sort' logic of comparing the elems,
// and writing the elem to the ptr that is greater.
// Finally, if n ptr hasn't run out (only when m is zero),
// we just copy each of its elems into nums1 (buffer only part).
func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	for p := m + n - 1; m > 0 && n > 0; p-- { // m and n use '>' since they are actual lens
		if nums1[m-1] > nums2[n-1] {
			nums1[p] = nums1[m-1]
			m--
		} else {
			nums1[p] = nums2[n-1]
			n--
		}
	}
	// if nums1 is empty
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n--
	}

	return nums1
}
