package arrays

func MergeAlternately(word1 string, word2 string) (result string) {
	i := 0
	for i < len(word1) || i < len(word2) {
		if i < len(word1) {
			result += string(word1[i])
		}
		if i < len(word2) {
			result += string(word2[i])
		}
		i++
	}
	return
}
