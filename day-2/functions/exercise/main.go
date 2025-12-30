package main

import "fmt"

func allDetails(a float64, b float64)(sum float64, difference float64, product float64, quotient float64, ok bool){
	sum = a+b
	difference = a-b
	product = a*b
	
	if b == 0 {
		quotient = 0
		ok = false
		return
	}
	quotient = a/b
	ok = true
	return

}

func main(){
	sum, difference, product, quotient, ok := allDetails(8,7)

	fmt.Println("Sum:",sum)
	fmt.Println("Differance:",difference)
	fmt.Println("Product:",product)	
	if ok{
		fmt.Println("Quotient:",quotient)
	}else{
		fmt.Println("Can't devisible by zero")
	}
}