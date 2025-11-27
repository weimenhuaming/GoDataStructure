package queue

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 插入元素
	heap.Push(&pq, &Item{Value: "A", Priority: 3})
	heap.Push(&pq, &Item{Value: "B", Priority: 1}) // 优先级最高
	heap.Push(&pq, &Item{Value: "C", Priority: 2})

	// 按优先级顺序弹出
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.Value, item.Priority)
	}
}
