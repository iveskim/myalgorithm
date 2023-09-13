package SequentialStack

type SqStack struct {
	data     []interface{}
	length   int
	capacity int
}

func New(capacity int) *SqStack {
	return &SqStack{
		data:     make([]interface{}, capacity),
		length:   0,
		capacity: capacity,
	}
}

func (s *SqStack) Push(v interface{}) {
	if s.length >= s.capacity {
		panic("")
	}
	s.data[s.length] = v
	s.length++
}

func (s *SqStack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	v := s.data[s.length-1]
	s.length--
	return v
}

func (s *SqStack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.data[s.length-1]
}

func (s *SqStack) Len() int {
	return s.length
}
