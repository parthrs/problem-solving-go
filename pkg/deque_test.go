package pkg

import "testing"

func TestDeque(t *testing.T) {
	dq := NewDeque[int](4)
	if dq.Len != 0 {
		t.Error("Length of empty Deque not zero")
	}
	if dq.Cap != 4 {
		t.Error("Capacity of Deque not as declared")
	}
	dq.PushBack(1)
	dq.PushFront(2)
	if dq.Len != 2 {
		t.Errorf("Expected length of Deque to be 2, got %d", dq.Len)
	}
	dq.PushFront(3)
	dq.PushFront(4)
	success := dq.PushFront(5)
	if success {
		t.Errorf("Deque exceeded capacity still push front returned true")
	}
	success = dq.PushBack(5)
	if success {
		t.Errorf("Deque exceeded capacity still push back returned true")
	}
	val, _ := dq.PopBack()
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
	val, _ = dq.PopFront()
	if val != 4 {
		t.Errorf("Expected 2, got %d", val)
	}
	dq.PopBack()
	dq.PopBack()
	if dq.Len != 0 {
		t.Errorf("Expected length of Deque to be 0, got %d", dq.Len)
	}
	_, success = dq.PopFront()
	if success {
		t.Errorf("Deque empty pop front returned true")
	}
	_, success = dq.PopBack()
	if success {
		t.Errorf("Deque empty pop back returned true")
	}
}
