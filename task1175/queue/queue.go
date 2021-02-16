package queue

import "fmt"

type Node struct {
	Value int
	Next *Node
}

type Queue struct {
	Size int
	Head *Node
	Tail *Node
	
}

func NewQueue() Queue {
	return Queue{
		Size: 0,
		Head: nil,
		Tail: nil,
	}
}

func (q Queue) Print() {
	if q.Head == q.Tail {
		fmt.Println(fmt.Errorf("queue is empty"))
	} else {
		for node := q.Head; node != nil; node = node.Next {
			fmt.Printf("%v ", node.Value)
		}
		fmt.Println()
	}
}

func (q *Queue) Queue(el int) {
	if q.Head == nil {
		q.Tail = &Node{el, nil}
		q.Head = q.Tail
	} else {
		q.Tail.Next = &Node{el, nil}
		q.Tail = q.Tail.Next
	}
	q.Size++
}

func (q *Queue) Dequeue() (int, error) {
	size := q.Size
	if size == 0 {
		return 0, fmt.Errorf("queue is emtpy")
	}
	el := q.Head.Value
	q.Head = q.Head.Next
	q.Size--
	return el, nil
}