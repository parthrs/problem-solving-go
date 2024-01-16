package arrays

func KidsWithCandies(candies []int, extraCandies int) (result []bool) {
	maxCandies := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}

	result = make([]bool, len(candies))

	for i, num := range candies {
		if extraCandies+num >= maxCandies {
			result[i] = true
		}
	}
	return
}
