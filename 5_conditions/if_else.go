package main

import "fmt"

func main() {

	// if-else statement
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if statement
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// if with OR conditions
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("At least one of the numbers is even")
	}

	//if-else if statement with AND condition
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

}
