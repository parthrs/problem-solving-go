package main

import "fmt"

/*
Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

Implement the LRUCache class:

	LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
	int get(int key)
			Return the value of the key if the key exists, otherwise return -1.
	void put(int key, int value)
			Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache.
			If the number of keys exceeds the capacity from this operation, evict the least recently used key.

The functions get and put must each run in O(1) average time complexity.
*/

type Dequeue struct {
	length   int32
	li       []int32
	capacity int32
}

func NewDequeue(size int32) Dequeue {
	return Dequeue{
		length:   0,
		li:       []int32{},
		capacity: size,
	}
}

func (d *Dequeue) appendLeft(k int32) {
	if d.length < d.capacity {
		d.li = append([]int32{k}, d.li...)
		d.length += 1
	}
}

func (d *Dequeue) appendRight(k int32) {
	if d.length < d.capacity {
		d.li = append(d.li, []int32{k}...)
		d.length += 1
	}
}

func (d *Dequeue) popLeft() int32 {
	var retVal int32 = -1
	if d.length > 0 {
		retVal = d.li[0]
		d.li = d.li[1:]
		d.length -= 1
		return retVal
	}
	return retVal
}

func (d *Dequeue) popRight() int32 {
	var retVal int32 = -1
	if d.length > 0 {
		retVal = d.li[d.length-1] // note the use of length - 1 since d.length gives the len, but really 'points' to the next location
		d.li = d.li[:d.length-1]
		d.length -= 1
		return retVal
	}
	return retVal
}

type LRUCache struct {
	keyMap     map[int32]int32
	keyDequeue Dequeue
}

func NewLRUCache(size int32) *LRUCache {
	return &LRUCache{
		keyMap:     map[int32]int32{},
		keyDequeue: NewDequeue(size),
	}
}

func (c *LRUCache) put(k int32, v int32) {
	if c.keyDequeue.length == c.keyDequeue.capacity {
		key := c.keyDequeue.popLeft()
		delete(c.keyMap, key)
	}
	c.keyDequeue.appendRight(k)
	c.keyMap[k] = v
}

func (c *LRUCache) get(k int32) int32 {
	if val, ok := c.keyMap[k]; ok {
		c.markRecent(k)
		return val
	}
	return -1
}

func (c *LRUCache) markRecent(k int32) {
	dqStack := NewDequeue(c.keyDequeue.length - 1)
	elem := c.keyDequeue.popRight()
	for c.keyDequeue.length >= 1 && elem != k {
		dqStack.appendRight(elem)
		elem = c.keyDequeue.popRight()
	}
	for dqStack.length >= 1 {
		c.keyDequeue.appendRight(dqStack.popRight())
	}
	c.keyDequeue.appendRight(elem)
}

func main() {
	c := NewLRUCache(3)
	c.put(1, 100)
	c.put(2, 200)
	c.put(3, 300)
	fmt.Println(c)
	fmt.Println(c.get(1))
	fmt.Println(c)
}
