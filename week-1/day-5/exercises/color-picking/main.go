package main

import (
	"fmt"
	"sync"
)

func main(){
	color := []string{"green","orange"}
	pick := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		defer wg.Done()
		for c := range pick{
			fmt.Println("Color:",c)
		}
	}()
	
	for _,v := range color{
		pick <- v
	}
	close(pick)
	wg.Wait()
}