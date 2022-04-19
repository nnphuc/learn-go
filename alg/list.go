package main

import "fmt"

type Node struct {
	Next *Node
	Data interface{}
}

func CreateNode(data interface{}, next *Node) *Node {
	node := new(Node)
	node.Data = data
	node.Next = next
	return node
}

func PrintNode(node *Node) {
	cur := node
	for cur != nil {
		fmt.Printf("-> %v ", cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}

type IQueue interface {
	Top() interface{}
	Peek() interface{}
	Size() int
	Push(interface{})
}

type Queue struct {
	size  int
	first *Node
	last  *Node
}

func NewQueue() *Queue {
	q := new(Queue)
	q.size = 0
	q.first = nil
	q.last = nil
	return q
}

func (q *Queue) Top() interface{} {
	if q.size == 0 {
		return nil
	}
	return q.last.Data
}

func (q *Queue) Push(data interface{}) {
	if q.first == nil {
		q.first = CreateNode(data, nil)
		q.last = q.first
	} else {
		q.last.Next = CreateNode(data, nil)
		q.last = q.last.Next
	}
	q.size++
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Peek() interface{} {
	if q.size == 0 {
		return nil
	} else {
		q.size--
		data := q.first.Data
		q.first = q.first.Next
		return data
	}
}

func main() {
	list := CreateNode(1, CreateNode("banana", CreateNode(3.0, nil)))
	PrintNode(list)
	q := NewQueue()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	for q.Size() > 0 {
		fmt.Printf("%v\n", q.Peek())
	}
	var queue IQueue
	queue = q
	fmt.Printf("%v\n", queue)
}
