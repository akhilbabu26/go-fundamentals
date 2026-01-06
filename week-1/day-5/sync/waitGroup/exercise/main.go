package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup

	for i := 1; i<=5; i++{
		wg.Add(1)
		go func (n int){
			defer wg.Done()
			fmt.Println("count:",n)
		}(i)
	}

	wg.Wait()
	fmt.Println("All done")
}