package main

import (
	"fmt"
	"sync"
)

func main(){
	odd := make(chan struct{})
	even := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func (){
		defer wg.Done()
		for i:=1; i<=9; i+=2{
			<-odd
			fmt.Println("ODD:",i)
			if i < 9{
				even <-  struct{}{}
			}
		}
		close(even)
	}()

	go func (){
		defer wg.Done()
		for i:=2; i<=10; i+=2{
			<-even
			fmt.Println("EVEN:",i)
			if i < 10{
				odd <- struct{}{}
			}
		}
	}()

	odd <- struct{}{}
	wg.Wait()
}