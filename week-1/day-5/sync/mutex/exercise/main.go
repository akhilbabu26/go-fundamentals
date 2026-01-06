package main

import (
	"sync"
	"fmt"
)

var(
	counter = 0
	mu sync.Mutex
)

func incerment(){
	mu.Lock()
	counter += 1
	mu.Unlock()
}

func main(){
	wg := sync.WaitGroup{}

	for i := 1; i<=5; i++{
		wg.Add(1)

		go func(){
			defer wg.Done()
			incerment()
		}()
	}

	wg.Wait()
	fmt.Println("Result:",counter)
}