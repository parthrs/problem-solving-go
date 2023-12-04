package stringrelated

import (
	"fmt"
)

/*
Given a set of words, compress the words to a smaller version.
The idea is to use the number of characters between the first
and the last, along with the first and last characs.
"apple" -> "a3e"
"kubernetes" -> k8s

The output should be unique, unambiguous (and a bunch of other things).
Questions:
- Can there be numbers? A: No. Only a-z (lowercase alphabets)
- What if there's zero len? A: Len only greater than 3.
- What if the shortned version is same for multiple words? (Also ask, can there be more than 2?)
  For this, use the following algorithm:
		- Keep adding one more letter from the left to the shortened
		  version, until you get a unique value
	e.g.
	"boggle" -> "bo3e" -> "bog2e"
	"bobble" -> "bo3e" -> "bob2e"
- What if the input contains all same characs?
Note: Here, I didn't get the "unambigous" part down.
Its not enough that you resolve the word that you find conflict,
you also solve all the conflicted words.
Also, in the interview, I was overwhelmed to not think
that there can be more than 2 clashes, and that way you can
use a map(key, []string).
*/

// ShortenString slices the input string s to return its shorter
// version using the logic:
// The startPtr instructs what character index from left should
// be preserved, the encoding is produced by counting the characters
// from index startPtr + 1.
func ShortenString(s string, startPtr int) string {
	return fmt.Sprintf("%s%d%s", string(s[:startPtr+1]), len(s)-(2+startPtr), string(s[len(s)-1]))
}

// StringCompression returns a clash free compress-encoding of a slice of
// strings. It resolves clashes by recursively calling self on a smaller
// input (i.e. the clash part).
func StringCompression(input []string, startPtr int) map[string][]string {
	// First "sweep", gather all the compressed
	// versions of the input
	// Here start with index 0 (provided by input)
	wordMap := map[string][]string{}
	for _, word := range input {
		shortened := ShortenString(word, startPtr)
		wordMap[shortened] = append(wordMap[shortened], word)
	}

	// Populate a second map in case there's a clash
	wordClash := map[string][]string{}
	for k, v := range wordMap {
		if len(v) > 1 {
			wordClash[k] = []string{}
			wordClash[k] = append(wordClash[k], v...) // This is done to copy (using append) the result slice
		}
	}

	// If there are clashes, for each clash:
	// - Call self on the list of words that caused clash
	// - Offset starting index by 1
	// - Update the main map with result
	// - Remove the clash key (and its value)
	for k, v := range wordClash {
		resultMap := StringCompression(v, startPtr+1)
		for key, val := range resultMap {
			wordMap[key] = []string{}
			wordMap[key] = append(wordMap[key], val...)
		}
		delete(wordMap, k)
	}

	// Finally return the map
	return wordMap
}
