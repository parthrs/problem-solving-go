package main

import "fmt"

/*
Given a string s, return the longest palindromic substring in s.
*/

func longestStr(a string, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

/*
func IsPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	retVal := true
	upperCount := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if string(s[i]) != string(s[upperCount]) {
			retVal = false
			break
		}
		upperCount -= 1
	}
	return retVal
}
*/

func expandingPalindromeSearch(revPtr int, fwdPtr int, s string) string {
	for revPtr-1 >= 0 && fwdPtr+1 < len(s) {
		if string(s[fwdPtr+1]) != string(s[revPtr-1]) {
			break
		}
		fwdPtr += 1
		revPtr -= 1
	}

	return string(s[revPtr : fwdPtr+1])
}

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	maxPalindrome := string(s[0])
	currentPalindrome := string(s[0])
	count := 1

	for count < len(s) {

		if string(s[count]) == string(s[count-1]) {
			// Repeating characters
			currentPalindrome = expandingPalindromeSearch(count-1, count, s)
		} else if count >= 2 && string(s[count]) == string(s[count-2]) {
			// Mirror image of the word
			currentPalindrome = expandingPalindromeSearch(count-1, count-1, s)
		}

		count += 1
		maxPalindrome = longestStr(maxPalindrome, currentPalindrome)
	}
	return maxPalindrome
}

func main() {
	fmt.Println(LongestPalindrome("ccc"))
}
