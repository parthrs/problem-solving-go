package twopointer

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	inputHaystack := []string{
		"sadbutsad",
		"leetcode",
		"hello",
		"a",
	}
	inputNeedle := []string{
		"sad",
		"leeto",
		"ll",
		"a",
	}
	expected := []int{0, -1, 2, 0}

	for i := range inputHaystack {
		result := StrStr(inputHaystack[i], inputNeedle[i])
		if equal := result == expected[i]; !equal {
			t.Errorf("Needle: %s, Haystack: %s, Expected: %v, Got:%v\n", inputNeedle[i], inputHaystack[i], expected[i], result)
		}
	}
}
