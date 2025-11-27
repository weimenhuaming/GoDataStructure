package queue

// Item 定义元素类型
type Item struct {
	Value    string // 元素值
	Priority int    // 优先级（数值越小优先级越高）
	Index    int    // 堆中的索引（用于维护）
}

// PriorityQueue 定义优先队列（本质是一个堆）
type PriorityQueue []*Item

// =====================================================================================================================
// 实现 heap.Interface 的方法 => 实现一个大顶堆

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push 和 Pop 要修改队列长度，需用指针接收者
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // 标记已移除
	*pq = old[0 : n-1]
	return item
}
