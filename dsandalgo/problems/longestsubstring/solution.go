package main

import (
	"fmt"
)

/*
Given a string s, find the length of the longest substring without repeating characters.
Clarifying question: The repeating characters can be non-subsequent or non-continous?
*/

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getPos(s string, str string) int {
	for i, j := range str {
		if s == string(j) {
			return i
		}
	}
	return -1
}

func LengthOfLongestSubstring(s string) int {
	longestLen := 0
	seenMap := make(map[string]int)
	prevChar := ""
	subStr := ""

	for _, i := range s {
		if _, lookup := seenMap[string(i)]; lookup {
			if prevChar == string(i) {
				subStr = ""
				seenMap = make(map[string]int)
			} else {
				subStr = subStr[getPos(string(i), subStr)+1:]
			}
		}
		seenMap[string(i)] = 1
		subStr += string(i)
		longestLen = maxInt(longestLen, len(subStr))
		prevChar = string(i)
	}
	return longestLen
}

func main() {
	fmt.Println(LengthOfLongestSubstring("ohvhjdml"))
}
