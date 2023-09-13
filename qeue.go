package queue

type (
	Queue struct {
		head, tail *node
		length     int
	}
	node struct {
		value interface{}
		next  *node
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{}
}

// Take the next item off the front of the queue
func (q *Queue) Get() interface{} {
	if q.length == 0 {
		return nil
	}
	n := q.head
	if q.length == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.length--
	return n.value
}

// Put an item on the end of a queue
func (q *Queue) Put(v interface{}) {
	n := &node{v, nil}
	if q.length == 0 {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.length++
}

// Return the number of items in the queue
func (q *Queue) Len() int {
	return q.length
}

// Return the first item in the queue without removing it
func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.head.value
}
