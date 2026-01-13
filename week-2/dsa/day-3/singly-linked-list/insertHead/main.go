// package main

// import "fmt"

// // Node represents a single node in the linked list
// type Node struct {
// 	data int
// 	next *Node
// }

// // LinkedList represents the singly linked list
// type LinkedList struct {
// 	head *Node
// }

// // InsertAtHead inserts a new node at the beginning
// func (ll *LinkedList) InsertAtHead(value int) {
// 	newNode := &Node{data: value}
// 	newNode.next = ll.head
// 	ll.head = newNode
// }

// // InsertAtTail inserts a new node at the end
// func (ll *LinkedList) InsertAtTail(value int) {
// 	newNode := &Node{data: value}

// 	if ll.head == nil {
// 		ll.head = newNode
// 		return
// 	}

// 	temp := ll.head
// 	for temp.next != nil {
// 		temp = temp.next
// 	}

// 	temp.next = newNode
// }

// // Display prints the linked list
// func (ll *LinkedList) Display() {
// 	temp := ll.head
// 	for temp != nil {
// 		fmt.Print(temp.data," ")
// 		temp = temp.next
// 	}
// 	fmt.Println("NULL")
// }

// // Main function
// func main() {
// 	list := LinkedList{}

// 	list.InsertAtHead(10)
// 	list.InsertAtHead(20)
// 	list.InsertAtTail(30)
// 	list.InsertAtTail(40)

// 	list.Display()
// }

package main

import "fmt"

type Node struct{
	value int
	next *Node
}

type LinkedList struct{
	head *Node
}

func (ll *LinkedList) InsertHead(val int){
	newNode := &Node{value: val}
	newNode.next = ll.head
	ll.head = newNode
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

	list.InsertHead(155)
	list.InsertHead(166)
	list.InsertHead(177)

	list.Display()
	
}