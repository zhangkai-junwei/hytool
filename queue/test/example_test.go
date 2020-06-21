package test

import (
	"fmt"
	"hyTool/queue"
	"testing"
)

func TestQueue(t *testing.T) {
	q := queue.NewQueue("circle", 5)

	q.Enqueue(10)
	fmt.Println(q.Length())
	fmt.Println(q.Dequeue())
}
