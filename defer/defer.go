package main

import (
	"fmt"
)

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Println("result :", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i :", i)
	}
}

func deferLoop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("defer i :", i)
	}
}

func main() {
	fmt.Println("Hello World")
	defer fmt.Println("End World")
	defer add(15, 15)

	loop()
	deferLoop()
}
