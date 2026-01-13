package main

import "fmt"

type Node struct{
	value int
	next *Node
}

type LinkedList struct{
	head *Node
}

// insert first
func (ll *LinkedList) InsertHead(val int){
	newNode := &Node{value: val}
	newNode.next = ll.head
	ll.head = newNode
}

// insert last
func (ll *LinkedList)InsertLast(val int){
	newNode := &Node{value : val}
	
	if ll.head == nil{
		ll.head = newNode
		return
	}
	temp := ll.head

	for temp.next != nil{
		temp = temp.next
	}
	temp.next = newNode
}

//search
func (ll *LinkedList)Search(val int)bool{
	temp := ll.head

	for temp != nil {
		if temp.value == val{
			return true
		}
		temp = temp.next
	}
	return false
}


func (ll *LinkedList)Display(){
	temp := ll.head

	for temp != nil{
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main(){
	list := LinkedList{}

	list.InsertHead(12)
	list.InsertHead(22)
	list.InsertHead(32)
	list.Display()

	list.InsertLast(55)
	list.Display()

	result := list.Search(12)

	fmt.Println(result)
}