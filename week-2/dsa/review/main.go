package main

import "fmt"

type Node struct{
	value int
	prev *Node
	next *Node
}

type LinkedList struct{
	head *Node
	tail *Node
}

func (ll *LinkedList)Insert(val int){
	newNode := &Node{value: val}

	if ll.head == nil{
		ll.head = newNode
		ll.tail = newNode
		return
	}

	newNode.next = ll.head
	ll.head.prev = newNode
	ll.head = newNode
}

func (ll *LinkedList)PrintList(){
	temp := ll.head

	for temp != nil{
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main(){
	l := LinkedList{}

	l.Insert(23)
	l.Insert(56)
	l.Insert(26)
	l.PrintList()
}

// package main

// import "fmt"

// type Stack struct{
// 	arr []int
// 	top int
// }

// func newStack()*Stack{
// 	return &Stack{
// 		arr: []int{},
// 		top: -1,
// 	}
// }

// func (s *Stack)Push(val int){
// 	s.arr = append(s.arr, val)
// 	s.top++
// }

// func (s *Stack)Pop(){
// 	s.arr = append(s.arr[:s.top], s.arr[s.top+1:]...)
// 	s.top--
// }

// func (s *Stack)printS(){
// 	fmt.Println(s.arr)
// }

// type Queue struct{
// 	arr []int
// 	front int
// 	rear int
// }

// func newQueue()*Queue{
// 	return &Queue{
// 		arr: []int{},
// 		front: 0,
// 		rear: -1,
// 	}
// }

// func (q *Queue)Enqueu(val int)[]int{
// 	q.arr = append(q.arr, val)
// 	q.rear++
// 	return q.arr
// }

// func (q *Queue)Dequeue(){
// 	q.arr = append(q.arr[:q.front], q.arr[:q.rear]...)
// }

// func (q *Queue)printQ(){
// 	fmt.Println(q.arr)
// }


// func main(){
// 	s := newStack()
// 	s.Push(45)
// 	s.Push(33)
// 	s.Push(44)
// 	s.printS()
// 	s.Pop()
// 	s.printS()

// 	q := newQueue()
// 	q.Enqueu(23)
// 	q.Enqueu(23)
// 	q.Enqueu(23)
// 	q.printQ()
// 	q.Dequeue()
// 	q.printQ()
// }