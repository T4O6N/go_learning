package main

import "fmt"

var product = make(map[string]float32)

func main() {
	product["Milk"] = 1000
	product["Bread"] = 2000
	product["Water"] = 3000
	fmt.Println("products =", product)

	product["Milk"] = 1500
	fmt.Println("update product =", product)

	delete(product, "Water")
	fmt.Println("delete product =", product)

	value1 := product["Milk"]
	fmt.Println("value1 =", value1)

	courseName := map[string]string{"101": "Golang", "102": "Python"}
	fmt.Println("courseName =", courseName)
}
