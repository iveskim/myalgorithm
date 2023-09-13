package sqqueue

type Queue struct {
	data                         []interface{}
	head, tail, length, capacity int
}

func New(capacity int) *Queue {
	return &Queue{
		data: make([]interface{}, capacity),
	}
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Put(v interface{}) {
	if q.length == q.capacity {
		panic("")
	}

	q.data[q.tail] = v

	q.tail = (q.tail + 1) % q.capacity
	q.length++
}

func (q *Queue) Get() interface{} {
	if q.length == 0 {
		return nil
	}
	v := q.data[q.head]
	q.data[q.head] = nil
	q.head = (q.head + 1) % q.capacity
	q.length--

	return v
}

func (q *Queue) Peek() interface{} {
	if q.length == 0 {
		return nil
	}
	return q.data[q.head]
}
