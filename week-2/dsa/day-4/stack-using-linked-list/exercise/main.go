package main

import "fmt"

type Node struct{
	value int
	next *Node
}

type Stack struct{
	top *Node
}

func NewStack()*Stack{
	return &Stack{top: nil}
}

//push
func (s *Stack)Push(val int){
	newNode := &Node{
		value: val,
		next: s.top,
	}
	s.top = newNode
}

//pop
func (s *Stack)Pop(){
	if s.top == nil{
		fmt.Println("Stack is empty")
		return
	}

	val := s.top.value // pop value
	s.top = s.top.next
	fmt.Printf("Last value removed: %d, and new list:\n",val)
}

//print
func (s *Stack)PrintStack(){
	if s.top == nil{
		fmt.Println("Stack is Empty")
		return
	}

	current := s.top

	for current !=nil{
		fmt.Println(current.value)
		current = current.next
	}
}


func main(){
	stack := NewStack()
	
	stack.Push(55)
	stack.Push(100)
	stack.Push(77)

	stack.PrintStack()
	stack.Pop()
	stack.PrintStack()
}