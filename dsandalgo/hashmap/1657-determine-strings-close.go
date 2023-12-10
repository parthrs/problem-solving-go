package hashmap

func CloseStrings(word1 string, word2 string) bool {
	// Ensure words have same lengths
	if len(word1) != len(word2) {
		return false
	}

	// Count occurences of characs
	// for each word
	characMap1 := map[rune]int{}
	characMap2 := map[rune]int{}

	for _, ch := range word1 {
		characMap1[ch] += 1
	}
	for _, ch := range word2 {
		characMap2[ch] += 1
	}

	for k, _ := range characMap1 {
		if _, found := characMap2[k]; !found {
			return false
		}
	}

	for k, _ := range characMap2 {
		if _, found := characMap1[k]; !found {
			return false
		}
	}

	countMap1 := map[int]rune{}
	countMap2 := map[int]rune{}

	for k, v := range characMap1 {
		countMap1[v] = k
	}
	for k, v := range characMap2 {
		countMap2[v] = k
	}

	for k, _ := range countMap1 {
		if _, found := countMap2[k]; !found {
			return false
		}
		delete(countMap2, k)
	}
	return len(countMap2) == 0
}
