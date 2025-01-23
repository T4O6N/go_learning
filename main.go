package main

import "fmt"

var numInt, numInt2 int = 1000, 2000
var msg string = "Hello"

func main() {
	numFloat := 5.55
	fmt.Println(numInt)
	fmt.Println(numInt2)
	fmt.Println(numFloat)

	fmt.Println(numInt + numInt2)
	fmt.Println(float64(numInt) + numFloat)

	fmt.Println(msg)
	fmt.Println(msg + " Jong Event Service")
	fmt.Println("my money:", numInt)
}
