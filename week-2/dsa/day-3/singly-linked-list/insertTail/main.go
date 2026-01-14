// package main

// import "fmt"

// type Node struct{
// 	value int
// 	next *Node
// }

// type LinkedList struct{
// 	head *Node
// }

// func (ll *LinkedList) InsertTail(val int){
// 	newNode := &Node{value: val}

// 	if ll.head == nil{
// 		ll.head = newNode
// 		return
// 	}

// 	temp := ll.head

// 	for temp.next != nil{
// 		temp = temp.next
// 	}
// 	temp.next = newNode

// }

// func (ll *LinkedList) Display(){
// 	temp := ll.head

// 	for temp != nil{
// 		fmt.Println(temp.value)
// 		temp = temp.next
// 	}
// }

// func main(){
// 	list := LinkedList{}

// 	list.InsertTail(155)
// 	list.InsertTail(166)
// 	list.InsertTail(177)

// 	list.Display()
	
// }

package main

import "fmt"

type Node struct{
	value int
	next *Node
}

func insertTail(head *Node, val int)*Node{
	newNode := &Node{value: val}

	if head == nil{
		return newNode
	}

	current := head
	for current.next != nil{
		current = current.next
	}

	current.next = newNode
	return head
}

func printList(head *Node){
	current := head

	for current != nil{
		fmt.Println(current.value)
		current = current.next
	}
}


func main(){
	var head *Node = nil

	val := []int{100,500,100,700,400,500,600,700}

	for _,v := range val{
		head = insertTail(head, v)
	} 

	printList(head)

}