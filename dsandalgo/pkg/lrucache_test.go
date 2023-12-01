package pkg

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(3)
	c.put(1, 100)
	c.put(2, 200)
	c.put(3, 300)
	if c.get(4) != -1 {
		t.Error("Incorrect return value when elem not present")
	}
	if c.get(1) != 100 {
		t.Error("Method get failed to get correct value")
	}
	c.put(4, 400) // 2 should get evicted
	if c.get(2) != -1 {
		t.Error("Failed to evict")
	}
}
