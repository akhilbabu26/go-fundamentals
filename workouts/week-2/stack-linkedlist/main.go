package main

import "fmt"

type Node struct{
	value int
	next *Node
}

type Stack struct{
	top *Node
}

func (s *Stack)Push(val int){
	newNode := &Node{
		value: val,
		next: s.top,
	}

	s.top = newNode
}

func (s *Stack)Pop(){
	s.top = s.top.next
}

func (s *Stack)PrintStack(){
	temp := s.top

	for temp != nil{
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main(){
	s := Stack{}
	s.Push(15)
	s.Push(33)
	s.Push(54)
	s.Push(12)
	s.PrintStack()
	s.Pop()
	s.PrintStack()
}

