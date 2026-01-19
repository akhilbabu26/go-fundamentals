package main

import "fmt"

type Node struct{
	data int
	prev *Node
	next *Node
}

type LinkedList struct{
	head *Node
	tail *Node
}

func (ll *LinkedList)InsertValue(val int){
	newNode := &Node{data: val}

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
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main(){
	ll := LinkedList{}

	ll.InsertValue(77)
	ll.InsertValue(87)
	ll.InsertValue(74)
	ll.PrintList()
}