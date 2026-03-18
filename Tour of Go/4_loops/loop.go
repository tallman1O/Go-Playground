package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("current sum: ", sum)
	}
	fmt.Println("total sum: ", sum)
}

// In Go, there is only one type of loop: the for loop. The for loop can be used in three different ways:

// 1. Traditional for loop: This is the most common way to use a for loop,
// where you specify an initialization statement, a condition, and a post statement.
// Example:
// for i := 0; i < 5; i++ {
//     fmt.Println(i)
// }

// 2. While-like for loop: You can omit the initialization and post statements to
// create a loop that behaves like a while loop.
// Example:
// i := 0
// for i < 5 {
//     fmt.Println(i)
//     i++
// }

// 3. Infinite loop: You can omit all three components to create an infinite loop.
// Example:
// for {
//     fmt.Println("This will run forever")
// }

// The for loop is versatile and can be used in various ways to iterate over collections,
// perform repeated actions, and more.
