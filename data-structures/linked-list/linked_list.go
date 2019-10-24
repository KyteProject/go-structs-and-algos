package list

import (
	"errors"
)

// Node is the structure for single linked list node
type Node struct {
	data interface{}
	next *Node
}

// LinkedList is the structure for a linked list
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// NewList makes a new list
func NewList() *LinkedList {
	list := new(LinkedList)
	list.size = 0
	return list
}

// NewNode creates a new node
func NewNode(data interface{}) *Node {
	return &Node{data, nil}
}

// IsEmpty returns true is the list is empty
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// Len returns the length of the linked list
func (l *LinkedList) Len() int {
	return l.size
}

// Clear clears the list
func (l *LinkedList) Clear() {
	l.head, l.tail, l.size = nil, nil, 0
}

// PeekFirst returns the data held in the head
func (l *LinkedList) PeekFirst() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.head.data
}

// PeekNext returns the data held in the next node
func (l *LinkedList) PeekNext() interface{} {
	if l.IsEmpty() {
		return nil
	}

	if l.Len() <= 1 {
		return nil
	}

	return l.head.next.data
}

// PeekLast returns the data held in the tail
func (l *LinkedList) PeekLast() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.tail.data
}

// Add an element to the tail of the list
func (l *LinkedList) Add(item interface{}) {
	l.AddFirst(item)
}

// AddFirst adds to the beginning of the linked list
func (l *LinkedList) AddFirst(item interface{}) {
	node := NewNode(item)

	if l.IsEmpty() {
		l.head, l.tail = node, node
	} else {
		node.next = l.head
		l.head = node
	}

	l.size++
}

// AddLast adds to the tail of the linked list
func (l *LinkedList) AddLast(item interface{}) {
	node := NewNode(item)

	if l.IsEmpty() {
		l.head, l.tail = node, node
	} else {
		l.tail.next = node
		l.tail = node
	}

	l.size++
}

// RemoveFirst deletes the node held in head
func (l *LinkedList) RemoveFirst() (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("list is empty")
	}

	data := l.head.data
	l.head = l.head.next
	l.size--

	if l.IsEmpty() {
		l.tail = nil
	}

	return data, nil
}

// RemoveLast deletes the node held in tail
func (l *LinkedList) RemoveLast() (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("list is empty")
	}

	if l.head.next == nil {
		return l.RemoveFirst()
	}

	// Traverse the list
	trav := l.head
	for ; trav.next.next != nil; trav = trav.next {
	}

	ret := trav.next.data
	trav.next = nil
	l.tail = trav
	l.size--

	return ret, nil
}

// RemoveByValue deletes a specific node by its value
func (l *LinkedList) RemoveByValue(item interface{}) (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("list is empty")
	}

	if l.PeekFirst() == item {
		return l.RemoveFirst()
	} else if l.PeekLast() == item {
		return l.RemoveLast()
	}

	trav := l.head
	for ; trav.next != nil; trav = trav.next {
		if trav.next.data == item {
			ret := trav.next.data
			trav.next = trav.next.next
			l.size--

			return ret, nil
		}
	}

	return nil, errors.New("unable to find item")
}
