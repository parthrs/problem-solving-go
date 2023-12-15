package hashmap

import "sort"

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

	count1 := []int{}
	count2 := []int{}
	for _, v := range characMap1 {
		count1 = append(count1, v)
	}
	for _, v := range characMap2 {
		count2 = append(count2, v)
	}

	sort.Ints(count1)
	sort.Ints(count2)

	for i := range count1 {
		if count1[i] != count2[i] {
			return false
		}
	}
	return true
}
