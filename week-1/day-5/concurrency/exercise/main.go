package main

import (
	"fmt"
	"sync"
)

func main(){
	odd := make(chan struct{})
	even := make(chan struct{})
	var wait sync.WaitGroup
	wait.Add(2)

	// Odd goroutine
	go func(){
		defer wait.Done()
		for i := 1; i<=9; i+=2{
			<-odd
			fmt.Println("Odd:", i)
			even <- struct{}{}
		}
	}()
	
	//Even goroutine
	go func(){
		defer wait.Done()
		for i := 2; i<=10; i+=2{
			<-even
			fmt.Println("Even:", i)
			if i < 10{
				odd <- struct{}{}
			}
		}
	}()

	odd <- struct{}{}// giving odd a value
	wait.Wait()
}