package main

import "fmt"

const count = 10

func main() {
	i := 0
	for i < 10 {
		fmt.Println("i =", i)
		i = i + 1
	}

	for f := 0; f < count; f++ {
		fmt.Println("f =", f)
	}

	var input string

	for {
		fmt.Scanf("%s", &input)
		fmt.Println("input =", input)
		if input == "exit" {
			break
		}
	}
}
