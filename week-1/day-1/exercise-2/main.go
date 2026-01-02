package main

import "fmt"

func main(){
	value1 := 2
	value2 := 3
	operator := '-'

	switch operator{
		case '+':
			result := value1 + value2
			fmt.Println(result)
		case '-':
			result := value1 - value2
			fmt.Println(result)
		case '*':
			result := value1 * value2
			fmt.Println(result)
		case '/':
			result := float64(value1) / float64(value2)
			fmt.Println(result)
	}
}