package main

import "fmt"

func main() {
	var CarBrandName = []string{"Toyota", "Mazda"}
	fmt.Println(CarBrandName)
	CarBrandName = append(CarBrandName, "Honda", "Nissan", "Suzuki")
	fmt.Println(CarBrandName)

	CarBrandNameIChoose := CarBrandName[2:5]
	fmt.Println(CarBrandNameIChoose)

	CarBrandNameIChoose = CarBrandName[:3]
	fmt.Println(CarBrandNameIChoose)
}
