package main

import "fmt"

// insert value into slice
func insertArr(arr []int, index int, val int)[]int{
	if index < 0 || index > len(arr){
		fmt.Println("ERROR: invalid input")
		return arr
	}

	arr = append(arr, 0)
	copy(arr[index+1:], arr[index:])
	arr[index] = val
	return arr
}
// delete value from slice
func deleteValue(arr []int, index int)[]int{
	if index < 0 || index >= len(arr){
		fmt.Println("ERROR: invalid input")
		return arr
	}
	return append(arr[:index],arr[index+1:]...)
}
// search from slice
func searchValue(arr []int,val int)int{

	for i,v := range arr{
		if v == val{
			return i
		}
	}
	return -1
}
// global node
type Node struct{
	data int
	next *Node
}

// Linked List insert
type LinkedList struct{
	head *Node
}

//insert
func (ll *LinkedList)Insert(val int){
	newNode := &Node{data: val}
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
// delete from linked list
func (ll *LinkedList)DeleteVal(val int){
	if ll.head == nil{
		return
	}
	if ll.head.data == val{
		ll.head = ll.head.next
		return
	}
	prev := ll.head
	curr := ll.head.next
	for curr != nil && curr.data != val{
		prev = curr
		curr = curr.next
	}
	if curr != nil{
		prev.next = curr.next
	}
}
// convert array into linked list
func arrayToLl(arr []int)*LinkedList{
	ll := &LinkedList{}

	for _,v := range arr{
		ll.Insert(v)
	}
	return ll
}

// print linked list
func (ll *LinkedList)PrintList(){
	temp := ll.head

	fmt.Println("Linked list: ")
	for temp != nil{
		fmt.Println(temp.data, " ")
		temp = temp.next
	}
}

// stack
type Stack struct{
	items []int
	top int
}

func newStack()*Stack{
	return &Stack{
		items: []int{},
		top: -1,
	}
}

// stack pop
func (s *Stack)Push(val int){
	s.items = append(s.items,val)
	s.top++
}
// stack pop
func (s *Stack)Pop(){
	if  s.top == -1{
		fmt.Println("Empty Stack")
		return
	}
	s.items = s.items[:len(s.items)-1]
	s.top--
}
// stack peek
func (s *Stack)Peek(){
	if s.top == -1{
		fmt.Println("Empty Stack")
		return
	}
	fmt.Println("Top element of stack is:", s.items[s.top])
}
// print stack
func (s *Stack)PrintStack(){
	fmt.Println("Your stack:")
	fmt.Println(s.items)
}

// Queue
type Queue struct{
	front *Node
	rear *Node
}

func newQueue()*Queue{
	return &Queue{
		front: nil,
		rear: nil,
	}
}
// eequeue
func (q *Queue)Enqueue(val int){
	newNode := &Node{data: val}
	if q.front == nil{
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}
// dequeue
func (q *Queue)Dequeue()int{
	if q.front == nil{
		fmt.Println("Queue is empty")
		return -1
	}
	val := q.front.data
	q.front = q.front.next

	if q.front == nil{
		q.rear = nil
	}
	return val
}
// print queue
func (q *Queue)PrintQ(){
	if q.front == nil{
		fmt.Println("Empty Queue")
		return
	}
	temp := q.front
	fmt.Println("Queue:")
	for temp != nil{
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main(){
	//insert
	arr := []int{4,8,70,25,14,22}
	fmt.Println("Befor insert:",arr)
	result := insertArr(arr,4,1000)
	fmt.Println("After insert:",result)
	//delete
	deleteArr := deleteValue(arr, 2)
	fmt.Println("After delete:", deleteArr)
	//search 
	searchArr := searchValue(arr,14)
	fmt.Printf("value is in No.%d position\n",searchArr)

	//linked list
	ll := &LinkedList{}
	ll.Insert(55)
	ll.Insert(74)
	ll.Insert(18)
	ll.PrintList()
	ll.DeleteVal(74)
	ll.PrintList()
	lLresult := arrayToLl([]int{1,7,8,88,4,6,99})// convert array into ll
	lLresult.PrintList() 

	//stack
	head := newStack()
	head.Push(55)
	head.Push(43)
	head.Push(22)
	head.PrintStack()
	head.Pop()
	head.PrintStack()
	head.Peek()

	// Queue
	q := newQueue()
	q.Enqueue(77)
	q.Enqueue(43)
	q.Enqueue(66)
	q.PrintQ()
	q.Dequeue()
	q.PrintQ()
}