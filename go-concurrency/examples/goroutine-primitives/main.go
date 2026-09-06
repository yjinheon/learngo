package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	go someFunc("1")
	go someFunc("2")
	go someFunc("3")
	go someFunc("4")
	go someFunc("5")

	time.Sleep(time.Second * 2)
	fmt.Println("hi")
}
