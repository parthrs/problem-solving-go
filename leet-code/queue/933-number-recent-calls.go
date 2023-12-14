package queue

type RecentCounter struct {
	Timestamps  []int
	Windowstart int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Timestamps:  make([]int, 0, 10001),
		Windowstart: 0,
	}
}

// LC insists we use 'this' ¯\_(ツ)_/¯
func (this *RecentCounter) Ping(t int) int {
	this.Timestamps = append(this.Timestamps, t)
	rangeStart := t - 3000
	for this.Timestamps[this.Windowstart] < rangeStart {
		this.Windowstart++
	}
	return len(this.Timestamps[this.Windowstart:])
}
