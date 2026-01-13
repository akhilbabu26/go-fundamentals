package main

import "fmt"

type Node struct{
	value int
	prev *Node
	next *Node
}

type Dll struct{
	head *Node
	tail *Node
}

func (dll *Dll)InsertFront(val int){
	newNode := &Node{value: val}

	if dll.head == nil{
		dll.head = newNode
		dll.tail = newNode
		return
	}

	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
}

func (dll *Dll)DeleteLast(){
	if dll.tail == nil{
		fmt.Println("List is empty")
		return
	}

	// If only one node
	if dll.head == dll.tail{
		dll.head = nil
		dll.head = nil
		return
	}

	dll.tail = dll.tail.prev
	dll.tail.next = nil
}

func (dll *Dll)PrintForward(){
	current := dll.head

	for current != nil{
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
	dll := Dll{}

	dll.InsertFront(10)
	dll.InsertFront(20)
	dll.InsertFront(30)

	fmt.Println("After InsertFront operations:")
	dll.PrintForward()

	dll.DeleteLast()

	fmt.Println("After DeleteLast operation:")
	dll.PrintForward()
}