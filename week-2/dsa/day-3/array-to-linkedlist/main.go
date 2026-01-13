package main

import "fmt"

type Node struct{
	value int
	next *Node
}

func arrayToLinkedList(arr []int)*Node{
	if len(arr)==0{
		return nil
	}

	headNode := &Node{value: arr[0]}
	currentHead := headNode

	for i := 1; i < len(arr); i++{
		newNode := &Node{value: arr[i]}
		currentHead.next = newNode
		currentHead = newNode
	}
	return headNode
}

func printList(head *Node){
	for head != nil{
		fmt.Println(head.value)
		head = head.next
	}
}

func main(){
	arr := []int{8,17,5,4,8,99,100,54,700}
	head := arrayToLinkedList(arr)

	printList(head)

}