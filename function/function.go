package main

import "fmt"

func name() {
	fmt.Println("Hello, My name is Golang")
}

// func sum(value1 int, value2 int) {
// 	result := value1 + value2
// 	fmt.Println("Sum Result =", result)
// }

func sum(value1, value2 int) int {
	return value1 + value2
}

func main() {
	name()
	// sum(2, 2)

	result := sum(5, 5)
	fmt.Println("Sum Result =", result)
}
