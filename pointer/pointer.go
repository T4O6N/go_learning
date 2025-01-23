package main

import "fmt"

func zerovalue(ivalue int) {
	fmt.Println("ivalue =", ivalue)
	ivalue = 0
}

func zeropointer(ipointer *int) {
	*ipointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)

	zerovalue(i)
	fmt.Println("i from function zerovalue =", i)

	zeropointer(&i)
	fmt.Println("i address from function zeropointer =", &i)
}
