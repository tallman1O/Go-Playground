package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer p
	fmt.Println(j) // see the new value of j
}
