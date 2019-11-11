package main

// MyCircularDeque(k)：构造函数,双端队列的大小为k。
// insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
// insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
// deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
// deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
// getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
// getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
// isEmpty()：检查双端队列是否为空。
// isFull()：检查双端队列是否满了。

type CircularDeque struct {
	items []int
	n     int
	head  int
	tail  int
}

func NewCircularDeque(n int) CircularDeque {
	return CircularDeque{
		make([]int, n+1, n+1), n + 1, 0, 0,
	}
}

func (queue *CircularDeque) IsEmpty() bool {
	if queue.head == queue.tail {
		return true
	}
	return false
}

func (queue *CircularDeque) IsFull() bool {
	if (queue.tail+1)%queue.n == queue.head {
		return true
	}
	return false
}

// 如果head = 0, item放在末尾
// 如果head > 0, item放在前一位置
func (queue *CircularDeque) InsertFront(item int) bool {
	if queue.IsFull() {
		return false
	}
	queue.head = (queue.n + queue.head - 1) % queue.n
	queue.items[queue.head] = item
	return true
}

func (queue *CircularDeque) DeleteFront() bool {
	if queue.IsEmpty() {
		return false
	}
	queue.head = (queue.head + 1) % queue.n
	return true
}

func (queue *CircularDeque) InsertLast(item int) bool {
	if queue.IsFull() {
		return false
	}
	queue.items[queue.tail] = item
	queue.tail = (queue.tail + 1) % queue.n
	return true
}

func (queue *CircularDeque) DeleteLast() bool {
	if queue.IsEmpty() {
		return false
	}
	queue.tail = (queue.n + queue.tail - 1) % queue.n
	return true
}

func (queue *CircularDeque) GetFront() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.items[queue.head]
}

func (queue *CircularDeque) GetRear() int {
	if queue.IsEmpty() {
		return -1
	}
	return queue.items[(queue.n+queue.tail-1)%queue.n]
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
