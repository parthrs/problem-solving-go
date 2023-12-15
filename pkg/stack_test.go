package pkg

import "testing"

func TestNewStack(t *testing.T) {
	// Stack with capacity of 5 elems
	s := NewStack[int](5)
	// Check len and capacity
	if length := s.Len; length != 0 {
		t.Errorf("Len expected to be 0, got %v", length)
	}
	if capacity := s.Cap; capacity != 5 {
		t.Errorf("Capacity expected to be 5, got %v", capacity)
	}

	// Push 5 elems 4 through 0 in order
	// Check for return status (of push)
	for i := 4; i >= 0; i-- {
		status := s.Push(i)
		if !status {
			t.Errorf("Expected true while push(%d), got %v", i, status)
		}
	}

	// Push beyond capacity, expect unsuccessful
	// operation
	status := s.Push(5)
	if status {
		t.Errorf("Expected false while push(5), got %v", status)
	}

	// Expect length to be intact and full
	if length := s.Len; length != 5 {
		t.Errorf("Len expected to be 5, got %v", length)
	}

	// Pop two elems and check return values
	// Check length after two pop operations
	if val := s.Pop(); val != 0 {
		t.Errorf("Expected val to be 0, got %d", val)
	}
	if val := s.Pop(); val != 1 {
		t.Errorf("Expected val to be 1, got %d", val)
	}
	if length := s.Len; length != 3 {
		t.Errorf("Len expected to be 3, got %v", length)
	}
	s.Pop() // 2
	s.Pop() // 3
	s.Pop() // 4
	if length := s.Len; length != 0 {
		t.Errorf("Len expected to be 0, got %v", length)
	}
}
