package main

import (
	"strings"
)

func AnagramCheck(firstWord string, anagram string) bool {
	/*
		An anagram of a word is another word made by reshuffling the letters, such that no letter is left unused and no new letters
		are added.
		https://en.wikipedia.org/wiki/Anagram
	*/

	/*
		Self-notes/TOL:
		- Check 0: Length of both the words
		- Approach 1: Use a map and then update it while iterating over both the words
		- Approach 2: Sort both the words and then loop over both simultaneously, all letter should match
	*/

	// This only works if we get rid of the blank spaces first
	// Compare lengths of the two words
	//if len(firstWord) != len(anagram) {
	//	fmt.Printf("%s is not an anagram of %s\n", anagram, firstWord)
	//	os.Exit(0)
	//}

	// Approach 1
	letterMap := make(map[string]int)

	for _, letter := range anagram {
		let := strings.ToLower(string(letter))
		if let == " " {
			continue
		}
		letterMap[let] += 1
	}

	for _, letter := range firstWord {
		let := strings.ToLower(string(letter))
		if let == " " {
			continue
		}
		if _, ok := letterMap[let]; ok {
			letterMap[let] -= 1
		} else {
			return false
		}
	}

	for _, count := range letterMap {
		if count != 0 {
			return false
		}
	}
	return true
}
