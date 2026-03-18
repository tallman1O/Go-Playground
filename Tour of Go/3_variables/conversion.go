package main

import "math"

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	println(x, y, z)
}

// In Go, we need to explicitly convert types.
// In the example above, we convert the integer values of x and y to float64
// before performing the square root operation, and then convert the result back to uint.
// IDES Can catch these errors  at compile time itself.
