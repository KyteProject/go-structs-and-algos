package priorityqueue

// PriorityQueue is the heap where items are stored
type PriorityQueue []*Item

// Item is the data to be stored
type Item struct {
	value    interface{}
	priority int
	index    int
}

// Len returns the length of the queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less compares for greatest priority
func (pq PriorityQueue) Less(x, y int) bool {
	return (pq)[x].priority > (pq)[y].priority
}

// Push adds an item to the queue
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

// Pop removes the item with highest priority
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item.value
}

func (pq *PriorityQueue) Swap(x, y int) {
	(*pq)[x], (*pq)[y] = (*pq)[y], (*pq)[x]
	(*pq)[x].index = x
	(*pq)[y].index = y
}
