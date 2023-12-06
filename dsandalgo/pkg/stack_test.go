package pkg

import "testing"

func TestStack(t *testing.T) {
	s := NewStack[int]()
	s.Add(5)
	s.Add(10)
	if len(s.S) != 2 {
		t.Errorf("Len expected to be 2, got %v", len(s.S))
	}
	if pop := s.Pop(); pop != 10 {
		t.Errorf("Expected elem 10, got %v", pop)
	}
}
