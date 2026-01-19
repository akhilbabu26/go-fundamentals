package main

import "fmt"

type Node struct{
	prev *Node
	value int
	next *Node
}

type LinkedList struct{
	head *Node
	tail *Node
}

func (ll *LinkedList)InsertHead(val int){
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

func (ll *LinkedList)InsertTail(val int){
	newNode := &Node{value: val}

	if ll.tail == nil{
		ll.tail = newNode
		ll.head = newNode
		return
	}

	newNode.prev = ll.tail
	ll.tail.next = newNode
	ll.tail = newNode
}

func (ll *LinkedList)PrintList(){
	current := ll.head

	for current != nil{
		fmt.Println(current.value)
		current = current.next
	}
}

func main(){
	ll := LinkedList{}

	ll.InsertHead(23)
	ll.InsertHead(98)
	ll.InsertHead(65)
	ll.InsertHead(87)
	ll.PrintList()

	ll.InsertTail(90)
	ll.InsertTail(33)
	ll.InsertTail(24)
	ll.InsertTail(65)
	ll.InsertTail(88)
	ll.PrintList()
}

