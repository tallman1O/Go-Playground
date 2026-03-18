package main

import "fmt"

//A closure = function + captured variables from outer scope
//It remembers variables even after the outer function is done.

func adder() func(int) int {
	sum := 0 // this variable gets "captured"
	return func(x int) int {
		sum += x // modifies captured variable
		return sum
	}
}

func main() {
	pos := adder()

	fmt.Println(pos(1)) // 1
	fmt.Println(pos(2)) // 3
	fmt.Println(pos(3)) // 6
}

// adder() runs → creates sum = 0
//returns an anonymous function
//that function keeps access to sum
//every time you call it → it updates the SAME sum
//It’s not resetting. It’s stateful.
//Each call to adder() creates a separate closure

//Where closures are actually useful
//1. maintaining state without structs (quick hacks)
//2. middleware (very common in backend)
//3. function factories
//4. callbacks with memory
