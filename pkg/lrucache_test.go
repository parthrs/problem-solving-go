package pkg

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(3)
	c.put(1, 100) // (1)
	if got := c.Head.Value[1]; got != 100 {
		t.Errorf("Cache head value incorrect (expected 100, got %d)", got)
	}
	if got := c.Tail.Value[1]; got != 100 {
		t.Errorf("Cache tail value incorrect (expected 100, got %d)", got)
	}
	c.put(2, 200) // (1, 2)
	if got := c.Head.Value[1]; got != 100 {
		t.Errorf("Cache head value incorrect (expected 100, got %d)", got)
	}
	if got := c.Tail.Value[1]; got != 200 {
		t.Errorf("Cache tail value incorrect (expected 200, got %d)", got)
	}
	if got := c.Head.Next.Value[1]; got != 200 {
		t.Errorf("Cache head.Next value incorrect (expected 200, got %d)", got)
	}
	if got := c.Tail.Previous.Value[1]; got != 100 {
		t.Errorf("Cache tail.Previous value incorrect (expected 100, got %d)", got)
	}
	c.put(3, 300) // (1,2,3)
	if got := c.Tail.Value[1]; got != 300 {
		t.Errorf("Cache tail value incorrect (expected 300, got %d)", got)
	}
	if c.get(4) != -1 {
		t.Error("Incorrect return value when elem not present")
	}
	if c.get(1) != 100 { // (2,3,1)
		t.Error("Method get failed to get correct value")
	}
	c.put(4, 400) // 2 should get evicted (3,1,4)
	if result := c.get(2); result != -1 {
		t.Errorf("Failed to evict, expected: %d, Got: %d", -1, result)
	}
	if got := c.Head.Value[1]; got != 300 {
		t.Errorf("Cache head value incorrect (expected 300, got %d)", got)
	}
	if got := c.Tail.Value[1]; got != 400 {
		t.Errorf("Cache tail value incorrect (expected 400, got %d)", got)
	}
	if got := c.Head.Next.Value[1]; got != 100 {
		t.Errorf("Cache head.Next value incorrect (expected 100, got %d)", got)
	}
	if got := c.Tail.Previous.Value[1]; got != 100 {
		t.Errorf("Cache tail.Previous value incorrect (expected 100, got %d)", got)
	}
}

// Cache with size 1
func TestLRUCacheLcTest1(t *testing.T) {
	c := NewLRUCache(1)
	c.put(2, 1)
	if result := c.get(2); result != 1 {
		t.Errorf("c.get(2) - Expected %v, Got %v", 1, result)
	}
	c.put(3, 2)
	if result := c.get(2); result != -1 {
		t.Errorf("c.get(2) - Expected %v, Got %v", -1, result)
	}
	if result := c.get(3); result != 2 {
		t.Errorf("c.get(3) - Expected %v, Got %v", 2, result)
	}
}

// Overwriting an existing key
func TestLRUCacheLcTest2(t *testing.T) {
	c := NewLRUCache(2)
	if result := c.get(2); result != -1 {
		t.Errorf("c.get(2) - Expected %v, Got %v", -1, result)
	}
	c.put(2, 6)
	if result := c.get(1); result != -1 {
		t.Errorf("c.get(1) - Expected %v, Got %v", -1, result)
	}
	c.put(1, 5)
	c.put(1, 2)
	if result := c.get(1); result != 2 {
		t.Errorf("c.get(1) - Expected %v, Got %v", 2, result)
	}
	if result := c.get(2); result != 6 {
		t.Errorf("c.get(3) - Expected %v, Got %v", 6, result)
	}
}

// ["LRUCache","get","get","put","get","put","put","put","put","get","put"]
// [[1],[6],[8],[12,1],[2],[15,11],[5,2],[1,15],[4,2],[5],[15,15]]

func TestLRUCacheLcTest3(t *testing.T) {
	c := NewLRUCache(1)
	if result := c.get(6); result != -1 {
		t.Errorf("c.get(6) - Expected %v, Got %v", -1, result)
	}
	if result := c.get(8); result != -1 {
		t.Errorf("c.get(8) - Expected %v, Got %v", -1, result)
	}
	c.put(12, 1)
	if result := c.get(2); result != -1 {
		t.Errorf("c.get(2) - Expected %v, Got %v", -1, result)
	}
	c.put(15, 11)
	c.put(5, 2)
	c.put(1, 15)
	c.put(4, 2)
	if result := c.get(5); result != -1 {
		t.Errorf("c.get(5) - Expected %v, Got %v", -1, result)
	}
	c.put(15, 15)
}
