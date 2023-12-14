package pkg

/*
Here are the ways a deque could be implemented:
 - Using a dynamic array:
   - Insertions to the front are ammortized to O(1)
	 - Insertions/deletions in the middle are O(n)
	 - Random access by index is O(1)
 - Using a LinkedList:
   - All insertions are O(1)
	 - Insertions/deletions in the middle are O(1)
	 - Random access by index is O(n)
*/

type Deque struct {
	length   int32
	li       []int32
	capacity int32
}

func NewDeque(size int32) Deque {
	return Deque{
		length:   0,
		li:       []int32{},
		capacity: size,
	}
}

func (d *Deque) appendLeft(k int32) {
	if d.length < d.capacity {
		d.li = append([]int32{k}, d.li...)
		d.length += 1
	}
}

func (d *Deque) appendRight(k int32) {
	if d.length < d.capacity {
		d.li = append(d.li, []int32{k}...)
		d.length += 1
	}
}

func (d *Deque) popLeft() int32 {
	var retVal int32 = -1
	if d.length > 0 {
		retVal = d.li[0]
		d.li = d.li[1:]
		d.length -= 1
		return retVal
	}
	return retVal
}

func (d *Deque) popRight() int32 {
	var retVal int32 = -1
	if d.length > 0 {
		retVal = d.li[d.length-1] // note the use of length - 1 since d.length gives the len, but really 'points' to the next location
		d.li = d.li[:d.length-1]
		d.length -= 1
		return retVal
	}
	return retVal
}
