package misc

import "testing"

func TestSnapshotMap(t *testing.T) {
	s := NewSnapshotMap[string, int]()
	s.Put("a", 1)
	s.Put("b", 2)
	s.Put("c", 1)
	snapId := s.TakeSnapshot()
	if snapId != 1 {
		t.Errorf("Expected snapshot ID to be 1, got: %v", snapId)
	}

	s.Put("a", 10)
	s.Put("b", 3)
	snapId = s.TakeSnapshot()
	if snapId != 2 {
		t.Errorf("Expected snapshot ID to be 2, got: %v", snapId)
	}

	s.Put("b", 100)
	s.TakeSnapshot()
	s.Put("a", 100)

	result, err := s.Get("a", 3)
	if err != nil {
		t.Errorf("Expected non-error response for s.Get(a), got: %v", err)
	}
	if result != 10 {
		t.Errorf("For s.Get(a, 3): Expected 10, got %v", result)
	}

	result, err = s.Get("b", 3)
	if err != nil {
		t.Errorf("Expected non-error response for s.Get(b, 3), got: %v", err)
	}
	if result != 100 {
		t.Errorf("For s.Get(b, 3): Expected 100, got %v", result)
	}

	result, err = s.Get("a", 0)
	if err != nil {
		t.Errorf("Expected non-error response for s.Get(a, 0), got: %v", err)
	}
	if result != 100 {
		t.Errorf("For s.Get(a, 0): Expected 100, got %v", result)
	}

	result, err = s.Get("c", 3)
	if err != nil {
		t.Errorf("Expected non-error response for s.Get(c, 3), got: %v", err)
	}
	if result != 1 {
		t.Errorf("For s.Get(c, 3): Expected 1, got %v", result)
	}

	result, err = s.Get("d", 3)
	if err == nil {
		t.Errorf("Expected error response for s.Get(d, 3), got: %v", err)
	}
	if result != 0 {
		t.Errorf("Expected zero value for int 0, got %v", result)
	}

	s.Delete("a")

	_, err = s.Get("a", 0)
	if err == nil {
		t.Errorf("Expected error response for s.Get(a, 0), got: %v", err)
	}

	result, err = s.Get("a", 3)
	if err != nil {
		t.Errorf("Expected non-error response for s.Get(a, 3), got: %v", err)
	}
	if result != 10 {
		t.Errorf("For s.Get(a, 3): Expected 10, got %v", result)
	}
	s.TakeSnapshot() // 4
	s.Put("d", 599)
	s.Put("k", 111)
	s.Put("o", 55)
	s.TakeSnapshot() // 5
	s.Put("c", 699)
	s.TakeSnapshot() // 6
	result, err = s.Get("a", 6)
	if err == nil {
		t.Errorf("Expected error response for s.Get(a, 6), got: %v", err)
	}
	if result != 0 {
		t.Errorf("For s.Get(a, 6): Expected 0, got %v", result)
	}
	result, err = s.Get("a", 4)
	if err == nil {
		t.Errorf("Expected error response for s.Get(a, 4), got: %v", err)
	}
	if result != 0 {
		t.Errorf("For s.Get(a, 4): Expected 0, got %v", result)
	}
	result, err = s.Get("a", 3)
	if err != nil {
		t.Errorf("Expected non-error response for s.Get(a, 3), got: %v", err)
	}
	if result != 10 {
		t.Errorf("For s.Get(a, 3): Expected 0, got %v", result)
	}
}
