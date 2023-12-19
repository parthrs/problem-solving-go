package slidingwindow

/*
Same logic as lc-643
*/

func MaxVowels(s string, k int) int {
	start, end := 0, k-1
	count := 0
	for i := start; i <= end; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	maxCount := count
	end++

	for end < len(s) {
		if isVowel(s[start]) {
			count -= 1
		}
		if isVowel(s[end]) {
			count += 1
		}
		start++
		end++
		maxCount = findMax(count, maxCount)
	}
	return maxCount
}

func isVowel(b byte) bool {
	ch := rune(b)
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	return false
}

func findMax(i, j int) int {
	if i > j {
		return i
	}
	return j
}
