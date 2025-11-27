package queue

import "container/list"

// MonotonicQueue 这一个包下实现的是一个单调队列
// 单调递减的单调队列
type MonotonicQueue struct {
	q list.List // 直接使用包下的双向链表
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		q: list.List{},
	}
}

// Enqueue 入队:入队时，删除前面比入队元素小的元素
func (mq *MonotonicQueue) Enqueue(v int) {
	for mq.q.Len() > 0 && mq.q.Back().Value.(int) < v {
		mq.q.Remove(mq.q.Back())
	}
	mq.q.PushBack(v)
}

// Dequeue 出队:这里做判断的原因是只能出队首元素
// 元素已经被压缩了可能之前的最大值在后续已经被push操作压缩了
func (mq *MonotonicQueue) Dequeue(n int) {
	if n == mq.q.Front().Value.(int) {
		mq.q.Remove(mq.q.Front())
	}
}

// Max 返回最大值就是队首元素
func (mq *MonotonicQueue) Max() int {
	return mq.q.Front().Value.(int)
}
