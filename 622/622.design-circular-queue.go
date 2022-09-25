package leetcode

/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 */

// @lc code=start
type MyCircularQueue struct {
	queue []int
	size  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, 0),
		size:  k,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if len(this.queue) == this.size {
		return false
	}

	this.queue = append(this.queue, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.queue) == 0 {
		return false
	}
	this.queue = this.queue[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}

func (this *MyCircularQueue) Rear() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.queue) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.queue) == this.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end
