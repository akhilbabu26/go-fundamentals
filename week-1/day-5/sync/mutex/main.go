package main

import (
	"fmt"
	"sync"
)

var (
	counter = 0
	mu sync.Mutex
)

func increment(){
	mu.Lock()
	counter += 1
	mu.Unlock()
}

func main(){
	wg := sync.WaitGroup{}

	for i :=0; i<=5; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:",counter)
}
