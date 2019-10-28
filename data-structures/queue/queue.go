package queue

import "container/list"

// Queue is the queue structure
type Queue struct {
	queue *list.List
}

// New creates a new queue
func New() *Queue {
	q := new(Queue)
	q.queue = list.New()
	return q
}

// Len returns the length of the linked list
func (q *Queue) Len() int {
	return q.queue.Len()
}

// IsEmpty returns true is the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.queue.Len() == 0
}

// Clear clears the list
func (q *Queue) Clear() {
	q.queue.Init()
}

// Enqueue adds an item to the back of the queue
func (q *Queue) Enqueue(i interface{}) (e *list.Element) {
	return q.queue.PushBack(i)
}

// Dequeue removes the first item in the queue
func (q *Queue) Dequeue() (i interface{}) {
	if el := q.queue.Front(); el != nil {
		q.queue.Remove(el)
		return el.Value
	}
	return nil
}

// Peek returns the value in the beginning of queue
func (q *Queue) Peek() (i interface{}) {
	if el := q.queue.Front(); el != nil {
		return el.Value
	}
	return nil
}
