package main

import (
	"fmt"
	"errors"
)

type Students struct{
	Name string
	Age int
}

var students []Students

//storing
func AddStudents(a string, b int)(string, int, error){
	if a == ""{
		return a,b, errors.New("Enter Your name")
	}
	if b < 0{
		return a,b,errors.New("Age can't be less than zero")
	}

	students = append(students, Students{
		Name: a,
		Age: b,
	})

	return a,b,nil
}

//listing students
func ListStudents(){
	if len(students)==0{
		fmt.Println("No students found")
		return
	}
	fmt.Println("Student list:")
	for i,s := range students{
		fmt.Printf("ID:%d,:Name: %s, Age: %d\n",i,s.Name,s.Age)
	}
}

//updating
func UpdateStudent(index int, newName string, newAge int){

	if index < 0 || index >= len(students){
		fmt.Println("Error: enter valid number")
		return
	}

	students[index] = Students{
		Name: newName,
		Age: newAge,
	}
	fmt.Println("Updated students List:",students)
}

//deleting
func DeletingStudents(index int){
	if index < 0 || index >= len(students){
		fmt.Println("Error: enter valid number")
		return
	}
	students = append(students[:index],students[index+1:]...)
	fmt.Println("New student list:",students)
}

func main(){
	AddStudents("Akhil", 24)
	AddStudents("Manu", 44)
	AddStudents("Junaid", 34)

	ListStudents()

	UpdateStudent(2, "Jefin", 23)

	DeletingStudents(0)

	ListStudents()
}