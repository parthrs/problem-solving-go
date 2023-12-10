package stacks

func RemoveStars(s string) string {
	stack := []rune{}
	for _, ch := range s {
		if ch == '*' {
			// pop prev char
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
