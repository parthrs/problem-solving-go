package arrays

func KidsWithCandies(candies []int, extraCandies int) (result []bool) {
	maxCandies := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > maxCandies {
			maxCandies = candies[i]
		}
	}
	for _, num := range candies {
		maxCandy := false
		if extraCandies+num >= maxCandies {
			maxCandy = true
		}
		result = append(result, maxCandy)
	}
	return
}
