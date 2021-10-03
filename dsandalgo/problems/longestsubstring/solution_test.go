package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	inp := "abcabcbb"
	out := 3
	if result := LengthOfLongestSubstring(inp); result != out {
		t.Errorf("%d != %d for %s\n", result, out, inp)
	}

	inp = "bbbbb"
	out = 1
	if result := LengthOfLongestSubstring(inp); result != out {
		t.Errorf("%d != %d for %s\n", result, out, inp)
	}

	inp = "pwwkew"
	out = 3
	if result := LengthOfLongestSubstring(inp); result != out {
		t.Errorf("%d != %d for %s\n", result, out, inp)
	}

	inp = ""
	out = 0
	if result := LengthOfLongestSubstring(inp); result != out {
		t.Errorf("%d != %d for %s\n", result, out, inp)
	}

	inp = "pww"
	out = 2
	if result := LengthOfLongestSubstring(inp); result != out {
		t.Errorf("%d != %d for %s\n", result, out, inp)
	}

	inp = "aab"
	out = 2
	if result := LengthOfLongestSubstring(inp); result != out {
		t.Errorf("%d != %d for %s\n", result, out, inp)
	}

	inp = "dvdf"
	out = 3
	if result := LengthOfLongestSubstring(inp); result != out {
		t.Errorf("%d != %d for %s\n", result, out, inp)
	}

	inp = "ohvhjdml"
	out = 6
	if result := LengthOfLongestSubstring(inp); result != out {
		t.Errorf("%d != %d for %s\n", result, out, inp)
	}
}
