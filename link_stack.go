package stack

type (
	LinkedStack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *LinkedStack {
	return &LinkedStack{}
}

// Push a value onto the top of the stack
func (s *LinkedStack) Push(v interface{}) {
	s.top = &node{v, s.top}
	s.length++
}

// Pop the top item of the stack and return it
func (s *LinkedStack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	v := s.top.value
	s.top = s.top.prev
	s.length--
	return v
}

// View the top item on the stack
func (s *LinkedStack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Return the number of items in the stack
func (s *LinkedStack) Len() int {
	return s.length
}
