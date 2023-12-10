package hashmap

// Use hashmap to count the times you've seen an elem
// while traversing and updating it, till sum is matched.
// If sum matched decrement

func MaxOperations(nums []int, k int) (count int) {
	seenMap := map[int]int{}
	for _, num := range nums {
		seek := k - num
		if counter, found := seenMap[seek]; found && counter > 0 {
			seenMap[seek] -= 1
			count += 1
		} else {
			seenMap[num] += 1 // You reach here when found and 0 or not found
		}
	}
	return count
}
