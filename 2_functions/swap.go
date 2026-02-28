package main

// syntax overview: func function_name() (string, string) -
// When a function returns multiple values,
// we can specify the return types in parentheses.

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	println(a, b)
}
