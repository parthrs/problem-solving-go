package twopointer

import (
	"testing"
)

func TestReverse(t *testing.T) {
	inputs := []string{
		"the sky is blue",
		"  hello world!  ",
		"a good   example",
	}

	/*
	   0        9   13 16
	   thesarus sky is blue
	*/

	expecteds := []string{
		"blue is sky the",
		"world! hello",
		"example good a",
	}

	for i := range inputs {
		result := Reverse(inputs[i])
		if result != expecteds[i] {
			t.Errorf("Input: %s, Expected: %s, Got: %s", inputs[i], expecteds[i], result)
		}
	}
}
