package stacks

import (
	"testing"
)

func TestDecodeString(t *testing.T) {
	inputs := []string{
		"3[a]2[bc]",
		"3[a2[c]]",
		"2[abc]3[cd]ef",
	}
	expected := []string{
		"aaabcbc",
		"accaccacc",
		"abcabccdcdcdef",
	}
	for i, input := range inputs {
		if result := DecodeString(input); result != expected[i] {
			t.Errorf("Expected: %s, Got: %s", expected[i], result)
		}
	}
}
