package main

import "fmt"

// fibonacci returns a function that returns
// successive fibonacci numbers from each
// successive call

func fibo() func() int {
	first, second := 0, 1
	return func() int {
		res := first

		first, second = second, first+second

		return res
	}
}

func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
