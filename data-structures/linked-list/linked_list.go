package list

// import "fmt"

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
