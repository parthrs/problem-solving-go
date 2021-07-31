package main

import (
	"reflect"
	"testing"
)

func TestLRUCache(t *testing.T) {
	c := NewLRUCache(3)
	c.put(1, 100)
	c.put(2, 200)
	c.put(3, 300)
	if c.keyDequeue.length != 3 {
		t.Error("Length of cache not matching items added")
	}
	if c.get(1) != 100 {
		t.Error("Method get failed to get correct value")
	}
	if ok := reflect.DeepEqual(c.keyDequeue.li, []int32{2, 3, 1}); !ok {
		t.Error("Cache order not correct after get request")
	}
	c.put(4, 400)
	if c.keyDequeue.length != 3 {
		t.Error("Length of cache not matching items added")
	}
	if ok := reflect.DeepEqual(c.keyDequeue.li, []int32{3, 1, 4}); !ok {
		t.Error("Cache order not correct after get request")
	}
}
