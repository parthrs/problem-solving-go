package twopointer

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := []string{
		"A man, a plan, a canal: Panama",
		"race a car",
		"",
	}

	expected := []bool{
		true,
		false,
		true,
	}

	for i := range input {
		result := IsPalindrome(input[i])
		if result != expected[i] {
			t.Errorf("Input: %v, Expected: %v, Got: %v", input[i], expected[i], result)
		}
	}
}
