package main

import "fmt"

type Node struct{
	value int
	next *Node
}

type Queue struct{
	front *Node
	rear *Node
}

func NewQueue()*Queue{
	return &Queue{
		front: nil,
		rear: nil,
	}
}

func (q *Queue)Enqueue(val int){
	newNode := &Node{value: val}

	if q.rear == nil{
		q.front = newNode
		q.rear = newNode
		return
	}

	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue)Dequeue(){
	q.front = q.front.next
}

func (q *Queue)PrintQ(){
	temp := q.front

	for temp != nil{
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main(){
	q := NewQueue()
	q.Enqueue(55)
	q.Enqueue(22)
	q.Enqueue(11)
	q.PrintQ()
	q.Dequeue()  
	q.PrintQ()

}