package twopointer

/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
-- haystack and needle consist of only lowercase English characters.

Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
*/

func StrStr(haystack string, needle string) int {
	// This ensures both i and j are tracking the same
	// char, if there's no match
	// i -> track haystack chars
	// j -> track needle chars, hence get started
	// only when a match is there
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) { // Return haystack index on first full match of needle
				return i
			}
			if i+j == len(haystack) { // Needle is a part of haystack, if combined len is larger means either have not matched
				return -1
			}
			if needle[j] != haystack[i+j] { // i is rooted in haystack, and j moves, so see if it has run out
				break
			}
		}
	}
}
