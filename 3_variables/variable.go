package main

import "fmt"

var c, pyhon, java bool

const Pi = 3.14

func main() {

	var b string
	a := 7 // same as var a int = 7

	println(a, b, c, pyhon, java)
	fmt.Printf("%T\n", a)
	fmt.Println(Pi)
}

// Rules :
// 1. Variable names must start with a letter or an underscore.
// 2. Use := for local variables (Inside functions)
// 3. Use var for package-level variables (Outside functions)
