package arrays

import "testing"

func TestEuclideanGcd(t *testing.T) {
	result, _ := EuclideanGcd(48, 18)
	if result != 6 {
		t.Errorf("Expected 6, got %d", result)
	}
}

func TestGcdOfStrings(t *testing.T) {
	inputs := [][]string{
		{"ABCDEF", "ABC"},
		{"TAUXXTAUXXTAUXXTAUXX", "TAUXX"},
	}

	expecteds := []string{
		"",
		"TAUXX",
	}

	for i := range inputs {
		if result := GcdOfStrings(inputs[i][0], inputs[i][1]); result != expecteds[i] {
			t.Errorf("Expected %v, Got %v", expecteds[i], result)
		}
	}
}
