package main

import (
	"fmt"
	"math"
)

// compute takes a function (fn) as input and calls it with (3,4)
func compute(fn func(float64, float64) float64) float64 {
	return fn(5, 12) // execute passed function
}

func main() {
	// anonymous function assigned to hypot (calculates sqrt(x^2 + y^2))
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y) // Pythagorean formula
	}

	fmt.Println(hypot(3, 4))    // direct call → 5
	fmt.Println(compute(hypot)) // pass function → compute calls it → 13
}
