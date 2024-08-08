package main

import (q
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}


func main() {
	go someFunc("1") // a goroutine is started. asynchrounous
	go someFunc("2")
	go someFunc("3")
	go someFunc("4")
	go someFunc("5")

	time.Sleep(2 * time.Second) // wait for 1 second

	fmt.Println("Hello, World!")
}