package main

import "fmt"

func withdraw(amount, balance float64)(float64, error){
    if amount > balance{
        return balance, fmt.Errorf("insufficient funds: have: %.2f, need: %.2f ", balance, amount)
    }

    return balance - amount, nil   
}

func main(){
    newBalance, err := withdraw(3000,2000)

    if err != nil{
        fmt.Println("Transaction failed:", err)
        return
    }
    fmt.Println("New balance:", newBalance)
}
