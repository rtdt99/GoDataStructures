package datastructures

import "fmt"

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
}

type QueueNode struct {
	Val  int
	Next *QueueNode
}

func (q *Queue) Enqueue(v int) {
	if q.Head == nil && q.Tail == nil {
		n := &QueueNode{v, nil}
		q.Head = n
		q.Tail = n
		return
	}

	n := &QueueNode{v, nil}
	q.Tail.Next = n
	q.Tail = q.Tail.Next
}

func (q *Queue) Dequeue() int {
	if q.Head == nil {
		return -1
	}

	result := q.Head.Val

	temp := q.Head

	q.Head = q.Head.Next

	temp.Next = nil

	return result
}

func (q *Queue) Print() {

	temp := q.Head

	for temp != nil {
		fmt.Println(temp.Val)
		temp = temp.Next
	}
}
