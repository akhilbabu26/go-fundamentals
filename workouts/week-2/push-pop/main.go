package main

import "fmt"

type Stack struct{
	data []int
	top int
}

func newStack()*Stack{
	return &Stack{
		data: []int{},
		top: -1,
	}
}

func (s *Stack)Push(val int){
	s.data = append(s.data , val)
	s.top++
}

func (s *Stack)Pop(){
	s.data = s.data[:s.top]
	s.top--
}

func (s *Stack)PrintStack(){
	for _,v:= range s.data{
		fmt.Println(v)
	}
	fmt.Println(s.data)
}

func main(){
	s := newStack()
	s.Push(45)
	s.Push(33)
	s.Push(23)
	s.Push(56)
	s.PrintStack()
	s.Pop()
	s.PrintStack()

}