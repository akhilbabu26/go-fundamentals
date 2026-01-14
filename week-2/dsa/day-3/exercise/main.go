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

func removeDuplicate(head *Node)*Node{
	if head == nil{
		return nil
	}

	visited := make(map[int]bool)
	current := head
	var prev *Node = nil

	for current != nil{
		if visited[current.value]{
			prev.next = current.next
		}else{
			visited[current.value] =true
			prev = current
		}
		current = current.next
	}
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
	
	fmt.Println("Befor duplicate removel:")
	printList(head)

	removeDuplicate(head)
	fmt.Println("After duplicate removel:")
	printList(head)

}