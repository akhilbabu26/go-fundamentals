package main

import "fmt"

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (ll *LinkedList)InsertHead(val int){
	newNode := &Node{data: val}
	newNode.next = ll.head
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

	ll.InsertHead(55)
	ll.InsertHead(75)
	ll.InsertHead(58)
	ll.PrintList()
}