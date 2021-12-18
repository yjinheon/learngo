package main

import "fmt"

// fibonacci is a function that returns a function that returns an int

func fibo() func() int {
	a := 0
	b := 1

	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibo()

	for i := 0; i < 34; i++ {
		fmt.Println(f())
	}
}
