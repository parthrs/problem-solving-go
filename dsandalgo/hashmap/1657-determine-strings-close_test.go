package hashmap

import (
	"testing"
)

func TestCloseStrings(t *testing.T) {
	if result := CloseStrings("aaabbbbccddeeeeefffff", "aaaaabbcccdddeeeeffff"); result {
		t.Errorf("Expected false, got true")
	}
}
