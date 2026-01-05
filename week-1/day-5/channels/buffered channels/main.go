package main

import (
	"fmt"
)

func main(){
	ch := make(chan int,2)

	ch <- 5
	ch <- 10

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}