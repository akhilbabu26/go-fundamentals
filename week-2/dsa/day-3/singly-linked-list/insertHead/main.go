// package main

// import "fmt"

// type Node struct{
// 	value int
// 	next *Node
// }

// type LinkedList struct{
// 	head *Node
// }

// func (ll *LinkedList) InsertHead(val int){
// 	newNode := &Node{value: val}
// 	newNode.next = ll.head
// 	ll.head = newNode
// }

// func (ll *LinkedList)Display(){
// 	temp := ll.head
// 	for temp != nil{
// 		fmt.Println(temp.value)
// 		temp = temp.next
// 	}
// }

// func main(){
// 	list := LinkedList{}

// 	list.InsertHead(155)
// 	list.InsertHead(166)
// 	list.InsertHead(177)

// 	list.Display()
	
// }

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

// Insert at Head
func insertHead(head *Node, value int) *Node {
	newNode := &Node{value: value}
	newNode.next = head
	return newNode
}

// Print List
func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	var head *Node = nil

	head = insertHead(head, 10)
	head = insertHead(head, 20)
	head = insertHead(head, 30)

	printList(head)
}
