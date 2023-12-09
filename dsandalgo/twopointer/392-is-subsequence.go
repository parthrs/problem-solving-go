package twopointer

func IsSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	i := 0 // ptr to traverse s
	for p := 0; p < len(t); p++ {
		if t[p] == s[i] {
			i++
			if i == len(s) { // This ensures we return asap when s is fully traversed
				return true
			}
		}
	}
	return false
}
