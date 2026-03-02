package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// When you run this program, you will see that the numbers are printed in reverse order, from 9 down to 0. \
// This is because the defer statements are executed in a last-in-first-out (LIFO) order when the surrounding function returns.
// In this case, the main function will return after the loop finishes, and all the deferred calls will be executed in reverse order.

// The output of the program will be:
// counting
// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0

// Defer is often used to ensure that resources are released properly, such as closing files or network connections, even if an error occurs. In this example, we are simply demonstrating the order of execution of deferred function calls.
