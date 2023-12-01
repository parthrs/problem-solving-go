package twopointer

import (
	"strings"
)

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
For example,

"A man, a plan, a canal: Panama" is a palindrome.
"race a car" is not a palindrome.


: Empty strings allowed? Consider empty string as a valid palindrome
: Only alphanumeric? Yes

*/

func IsPalindrome(s string) (palindrome bool) {
	if len(s) < 1 {
		return true
	}

	s = strings.ToLower(s)
	//a := "a"[0]
	//z := "z"[0]

	left, right := 0, len(s)-1

	for right > left {
		for !isChar(s[right]) {
			right--
		}
		for !isChar(s[left]) {
			left++
		}
		if s[left] != s[right] {
			return
		}
		left++
		right--
	}

	return true
}

func isChar(b byte) bool {
	return ((b >= 'a' && b <= 'z') || (b >= '0' && b <= '9'))
}
