package main

import "fmt"

func main(){
	//if-else statement
	var votingAge int = 50

	if votingAge > 18{
		fmt.Println("Eligible to vote")
	}else{
		fmt.Println("Not eligible")
	}

	//if - else if - else statement
	var marks int = 99

	if marks > 90{
		fmt.Println("Pass Grade A")
	}else if marks >= 70{
		fmt.Println("Pass Grade B")
	}else if marks >= 60{
		fmt.Println("Pass Grade C")
	}else if marks >= 50{
		fmt.Println("Pass Grade D")
	}else if marks >= 40{
		fmt.Println("Just Pass Grade E")
	}else{
		fmt.Println("Fail Grade F")
	}

	//short statement in if

	if score := 90; score > 90{
		fmt.Println("A+")
	}else if score >= 70{
		fmt.Println("A")
	}else if score >= 60{
		fmt.Printf("B")
	}else if score >= 50{
		fmt.Println("C")
	}else if score >= 40{
		fmt.Println("D")
	}else if score >= 30{
		fmt.Println("E")
	}else{
		fmt.Println("F")
	}

	//switch
	var age int = 25

	switch {
		case age < 18:
			fmt.Println("Minor")
		case age < 60:
			fmt.Println("Adult")
		default:
			fmt.Println("Senior")
	}

	// other switch cases
	var day int = 1

	switch day{
		case 7:
			fmt.Println("Weekend")
		case 1,2,3,4,5,6:
			fmt.Println("weekday")
		default:
			fmt.Println("Invalid day")
	}

}