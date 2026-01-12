package main

import (
	"fmt"
	"errors"
)

type Task struct{
	Name string
}

var task []Task

func AddTask(a string)(string, error){
	if a == ""{
		return a,errors.New("ERROR: Enter your task")
	}

	task = append(task,Task{Name: a})
	return a,nil
}

func UpdateTask(index int, tsk string)error{
	if index < 0 || index >= len(task) || tsk == ""{
		fmt.Println("ERROR: Invalid input")
	}

	task[index] = Task{Name: tsk}
	fmt.Println("Updated task:")
	return nil
}

func DeleteTask(index int)error{
	if index < 0 || index >= len(task){
		fmt.Println("Error: Task not found")
	}

	task = append(task[:index], task[index+1:]...)
	fmt.Println("Updated task:")
	return nil
}

func ListTask(){
	if len(task)==0{
		fmt.Println("There is no task!!")
	}

	for i,v:= range task{
		fmt.Printf("Task No:%d = %s\n",i,v)
	}
}

func main(){
	AddTask("Make dinner")
	AddTask("Make breakfast")
	AddTask("go to sleep")
	AddTask("Eat dinner")

	ListTask()

	UpdateTask(0, "wake up at 6'o clock")
	ListTask()

	DeleteTask(0)
	ListTask()
}