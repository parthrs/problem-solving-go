package stacks

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets. Open brackets must be closed in the correct order. Note that an empty string is also considered valid.
*/

func IsValid(s string) bool {
	stack := make([]rune, 0)
	for _, ch := range s {
		if ch == '[' || ch == '(' || ch == '{' {
			stack = append(stack, ch)
		} else if len(stack) > 0 && ch == ']' && stack[len(stack)-1] == '[' || len(stack) > 0 && ch == '}' && stack[len(stack)-1] == '{' || len(stack) > 0 && ch == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
