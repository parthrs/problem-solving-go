package misc

import (
	"reflect"
	"testing"
)

func TestShortenString(t *testing.T) {
	if result := ShortenString("apple", 0); result != "a3e" {
		t.Errorf("Expected: %v, Got: %v", "a3e", result)
	}
	if result := ShortenString("apple", 1); result != "ap2e" {
		t.Errorf("Expected: %v, Got: %v", "ap2e", result)
	}
	if result := ShortenString("boggle", 2); result != "bog2e" {
		t.Errorf("Expected: %v, Got: %v", "bog2e", result)
	}
}

func TestStringCompression(t *testing.T) {
	inputs := []string{
		"apple",
		"ankle",
		"kubernetes",
		"boggle",
		"bobble",
		"bottle",
	}
	expected := map[string][]string{
		"k8s":   {"kubernetes"},
		"ap2e":  {"apple"},
		"an2e":  {"ankle"},
		"bog2e": {"boggle"},
		"bob2e": {"bobble"},
		"bot2e": {"bottle"},
	}
	if result := StringCompression(inputs, 0); !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}
