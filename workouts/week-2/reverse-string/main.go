package main

import "fmt"

func reverse(s string)string{
	runes := []rune(s)
	
	head := 0
	tail := len(runes)-1

	for head < tail{
		runes[head], runes[tail] = runes[tail], runes[head]
		head++
		tail--
	}
	return string(runes)

}

func main(){
	result := reverse("Akhil")
	fmt.Println(result)
}