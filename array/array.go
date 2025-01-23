package main

import "fmt"

var productName [5]string

func main() {
	productName[0] = "Milk"
	productName[1] = "Bread"
	productName[2] = "Water"
	productName[3] = "Tea"
	productName[4] = "Coffee"
	fmt.Println(productName)

	price := [5]float32{1000, 2000, 3000, 4000, 5000}
	fmt.Println(price)
}
