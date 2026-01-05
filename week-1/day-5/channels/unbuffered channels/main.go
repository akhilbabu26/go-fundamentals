package main

import (
	"fmt"
)

// func main(){
// 	ch := make(chan int)

// 	go func(){
// 		fmt.Println("Sending...")
// 		ch <- 10
// 	}()

// 	fmt.Println("Receving...")
// 	value := <- ch
// 	fmt.Println("Receved:",value)
// }

func main() {
    ch := make(chan int)
    done := make(chan struct{})

    go func() {
		fmt.Println("Sending...")
		ch <- 10
        close(done)  
    }()

    fmt.Println("Receving...")
    value := <-ch
    fmt.Println("Receved:", value)
    <-done
}