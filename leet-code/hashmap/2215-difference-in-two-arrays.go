package hashmap

// Simple straightforward solution
func FindDifference(nums1 []int, nums2 []int) (result [][]int) {
	elemMap1 := map[int]bool{}
	for _, num := range nums1 {
		elemMap1[num] = true
	}
	elemMap2 := map[int]bool{}
	for _, num := range nums2 {
		elemMap2[num] = true
	}
	result = make([][]int, 2)
	for k, _ := range elemMap1 {
		if _, found := elemMap2[k]; found {
			delete(elemMap2, k)
		} else {
			result[0] = append(result[0], k)
		}
	}
	for k := range elemMap2 {
		result[1] = append(result[1], k)
	}

	return
}

// Using less memory then two maps
// Commented since not really hashmap
// Here the idea is to use a single array to track things (could be done with single map too)
// Loop over nums 1, mark found elems in tracker as 1 (1 means that elem was found in loop 1)
// Loop over nums 2, now if that index val is 0, means it wasn't found in loop 1 -> append result
//
//	also, set all elems seen in this loop to 2 (means that elem was found in loop 2)
//
// Finally, again loop over nums1, at this point we've marked elems as 1 in loop 1 and 2 in loop 2,
// so if we find elems that are still 1, then they were only seen in loop 1 so add them to result, and,
// to avoid ambiguity, also make them 0.
// func FindDifference(nums1 []int, nums2 []int) (result [][]int) {
// 	numTracker := make([]int, 2001, 2001)
// 	for _, num := range nums1 {
// 		numTracker[num+1000] = 1
// 	}
// 	result = make([][]int, 2, 2)
// 	for _, num := range nums2 {
// 		if numTracker[num+1000] == 0 {
// 			result[1] = append(result[1], num)
// 		}
// 		numTracker[num+1000] = 2
// 	}
// 	for _, num := range nums1 {
// 		if numTracker[num+1000] == 1 {
// 			result[0] = append(result[0], num)
// 			numTracker[num+1000] = 0
// 		}
// 	}
// 	return
// }
