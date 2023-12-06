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

// Using this over DoublyLinkedNode from pkg to pacify
// VSCode from complaining "undeclared name"
type Node[T comparable] struct {
	Previous *Node[T]
	Next     *Node[T]
	Value    T
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{
		Value: val,
	}
}

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
	KeyToNodePtrMap map[int]*Node[[2]int]
	Head            *Node[[2]int] // We manage our own DLL with head/tail ptrs
	Tail            *Node[[2]int]
	Capacity        int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		KeyToNodePtrMap: map[int]*Node[[2]int]{},
		Capacity:        capacity,
	}
}

// moveNodeToTail is only called when there are more
// than two elements in the cache
func (c *LRUCache) moveNodeToTail(node *Node[[2]int]) {
	// Node already at tail or only one elem in cache
	if node.Next == nil || len(c.KeyToNodePtrMap) == 1 {
		return
	}
	// Node at Head
	if node.Previous == nil {
		c.Head = node.Next
	} else { // Node in the middle
		node.Previous.Next = node.Next
	}
	node.Next.Previous = node.Previous
	c.Tail.Next = node
	node.Previous = c.Tail
	node.Next = nil
	c.Tail = node
}

// removeNodeAtHead is only called if cache
// is non-empty
func (c *LRUCache) removeNodeAtHead() {
	if c.Head.Next != nil {
		c.Head.Next.Previous = c.Head.Previous
	}
	c.Head = c.Head.Next
}

func (c *LRUCache) addNodeToTail(node *Node[[2]int]) {
	if len(c.KeyToNodePtrMap) == 0 {
		c.Head = node
	} else {
		c.Tail.Next = node
		node.Previous = c.Tail
	}
	c.Tail = node
}

// Add elements to the tail
// We have to also manage the add logic
// for the doubly linked list
func (c *LRUCache) put(k int, v int) {
	if node, found := c.KeyToNodePtrMap[k]; found {
		node.Value = [2]int{k, v}
		c.moveNodeToTail(node)
	} else {
		if len(c.KeyToNodePtrMap) == c.Capacity {
			headKey := c.Head.Value[0]
			c.removeNodeAtHead()
			delete(c.KeyToNodePtrMap, headKey)
		}
		node := NewNode[[2]int]([2]int{k, v})
		c.addNodeToTail(node)
		c.KeyToNodePtrMap[k] = node
	}
}

func (c *LRUCache) get(k int) int {
	if node, found := c.KeyToNodePtrMap[k]; found {
		c.moveNodeToTail(node)
		return node.Value[1]
	}
	return -1
}
