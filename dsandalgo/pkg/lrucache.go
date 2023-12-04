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

// This implementation maintains two datastructures to get O(1)
// complexity; A map of (key, pointers -> DLL node) and a doubly
// linked list as a priority queue.
// The important detail to get O(1) is find a way to avoid
// the O(n) runtime while moving the least recently used node
// to the "first" or "most recently used". This is achieved by
// making the value for the doubly linked node is an array that
// stores {key, value}. When its time to "freshen" a node,
// we already know its position in the queue with the pointer.
// Finally because of this detail, we must maintain our own
// Doubly linked list and not rely on the standard implementation.
type LRUCache struct {
	KeyToNodePtrMap map[int]*DoublyLinkedNode[[2]int]
	Head            *DoublyLinkedNode[[2]int] // We manage our own DLL with head/tail ptrs
	Tail            *DoublyLinkedNode[[2]int]
	Capacity        int
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		KeyToNodePtrMap: map[int]*DoublyLinkedNode[[2]int]{},
		Capacity:        size,
	}
}

// Add elements to the tail
// We have to also manage the add logic
// for the doubly linked list
func (c *LRUCache) put(k int, v int) {
	// If cache is full,
	//   Update map
	//   Remove node at head
	if len(c.KeyToNodePtrMap) == c.Capacity {
		delete(c.KeyToNodePtrMap, c.Head.Value[0])
		c.Head = c.Head.Next
		c.Head.Previous = nil
	}

	// Add node at tail
	//   Create node
	//   If cache empty, update head
	//   else update tail.next
	//   Add node at tail
	n := NewDoublyLinkedNode[[2]int]([2]int{k, v})
	if c.Tail == nil {
		c.Head = n
	} else {
		c.Tail.Next = n
	}
	n.Previous = c.Tail
	c.Tail = n
	c.KeyToNodePtrMap[k] = n
}

func (c *LRUCache) get(k int) int {
	// If key does not exist return -1
	if _, found := c.KeyToNodePtrMap[k]; !found {
		return -1
	}
	// If key exists and not on top, move it to top
	//   Remove node from current position
	//   Add it to tail
	node := c.KeyToNodePtrMap[k]
	// Node is already at the tail
	if node.Next == nil {
		return node.Value[1]
	}
	// Node is at head
	if node.Previous == nil {
		c.Head = node.Next
		node.Next.Previous = nil
		node.Next = nil
		c.Tail.Next = node
		node.Previous = node
		c.Tail = node
	} else { // In the middle
		node.Previous.Next = node.Next
		node.Next.Previous = node.Previous
		node.Next = nil
		node.Previous = c.Tail
		c.Tail = node
	}
	return c.Tail.Value[1]
}
