package main

import (
	dll "github.com/parthrs/LetsGo/dsandalgo/pkg/doublylinkedlist"
)

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

type LRUCache struct {
	keyMap   map[int32]int32
	keyList  *dll.DoublyLinkedList
	capacity int32
}

func NewLRUCache(size int32) *LRUCache {
	return &LRUCache{
		keyMap:   map[int32]int32{},
		keyList:  dll.NewDoublyLinkedList(),
		capacity: size,
	}
}

func (c *LRUCache) put(k int32, v int32) {
	if len(c.keyMap) == int(c.capacity) {
		delete(c.keyMap, c.keyList.Head.Value)
		c.keyList.RemoveNodeHead()
	}
	c.keyMap[k] = v
	c.keyList.AddNodeTail(k)
}

func (c *LRUCache) get(k int32) int32 {
	if val, ok := c.keyMap[k]; ok {
		c.keyList.RemoveNode(k)
		c.keyList.AddNodeTail(k)
		return val
	}
	return -1
}

/*
func main() {
	c := NewLRUCache(2)
	c.put(1, 100)
	c.put(2, 200)
	c.put(3, 300)
	fmt.Println(c.keyMap)
	fmt.Println(c.keyList.Head.Value)
	fmt.Println(c.keyList.Head.Next.Value)
	fmt.Println(c.keyList.Head.Next.Next.Value)
	fmt.Println(c.keyList.Tail.Value)
	fmt.Println(c.keyList.Tail.Previous.Value)
	fmt.Println(c.keyList.Tail.Previous.Previous.Value)
	fmt.Printf("c.get(1): %d\n", c.get(1))
	fmt.Println(c.keyList.Head.Value)
	fmt.Println(c.keyList.Head.Next.Value)
	fmt.Println(c.keyList.Head.Next.Next.Value)
	fmt.Println(c.keyList.Tail.Value)
	fmt.Println(c.keyList.Tail.Previous.Value)
	fmt.Println(c.keyList.Tail.Previous.Previous.Value)
	fmt.Println(c.keyMap)
}
*/
