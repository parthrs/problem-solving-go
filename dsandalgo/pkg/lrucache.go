package pkg

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
	keyList  *DoublyLinkedList[int32]
	capacity int32
}

func NewLRUCache(size int32) *LRUCache {
	return &LRUCache{
		keyMap:   map[int32]int32{},
		keyList:  NewDoublyLinkedList[int32](),
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
