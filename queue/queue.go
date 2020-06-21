package queue

type Queue interface {
	Length() int
	Capacity() int
	Front() *Node
	Rear() *Node
	Enqueue(value interface{}) bool
	Dequeue() interface{}
}

func NewQueue(name string, capacity int) Queue {
	var q Queue

	switch name {
	case "normal":
		q, _ = NewNormalQueue(capacity)
	case "unique":
		q, _ = NewUniqueQueue(capacity)
	case "circle":
		q, _ = NewCyclicQueue(capacity)
	}
	return q
}
