package main

import "fmt"

type Stack struct{
	arr []int
	top int
}

func NewStack()*Stack{
	return &Stack{
		arr: []int{},
		top: -1,
	}
}

//push
func (s *Stack)Push(val int){
	s.arr = append(s.arr, val)
	s.top++
}

//pop
func (s *Stack)Pop(){
	if s.top == -1{
		fmt.Println("Empty stack") // this is peek
	}

	s.arr = s.arr[:s.top]
	s.top--
	fmt.Println("After pop:")
}

func (s *Stack)PrintStack(){
	if s.top == -1{
		fmt.Println("Empty Stack")
		return
	}

	for i := 0; i <= s.top; i++{
		fmt.Println(s.arr[i])
	}
}

func main(){
	stack := NewStack()

	stack.Push(58)
	stack.Push(88)
	stack.Push(50)

	stack.PrintStack()

	stack.Pop()
	stack.PrintStack()
}