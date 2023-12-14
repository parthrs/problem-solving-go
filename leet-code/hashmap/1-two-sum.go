package hashmap

func TwoSum(nums []int, target int) []int {
	seenMap := map[int]int{}
	for index, i := range nums {
		seek := target - i
		if seekIndex, found := seenMap[seek]; found {
			return []int{seekIndex, index}
		}
		seenMap[i] = index
	}
	return []int{}
}
