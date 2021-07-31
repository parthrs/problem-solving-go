package dequeue

import "testing"

func TestDequeue(t *testing.T) {
	dq := NewDequeue(5)
	if dq.length != int32(0) {
		t.Error("Length of empty dequeue not zero")
	}
	if dq.capacity != int32(5) {
		t.Error("Capacity of dequeue not as declared")
	}
	dq.appendLeft(1)
	if dq.length != int32(1) {
		t.Error("Length of dequeue not as expected")
	}
	dq.appendRight(2)
	if dq.length != int32(2) {
		t.Error("Length of dequeue not as expected")
	}
	if dq.popLeft() != int32(1) {
		t.Error("Popleft didn't return expected item")
	}
	dq.appendLeft(0)
	if dq.popRight() != int32(2) {
		t.Error("Popright didn't return expected item")
	}
}
