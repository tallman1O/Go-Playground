package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	var s []int = primes[1:4]
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	fmt.Println(s)
	s[0] = 42

	names[2] = "mehul"
	fmt.Println(names)

	fmt.Println(s)
	appendOnSlice()

	// initiaze slice using make()
	s1 := make([]int, 5)
	s2 := make([]int, 5, 10) // length 5, capacity 10
	fmt.Println(s1, s2)
}
