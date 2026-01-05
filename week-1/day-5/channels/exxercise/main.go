package main

import "fmt"

func main(){
	jobs := make(chan int)
	done := make(chan struct{})

	//consumer
	go func(){
		for value := range jobs{
			fmt.Println("Consumed:",value)
		}
		fmt.Println("Consumer Finished")
		done <- struct{}{}
	}()
	
	//produced
	for i :=1; i<=5; i++{
		fmt.Println("Producer:",i)
		jobs <-i
	}

	close(jobs)
	<-done
	fmt.Println("All done")
}