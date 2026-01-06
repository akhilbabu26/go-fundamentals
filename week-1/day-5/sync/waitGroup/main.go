package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	wg.Add(4)

	go func (){
		defer wg.Done()
		fmt.Println("Frist msg")
	}()

	go func (){
		defer wg.Done()
		fmt.Println("Second msg")
	}()

	go func (){
		defer wg.Done()
		fmt.Println("3rd msg")
	}()

	go func(){
		defer wg.Done()
		fmt.Println("4th msg")
	}()

	wg.Wait()
	fmt.Println("All msg printed")
}