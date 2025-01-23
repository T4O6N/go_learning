package main

import "fmt"

func main() {
	value := 4
	switch value {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Invalid value")
	}

	var input string

	fmt.Println("- Input Color (blue, red, green, purple):")
	fmt.Scan(&input)
	fmt.Println("Input Color =", input)
	switch input {
	case "blue":
		fmt.Println("#0000FF")
	case "red":
		fmt.Println("#FF0000")
	case "green":
		fmt.Println("#00FF00")
	case "purple":
		fmt.Println("#FF00FF")
	default:
		fmt.Println("No Color Matched")
	}
}
