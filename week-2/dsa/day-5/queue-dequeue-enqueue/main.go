package main

import "fmt"

type Queue struct{
	arr []int
	front int
	rear int
}

func NewQueue()*Queue{
	return &Queue{
		arr: []int{},
		front: 0,
		rear:-1,
	}
}

func (q *Queue) Enqueue(val int){
	q.arr = append(q.arr, val)
	q.rear++

}

func (q *Queue) Dequeue()int{
	if q.isEmpty(){
		fmt.Println("Queue is empty")
		return -1
	}

	val := q.arr[q.front]
	q.front++
	return val
}

func (q *Queue)isEmpty()bool{
	return q.front > q.rear
}

func main(){
	q := NewQueue()

	q.Enqueue(45)
	q.Enqueue(5)
	q.Enqueue(4)

	fmt.Println("Dequeue:", q.Dequeue()) 
    fmt.Println("Dequeue:", q.Dequeue()) 
    fmt.Println("Dequeue:", q.Dequeue()) 
    fmt.Println("Dequeue:", q.Dequeue()) 

	
}