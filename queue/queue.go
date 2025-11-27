package queue

// 这里将实现一个简单的队列结构

type Queue struct {
	length int
}

func New() *Queue {
	return &Queue{
		length: 0,
	}
}

func (q *Queue) Enqueue(v interface{}) {

}

func (q *Queue) Dequeue() interface{} {

}
