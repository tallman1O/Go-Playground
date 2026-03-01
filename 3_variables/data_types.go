package main

import "fmt"

var (
	// boolean
	isGoFun bool = true

	// integer
	age int = 30

	// floating-point
	pi float64 = 3.14159

	// string
	greeting string = "Hello, Go!"

	// byte (alias for uint8)
	firstLetter byte = 'G'

	// rune (alias for int32, represents a Unicode code point)
	unicodeChar rune = '世'
)

func main() {
	fmt.Println("Boolean:", isGoFun)
	fmt.Println("Integer:", age)
	fmt.Println("Float:", pi)
	fmt.Println("String:", greeting)
	fmt.Printf("Byte: %c\n", firstLetter) // Print byte as character
	fmt.Printf("Rune: %c\n", unicodeChar) // Print rune as character

	// Default values of data types
	var i int
	var f float64
	var b bool
	var s string
	// printf is used to format the output, %v is a placeholder for the value of the variable,
	// %q is used to print the string with quotes
	fmt.Printf("int: %v, float64: %v, bool: %v, string: %q\n", i, f, b, s)
}

// Types of Data Types in Go:
// 1. Boolean: Represents true or false values.
// 2. Integer: Represents whole numbers (e.g., int, int8, int16, int32, int64).
// 3. Floating-point: Represents decimal numbers (e.g., float32, float64).
// 4. String: Represents a sequence of characters.
// 5. Aggregate Types: Arrays, Slices, Maps, Structs, etc.
// 6. Reference Types: Pointers, Channels, Interfaces, etc.
