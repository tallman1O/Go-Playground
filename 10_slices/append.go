package main

import "fmt"

func appendOnSlice() {
	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)
}
