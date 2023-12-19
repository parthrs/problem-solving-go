package stacks

import (
	"strconv"
	"unicode"
)

/*
This implements uses a stack to hold 3 elems for each combo
of - a num, '[' and set of characters that are being read in.
For every ']' those elems are popped, processed and result
added back to stack.
*/

func multiplyStrings(iStr string, s string) (res string) {
	if s == "" {
		return
	}
	i, _ := strconv.Atoi(iStr)
	for i > 0 {
		res += s
		i--
	}
	return
}

func stringify(i int) string {
	return strconv.Itoa(i)
}

func DecodeString(s string) string {
	stack := []string{""} // Our current "final string" is empty

	// Start processing
	num := 0
	for _, ch := range s {
		if unicode.IsDigit(ch) {
			num = num*10 + int(ch-'0') // Trick to accumulate number in case multiple digits
		} else if ch == '[' { // Finished processing number here
			stack = append(stack, stringify(num), "") // We expect an incoming start of str, so in prep send ""
			num = 0                                   // Also we are done processing num
		} else if ch == ']' {
			decoded, n, currStr := stack[len(stack)-3], stack[len(stack)-2], stack[len(stack)-1]
			decoded += multiplyStrings(n, currStr)
			stack = stack[:len(stack)-2]
			stack[len(stack)-1] = decoded
		} else {
			stack[len(stack)-1] += string(ch)
		}
	}

	return stack[0]
}
