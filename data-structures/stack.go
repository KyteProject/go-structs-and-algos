package stack

// Stack is a LIFO data-structure
type Stack struct {
	stack []interface{}
	depth int
}

// New creates a new stack
func New() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.depth = 0

	return s
}

// Push adds a new item to the top of the Stack
func (s *Stack) Push(item interface{}) {
	i := make([]interface{}, 1)
	i[0] = item
	s.stack = append(i, s.stack...)
	s.depth++
}

// Pop removes item from top of stack and returns it
func (s *Stack) Pop() (item interface{}) {
	if s.depth <= 0 {
		return nil
	}

	item, s.stack = s.stack[0], s.stack[1:]
	s.depth--

	return item
}

// IsEmpty returns true if stack is empty
func (s *Stack) IsEmpty() bool {
	return s.depth == 0
}

// Peek returns the top of stack without deleting
func (s *Stack) Peek() (item interface{}) {
	if s.depth <= 0 {
		return nil
	}

	return s.stack[0]
}
