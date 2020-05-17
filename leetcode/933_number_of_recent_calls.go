package goleetcode

import "container/list"

type RecentCounter struct {
	queue *list.List
}

func RecentCounterConstructor() RecentCounter {
	return RecentCounter{queue: list.New()}
}

// Learning: Elements from the list need to be cast into specific type before being used.
func (this *RecentCounter) Ping(t int) int {
	this.queue.PushBack(t)
	for this.queue.Front().Value.(int) < t-3000 {
		this.queue.Remove(this.queue.Front())
	}

	return this.queue.Len()
}
