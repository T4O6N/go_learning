package main

import "fmt"

// func main() {
// 	myMoney := 10000
// 	if myMoney > 1000 {
// 		fmt.Println("I have enough money")
// 	} else {
// 		fmt.Println("I don't have enough money")
// 	}
// }

// var score int

// func main() {
// 	fmt.Println("Grade calculator")
// 	fmt.Scanf("%d", &score)
// 	fmt.Println("Your score is =", score)
// }

var score int

func main() {
	fmt.Println("Grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println("Your score is =", score)
	if score >= 80 {
		fmt.Println("Your grade is A")
	} else if score >= 70 {
		fmt.Println("Your grade is B")
	} else if score >= 60 {
		fmt.Println("Your grade is C")
	} else if score >= 50 {
		fmt.Println("Your grade is D")
	} else {
		fmt.Println("Your grade is F")
	}
}
