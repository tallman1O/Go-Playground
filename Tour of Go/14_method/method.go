package main

// methods in Go are just functions with a receiver
import "fmt"

// Rectangle define a struct
type Rectangle struct {
	width, height float64
}

// Area method with receiver (r Rectangle)
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	rect := Rectangle{10, 5}
	fmt.Println(rect.Area()) // 50
}
