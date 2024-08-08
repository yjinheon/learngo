package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	var f func(int, int) int // function signature
	a, b := 10, 20

	f = add
	fmt.Println(f(a, b))
	f = sub
	fmt.Println(f(a, b))
}
