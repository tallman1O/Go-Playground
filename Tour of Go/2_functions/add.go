package main

// syntax overview: func function_name(parameter_list parameter_type) return_type {
//     // function body
// }

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func main() {
	println(add(1, 2))
	print(add2(2, 3))
}
