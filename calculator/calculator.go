package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v Must be a number", promt)
		panic(message)
	}

	return value
}

func getOperator() string {
	fmt.Print(("Operator is (+,-,*,/): "))
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func Add(value1, value2 float64) float64 {
	return value1 + value2
}

func Minus(value1, value2 float64) float64 {
	return value1 - value2
}

func Multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func Divide(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	value1 := getInput("value1 = ")
	value2 := getInput("value2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = Add(value1, value2)
	case "-":
		result = Minus(value1, value2)
	case "*":
		result = Multiply(value1, value2)
	case "/":
		result = Divide(value1, value2)
	default:
		panic("Wrong operator")
	}
	fmt.Printf("Result is = %v", result)
}
