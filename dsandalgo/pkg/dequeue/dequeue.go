package dequeue

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
