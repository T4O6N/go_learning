package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("Hello World")
	go f("Who i am")
	time.Sleep(5 * time.Second)
}
