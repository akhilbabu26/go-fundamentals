package main

import "fmt"

type Node struct{
	data int
	next *Node
}

type Queue struct{
	front *Node
	rear *Node
}

func newQueue()*Queue{
	return &Queue{
		front: nil,
		rear: nil,
	}
}

func (q *Queue)Enqueue(val int){
	newNode := &Node{data: val}

	if q.front == nil{
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue)Dequeue()int{
	if q.front == nil{
		fmt.Println("Queue is empty")
		return -1
	}

	val := q.front.data
	q.front = q.front.next

	return val
}

func (q *Queue)Display(){
	if q.front == nil{
		fmt.Println("Queue is empty")
		return
	}

	temp := q.front
	fmt.Println("Queue: ")
	for temp != nil{
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main(){
	q := newQueue()

	q.Enqueue(14)
	q.Enqueue(19)
	q.Enqueue(54)
	q.Enqueue(100)

	q.Display()
	q.Dequeue()
	q.Display()
}