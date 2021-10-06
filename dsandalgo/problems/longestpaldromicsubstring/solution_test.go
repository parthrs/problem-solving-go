package main

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	i := "a"
	o := "a"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}

	i = "ac"
	o = "a"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}

	i = "bb"
	o = "bb"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}

	i = "abcbaef"
	o = "abcba"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}

	i = "kweatopotaz"
	o = "atopota"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}

	i = "ccc"
	o = "ccc"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}

	i = "dddd"
	o = "dddd"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}

	i = "abba"
	o = "abba"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}

	i = "cbbd"
	o = "bb"
	if res := LongestPalindrome(i); res != o {
		t.Errorf("Expected %s got %s\n", o, res)
	}
}

/*
func TestIsPalindrome(t *testing.T) {
	i := "aca"
	o := true
	if res := IsPalindrome(i); res != o {
		t.Errorf("Expected %v got %v\n", o, res)
	}

	i = ""
	o = true
	if res := IsPalindrome(i); res != o {
		t.Errorf("Expected %v got %v\n", o, res)
	}

	i = "a"
	o = true
	if res := IsPalindrome(i); res != o {
		t.Errorf("Expected %v got %v\n", o, res)
	}

	i = "bb"
	o = true
	if res := IsPalindrome(i); res != o {
		t.Errorf("Expected %v got %v\n", o, res)
	}

	i = "abcba"
	o = true
	if res := IsPalindrome(i); res != o {
		t.Errorf("Expected %v got %v\n", o, res)
	}

	i = "atopota"
	o = true
	if res := IsPalindrome(i); res != o {
		t.Errorf("Expected %v got %v\n", o, res)
	}
}
*/
