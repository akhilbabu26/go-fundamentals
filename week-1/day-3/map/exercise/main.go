package main

import "fmt"

func main(){
	students := make(map[string]int)

	students["Amal"] = 90
	students["Akhil"] = 60
	students["Juniad"] = 80
	students["Jefin"] = 90
	
	for k, v:= range students{
		fmt.Println(k,":",v)
	}
	fmt.Println("-------------------------------------")

	//updating
	students["Akhil"] = 70
	students["jefin"] = 100

	for i:= range students{ // shorter for loop
		fmt.Println(i,":",students[i])
	}
	fmt.Println("-------------------------------------")

	// checking
	name := "Akhil"
	score, ok := students[name]
	if ok{
		fmt.Printf("%s's grade is %d",name,score)
	}else{
		fmt.Printf("%s's grade not found!! Please enter correct name",name)
	}

	fmt.Println("-------------------------------------")
}